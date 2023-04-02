package ros_hybrid_go

import (
	"github.com/HWZen/ros_hybrid_go/pkg/protobuf"
	ros_hybrid_go "github.com/HWZen/ros_hybrid_go/src/ros_hybrid_error"
	"google.golang.org/protobuf/proto"
)

type SubscriberCallbackFunc func(msg proto.Message)

type SubscribeHandler struct {
	Subscriber *_subscriber
	MsgCallback SubscriberCallbackFunc
}

type _subscriber struct {
	node        *Node
	Topic       string
	MsgType     string
	handlers	[]*SubscribeHandler
	MsgSample   proto.Message
}

type Subscriber = SubscribeHandler

func NewSubscriber(node *Node, topic string, msgTypeName string, msgType proto.Message, msgCallback SubscriberCallbackFunc) (*SubscribeHandler, error) {
	var subscriber *_subscriber = nil
	if s,exist := node.subscriberObservers[topic]; !exist {
		subscriber = &_subscriber{
			node:      node,
			Topic:     topic,
			MsgType:   msgTypeName,
			MsgSample: msgType,
		}
		err := subscriber.Subscribe()
		if err != nil {
			return nil, err
		}
		node.subscriberObservers[topic] = subscriber
	} else {
		subscriber = s.(*_subscriber)
	}
	handler := &SubscribeHandler{
		Subscriber: subscriber,
		MsgCallback: msgCallback,
	}
	subscriber.handlers = append(subscriber.handlers, handler)
	return handler, nil
}

func (subscriber *_subscriber) Subscribe() error {
	command := protobuf.Command{
		Type: protobuf.Command_SUBSCRIBE,
		Subscribe: &protobuf.Command_Subscribe{
			Topic: subscriber.Topic,
			Type:  subscriber.MsgType,
		},
	}

	return subscriber.node.Send(&command)
}

func (subscriber *_subscriber) Unsubscribe() error {
	command := protobuf.Command{
		Type: protobuf.Command_UNSUBSCRIBE,
		Unsubscribe: &protobuf.Command_Unsubscribe{
			Topic: subscriber.Topic,
		},
	}
	return subscriber.node.Send(&command)
}

func (handle *SubscribeHandler) Subscribe() error {
	for _, handler := range handle.Subscriber.handlers {
		if handler == handle {
			return ros_hybrid_go.NewError("already subscribed")
		}
	}
	handle.Subscriber.handlers = append(handle.Subscriber.handlers, handle)
	return nil
}

func (handle *SubscribeHandler) Unsubscribe() error {
	for i, handler := range handle.Subscriber.handlers {
		if handler == handle {
			handle.Subscriber.handlers = append(handle.Subscriber.handlers[:i], handle.Subscriber.handlers[i+1:]...)
			return nil
		}
	}
	return ros_hybrid_go.NewError("not subscribed")
}

func (subscriber *_subscriber) Update(publish proto.Message) {
	publishMsg := publish.(*protobuf.Command_Publish)
	msg := proto.Clone(subscriber.MsgSample)
	err := proto.Unmarshal(publishMsg.Data, msg)
	if err != nil {
		return
	}
	for _, handler := range subscriber.handlers {
		handler.MsgCallback(msg)
	}
}

func (subscriber *_subscriber) Shutdown() {
	subscriber.Unsubscribe()
}
