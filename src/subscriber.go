package ros_hybrid_go

import (
	"github.com/HWZen/ros_hybrid_go/pkg/protobuf"
	"google.golang.org/protobuf/proto"
	set "github.com/deckarep/golang-set"
)

type SubscriberCallbackFunc func(msg proto.Message)

type Subscriber struct {
	
	node 		*Node
	Topic 		string
	MsgType 	string
	MsgCallback SubscriberCallbackFunc
	MsgChannel 	chan []byte
	MsgSample 	proto.Message
}

func NewSubscriber(node *Node, topic string, msgTypeName string, msgType proto.Message, msgCallback SubscriberCallbackFunc) *Subscriber {
	subscriber := Subscriber{
		node: node,
		Topic: topic,
		MsgType: msgTypeName,
		MsgCallback: msgCallback,
		MsgChannel: make(chan []byte),
		MsgSample: msgType,

	}
	
	return &subscriber
}

func (subscriber *Subscriber) Subscribe() error {
	command := protobuf.Command{
		Type: protobuf.Command_SUBSCRIBE,
		Subscribe: &protobuf.Command_Subscribe{
			Topic: subscriber.Topic,
			Type: subscriber.MsgType,
		},
	}
	err := subscriber.node.Send(&command)
	if err != nil {
		return err
	}
	if _, ok := subscriber.node.subscriber_channels[subscriber.Topic]; !ok {
		subscriber.node.subscriber_channels[subscriber.Topic] = set.NewSet()
	}
	subscriber.node.subscriber_channels[subscriber.Topic].Add(&subscriber.MsgChannel)
	go subscriber.Run()
	return nil
}

func (subscriber *Subscriber) Unsubscribe() error {
	command := protobuf.Command{
		Type: protobuf.Command_UNSUBSCRIBE,
		Unsubscribe: &protobuf.Command_Unsubscribe{
			Topic: subscriber.Topic,
		},
	}
	err := subscriber.node.Send(&command)
	if err != nil {
		return err
	}
	subscriber.node.subscriber_channels[subscriber.Topic].Remove(&subscriber.MsgChannel)
	return nil
}

func (subscriber *Subscriber) Run() {
	for msgBuf, ok := <- subscriber.MsgChannel; ok; msgBuf, ok = <- subscriber.MsgChannel {
		go func(msgBuf []byte){
			msg := proto.Clone(subscriber.MsgSample)
			err := proto.Unmarshal(msgBuf, msg)
			if err != nil {
				// log it
				return
			}
			subscriber.MsgCallback(msg)
		}(msgBuf)
	}
}

