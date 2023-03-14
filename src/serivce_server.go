package ros_hybrid_go

import (
	"github.com/HWZen/ros_hybrid_go/pkg/protobuf"
	ros_hybrid_go "github.com/HWZen/ros_hybrid_go/src/ros_hybrid_error"
	"google.golang.org/protobuf/proto"
)

type ServiceCallbackFunc func(req proto.Message) proto.Message

type ServicerServer struct {
	node 			*Node
	ServiceName 	string
	TypeName 		string
	MsgSample 		proto.Message
	callbackFunc 	ServiceCallbackFunc
	channel 		chan *protobuf.Command_CallService
}

func NewServicer(node *Node, serviceName string, typeName string, msgSample proto.Message, callbackFunc ServiceCallbackFunc) (*ServicerServer, error) {
	if _, exist := node.service_server_channels[serviceName]; exist {
		// service already exist
		// log it 
		return nil, ros_hybrid_go.NewError("service already exist")
	}
	command := protobuf.Command{
		Type: protobuf.Command_ADVERTISE_SERVICE,
		AdvertiseService: &protobuf.Command_AdvertiseService{
			Service: serviceName,
			Type: typeName,
		},
	}

	err := node.Send(&command)
	if err != nil {
		return nil, err
	}
	
	servicer := ServicerServer{
		node: node,
		ServiceName: serviceName,
		TypeName: typeName,
		MsgSample: msgSample,
		callbackFunc: callbackFunc,
		channel: make(chan *protobuf.Command_CallService),
	}
	
	node.service_server_channels[serviceName] = &servicer.channel

	return &servicer, nil
}

func (servicer *ServicerServer) Advertise() {
	for callMsg, ok := <-servicer.channel; ok; callMsg, ok = <-servicer.channel {
		go func(callMsg *protobuf.Command_CallService) {
			response_service := &protobuf.Command_ResponseService{
				Service: callMsg.Service,
				Success: false,
				Seq: callMsg.Seq,
				ErrorMessage: proto.String("no set error message"),
			}
			command_response := protobuf.Command{
				Type: protobuf.Command_RESPONSE_SERVICE,
				ResponseService: response_service,
			}
			defer servicer.node.Send(&command_response)

			req := proto.Clone(servicer.MsgSample)
			err := proto.Unmarshal(callMsg.GetData(), req)
			if err != nil {
				// log it
				response_service.ErrorMessage = proto.String("proto.Unmarshal error: " + err.Error())
				return
			}

			res := servicer.callbackFunc(req)
			resBytes, err := proto.Marshal(res)
			if err != nil {
				// log it
				response_service.ErrorMessage = proto.String("proto.Marshal error: " + err.Error())
				return
			}

			response_service.Success = true
			response_service.Data = resBytes
			response_service.ErrorMessage = nil
		}(callMsg)
	}
}

func (servicer *ServicerServer) Unadvertise() {
	command := protobuf.Command{
		Type: protobuf.Command_UNADVERTISE_SERVICE,
		UnadvertiseService: &protobuf.Command_UnadvertiseService{
			Service: servicer.ServiceName,
		},
	}
	err := servicer.node.Send(&command)
	if err != nil {
		// log it
		return
	}

	delete(servicer.node.service_server_channels, servicer.ServiceName)
	close(servicer.channel)
}

