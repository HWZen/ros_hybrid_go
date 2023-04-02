package ros_hybrid_go

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/HWZen/ros_hybrid_go/pkg/protobuf"
	ros_hybrid_go "github.com/HWZen/ros_hybrid_go/src/ros_hybrid_error"
	"google.golang.org/protobuf/proto"
)

const delimiter_str = "___delimiter___"

var delimiter = []byte(delimiter_str)

type Node struct {
	connect                net.Conn
	connect_scan           *bufio.Scanner
	stop_cond              sync.Cond
	agent_config           *protobuf.AgentConfig
	subscriberObservers    map[string]msgObserver
	serviceServerObservers map[string]msgObserver
	serviceClientObservers map[string]msgObserver
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
		Node:       name,
		IsProtobuf: proto.Bool(true),
		Delimiter:  delimiter,
	}

	out, err := proto.Marshal(&agentConfig)
	if err != nil {
		return nil, err
	}
	conn.Write(out)

	if !scanner.Scan() {
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
		connect:                conn,
		connect_scan:           scanner,
		agent_config:           &agentConfig,
		stop_cond:              *sync.NewCond(&sync.Mutex{}),
		subscriberObservers:    make(map[string]msgObserver),
		serviceServerObservers: make(map[string]msgObserver),
		serviceClientObservers: make(map[string]msgObserver),
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
		case protobuf.Command_PUBLISH:
			publish := comment.GetPublish()
			if publish == nil {
				fmt.Println("get publish command error: ", comment)
				continue
			}
			if observer, ok := node.subscriberObservers[publish.Topic]; ok {
				go observer.Update(publish)
			} else {
				// log it: no observer for this topic
				fmt.Println("no observer for this topic: ", publish.Topic)
			}
		case protobuf.Command_CALL_SERVICE:
			callService := comment.GetCallService()
			if callService == nil {
				fmt.Println("get call service command error: ", comment)
				continue
			}
			if observer, ok := node.serviceServerObservers[callService.Service]; ok {
				go observer.Update(callService)
			} else {
				// log it: no observer for this service
				fmt.Println("no observer for this service: ", callService.Service)
			}
		case protobuf.Command_RESPONSE_SERVICE:
			responseService := comment.GetResponseService()
			if responseService == nil {
				fmt.Println("get response service command error: ", comment)
				continue
			}
			if observer, ok := node.serviceClientObservers[responseService.Service]; ok {
				go observer.Update(responseService)
			} else {
				// log it: no observer for this service
				fmt.Println("no observer for this service: ", responseService.Service)
			}
		case protobuf.Command_LOG:
			log := comment.GetLog()
			// log it
			fmt.Println("get log from server: \n", log)
		default:
			fmt.Println("get unknown command from server: \n", comment)
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

func (node *Node) Run() {
	recver(node)
}

func (node *Node) Shutdown() {
	// close all channels
	for _, observers := range node.subscriberObservers {
		observers.Shutdown()
	}
	node.subscriberObservers = make(map[string]msgObserver)

	for _, observers := range node.serviceServerObservers {
		observers.Shutdown()
	}
	node.serviceServerObservers = make(map[string]msgObserver)

	for _, observers := range node.serviceClientObservers {
		observers.Shutdown()
	}
	node.serviceClientObservers = make(map[string]msgObserver)

	// wait a second, unsubscribe\unadvertise all server
	time.Sleep(time.Second)
	// close connection
	node.connect.Close()
	// send shutdown signal
	node.stop_cond.Broadcast()
}

func (node *Node) Spin() {
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
	return err
}
