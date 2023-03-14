package ros_hybrid_go

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"time"
	"sync"

	"github.com/HWZen/ros_hybrid_go/pkg/protobuf"
	"github.com/HWZen/ros_hybrid_go/src/ros_hybrid_error"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	set "github.com/deckarep/golang-set"
)

const delimiter_str = "___delimiter___"
var delimiter = []byte(delimiter_str)

type Node struct {
	connect 					net.Conn
	connect_scan 				*bufio.Scanner
	stop_cond 					sync.Cond
	agent_config 				*protobuf.AgentConfig
	subscriber_channels 		map[string]set.Set
	service_server_channels 	map[string]*chan *protobuf.Command_CallService
	service_client_channels 	map[string]set.Set
	// set a logger
}

func init_connection(host string, name string) (*Node, error) {

	// connect to server
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(conn)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, delimiter); i >= 0 {
			// We have a full newline-terminated line.
			return i + len(delimiter), data[0:i], nil
		}
		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	})

	// login to server
	agentConfig := protobuf.AgentConfig{
		Node: name,
		IsProtobuf: proto.Bool(true),
		Delimiter: delimiter,
	}

	out, err := proto.Marshal(&agentConfig)
	if err != nil {
		return nil, err
	}
	conn.Write(out)

	if !scanner.Scan(){
		return nil, ros_hybrid_go.NewError("login failed")
	}
	data := scanner.Bytes()

	comment := &protobuf.Command{}
	if err := proto.Unmarshal(data, comment); err != nil {
		return nil, err
	}

	if comment.Type != protobuf.Command_LOG && comment.GetLog().Message != "login success\n" {
		return nil, ros_hybrid_go.NewError("login failed:" + comment.String())
	}

	return &Node{
		connect: conn,
		connect_scan: scanner,
		agent_config: &agentConfig,
		stop_cond: *sync.NewCond(&sync.Mutex{}),
		subscriber_channels: make(map[string]set.Set),
		service_server_channels: make(map[string]*chan *protobuf.Command_CallService),
		service_client_channels: make(map[string]set.Set),
		}, nil

}

func recver(node *Node) {
    for node.connect_scan.Scan() {
        data := node.connect_scan.Bytes()
        comment := &protobuf.Command{}
        if err := proto.Unmarshal(data, comment); err != nil {
			fmt.Println("unmarshal error: ", err)
            continue
        }
		switch comment.Type {
        case protobuf.Command_PUBLISH :
			publish := comment.GetPublish()
			if channelSet, ok := node.subscriber_channels[publish.Topic]; ok {
				for channel := range channelSet.Iter() {
					*channel.(*chan []byte) <- publish.GetData()
				}
			} else {
				response := protobuf.Command{
					Type: protobuf.Command_LOG,
					Log: &protobuf.Command_Log{
						Level : protobuf.Command_Log_ERROR,
						Message : "no subscriber for topic: " + publish.Topic,
						Time: timestamppb.Now(),
					},
				}
				err := node.Send(&response)
				if err != nil {
					// log it
					continue
				}
			}
		case protobuf.Command_CALL_SERVICE :
			call_service := comment.GetCallService()
			if channel, ok := node.service_server_channels[call_service.Service]; ok {
				*channel <- call_service
			} else {
				response := protobuf.Command{
					Type: protobuf.Command_LOG,
					Log: &protobuf.Command_Log{
						Level : protobuf.Command_Log_ERROR,
						Message : "no advertise for service: " + call_service.Service,
						Time: timestamppb.Now(),
					},
				}
				err := node.Send(&response)
				if err != nil {
					// log it
					continue
				}
			}
		case protobuf.Command_RESPONSE_SERVICE :
			response_service := comment.GetResponseService()
			if channelSet, ok := node.service_client_channels[response_service.Service]; ok {
				for channel := range channelSet.Iter() {
					*channel.(*chan *protobuf.Command_ResponseService) <- response_service
				}
			} else {
				response := protobuf.Command{
					Type: protobuf.Command_LOG,
					Log: &protobuf.Command_Log{
						Level : protobuf.Command_Log_ERROR,
						Message : "no client for service: " + response_service.Service,
						Time: timestamppb.Now(),
					},
				}
				err := node.Send(&response)
				if err != nil {
					// log it
					continue
				}
			}
		case protobuf.Command_LOG :
			log := comment.GetLog()
			// log it
			fmt.Println("get log from server: \n", log )
		default:
			fmt.Println("get unknown command from server: \n", comment )
		}
    }
	if err := node.connect_scan.Err(); err != nil {
		fmt.Println("error: ", err)
	}
	node.Shutdown()
}

func NewNode(host string, name string) (*Node, error) {
    node, err := init_connection(host, name)
    if err != nil {
        return nil, err
    }
    return node, nil
}



func (node *Node) Run(){
	recver(node)
}

func (node *Node) Shutdown(){
	// close all channels
	for _, chSet := range node.subscriber_channels {
		chSet.Each(func(v interface{}) bool {
			close(*v.(*chan []byte))
			return true
		})
	}
	for k := range node.subscriber_channels {
		delete(node.subscriber_channels, k)
	}
	
	for _, ch := range node.service_server_channels {
		close(*ch)
	}
	for k := range node.service_server_channels {
		delete(node.service_server_channels, k)
	}
	
	for _, chSet := range node.service_client_channels {
		chSet.Each(func(v interface{}) bool {
			close(*v.(*chan *protobuf.Command_ResponseService))
			return true
		})
	}
	for k := range node.service_client_channels {
		delete(node.service_client_channels, k)
	}
	

	// wait a second, unsubscribe\unadvertise all server
	time.Sleep(time.Second)
	// close connection
	node.connect.Close()
	// send shutdown signal
	node.stop_cond.Broadcast()
}

func (node *Node) Spin(){
	node.stop_cond.L.Lock()
	node.stop_cond.Wait()
	node.stop_cond.L.Unlock()
}

func (node *Node) Send(data *protobuf.Command) error {
	buf, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	_, err = node.connect.Write(buf)
	if err != nil {
		return err
	}
	_, err = node.connect.Write(delimiter)
	if err != nil {
		return err
	}
	return nil
}
