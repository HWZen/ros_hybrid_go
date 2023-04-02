package ros_hybrid_go

import (
	"github.com/HWZen/ros_hybrid_go/pkg/protobuf"
	"google.golang.org/protobuf/proto"
)

type Publisher struct {
	Node *Node
	Topic string
	Type string
}

func NewPublisher(node *Node, topic string, topicType string) (*Publisher, error) {
	publisher := Publisher{
		Node: node,
		Topic: topic,
		Type: topicType,
	}
	return &publisher, nil
}

func (publisher *Publisher) Publish(msg proto.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	command := protobuf.Command{
		Type: protobuf.Command_PUBLISH,
		Publish: &protobuf.Command_Publish{
			Topic: publisher.Topic,
			Data: data,
		},
	}
	return publisher.Node.Send(&command)
	
}

func (publisher *Publisher) Advertise() error {
	command := protobuf.Command{
		Type: protobuf.Command_ADVERTISE,
		Advertise: &protobuf.Command_Advertise{
			Topic: publisher.Topic,
			Type:  publisher.Type,
		},
	}
	return publisher.Node.Send(&command)
	
}


func (publisher *Publisher) Unadvertise() error {
	command := protobuf.Command{
		Type: protobuf.Command_UNADVERTISE,
		Unadvertise: &protobuf.Command_Unadvertise{
			Topic: publisher.Topic,
		},
	}
	return publisher.Node.Send(&command)
	
}
