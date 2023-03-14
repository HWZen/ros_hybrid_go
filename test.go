package main

import (
	"fmt"
	"time"

	"github.com/HWZen/ros_hybrid_go/example_protobuf/protobuf"
	"github.com/HWZen/ros_hybrid_go/src"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	node, err := ros_hybrid_go.NewNode("localhost:5150", "go_test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer node.Shutdown()
	go node.Run()
	publisher, err := ros_hybrid_go.NewPublisher(node, "test", "std_msgs/String")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer publisher.Unadvertise()

	subscriber := ros_hybrid_go.NewSubscriber(node, "test", "std_msgs/String", &protobuf.String{}, func(msg proto.Message) {
		real_msg := msg.(*protobuf.String)
		fmt.Println("Received: ", real_msg.Data)
		// sleep 1.1 seconds to simulate a long callback
		time.Sleep(time.Millisecond * 1100)
	})
	err = subscriber.Subscribe()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer subscriber.Unsubscribe()

	for i := 0; i < 10; i++ {
	publisher.Publish(&protobuf.String{Data:"Hello World!" + fmt.Sprint(i)})
	time.Sleep(time.Millisecond * 100)
	}


	service_server, err := ros_hybrid_go.NewServicer(node, 
		"test_service",
		"ros_hybrid_sdk/MyService",
		&protobuf.MyService_Request{}, 
		func(req proto.Message) proto.Message{
			real_req := req.(*protobuf.MyService_Request)
			fmt.Println("Received service request: ", real_req)
			sleep_millisecond := real_req.GetMyMsg().I
			time.Sleep(time.Millisecond * time.Duration(sleep_millisecond))
			return &protobuf.MyService_Response{
				Header: &protobuf.Header{
					Seq: 0,
					Stamp: timestamppb.Now(),
					FrameId: "test",
				},
			}
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	go service_server.Advertise()
	defer service_server.Unadvertise()

	time.Sleep(time.Second * 1)
	service_client := ros_hybrid_go.NewServiceCaller(node, "test_service", "ros_hybrid_sdk/MyService", &protobuf.MyService_Response{})
	defer service_client.Close()
	res, err := service_client.Call(&protobuf.MyService_Request{
		MyMsg: &protobuf.MyMsg{
			I: 1000,
			Strs: []string{"hello", "world", "one", "two", "three"},
			Int5: []*protobuf.Int32{{Data: 1}, {Data: 2}, {Data: 3}, {Data: 4}, {Data: 5}},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received service response: ", res)

	node.Spin()
	fmt.Println("exit")
}