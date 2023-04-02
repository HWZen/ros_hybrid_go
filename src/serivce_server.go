package ros_hybrid_go

import (
	"github.com/HWZen/ros_hybrid_go/pkg/protobuf"
	ros_hybrid_go "github.com/HWZen/ros_hybrid_go/src/ros_hybrid_error"
	"google.golang.org/protobuf/proto"
)

type ServiceCallbackFunc func(req proto.Message) proto.Message

type ServicerServer struct {
	node         *Node
	ServiceName  string
	TypeName     string
	MsgSample    proto.Message
	callbackFunc ServiceCallbackFunc
}


func (servicer *ServicerServer) Advertise() error{
	if _, exist := servicer.node.serviceServerObservers[servicer.ServiceName]; exist {
		// service already exist
		// log it
		return ros_hybrid_go.NewError("service already exist")
	}
	command := protobuf.Command{
		Type: protobuf.Command_ADVERTISE_SERVICE,
		AdvertiseService: &protobuf.Command_AdvertiseService{
			Service: servicer.ServiceName,
			Type:    servicer.TypeName,
		},
	}

	err := servicer.node.Send(&command)
	if err != nil {
		return err
	}
	servicer.node.serviceServerObservers[servicer.ServiceName] = servicer
	return nil;
}

func NewServicer(node *Node, serviceName string, typeName string, msgSample proto.Message, callbackFunc ServiceCallbackFunc) (*ServicerServer, error) {
	servicer := ServicerServer{
		node:         node,
		ServiceName:  serviceName,
		TypeName:     typeName,
		MsgSample:    msgSample,
		callbackFunc: callbackFunc,
	}

	err := servicer.Advertise()

	if err != nil {
		return nil, err
	}

	return &servicer, nil
}

func (servicer *ServicerServer) Update(call proto.Message) {
	callMsg := call.(*protobuf.Command_CallService)
	response_service := &protobuf.Command_ResponseService{
		Service:      callMsg.Service,
		Success:      false,
		Seq:          callMsg.Seq,
		ErrorMessage: proto.String("no set error message"),
	}
	command_response := protobuf.Command{
		Type:            protobuf.Command_RESPONSE_SERVICE,
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
}

func (servicer *ServicerServer) Shutdown() {
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
}


func (servicer *ServicerServer) Unadvertise() error {
	servicer.Shutdown()
	delete(servicer.node.serviceServerObservers, servicer.ServiceName)
	return ros_hybrid_go.NewError("servicer not found in node")
}
