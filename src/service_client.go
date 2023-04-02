package ros_hybrid_go

import (

	"github.com/HWZen/ros_hybrid_go/pkg/protobuf"
	ros_hybrid_go "github.com/HWZen/ros_hybrid_go/src/ros_hybrid_error"
	"google.golang.org/protobuf/proto"

)

type ServiceCaller struct {
	node          *Node
	ServiceName   string
	TypeName      string
	SrvSample     proto.Message
	seq           uint64
	reqHandle     map[uint64]func (*protobuf.Command_ResponseService)
}


func NewServiceCaller(node *Node, serviceName string, typeName string, srvSample proto.Message) *ServiceCaller {
	if caller, exist := node.serviceClientObservers[serviceName]; exist{
		return caller.(*ServiceCaller)
	}
	serviceCaller := &ServiceCaller{
		node:        node,
		ServiceName: serviceName,
		TypeName:    typeName,
		SrvSample:   srvSample,
		seq:         0,
		reqHandle:   make(map[uint64]func (*protobuf.Command_ResponseService)),
	}
	node.serviceClientObservers[serviceName] = serviceCaller
	return serviceCaller
}

func (caller *ServiceCaller) Update(resMsg proto.Message) {
	response := resMsg.(*protobuf.Command_ResponseService)
	if response.GetService() != caller.ServiceName {
		return
	}
	if handle := caller.reqHandle[response.GetSeq()]; handle != nil {
		handle(response)
		caller.reqHandle[response.GetSeq()] = nil
	}
}

func (caller *ServiceCaller) Shutdown() {

}

func (caller *ServiceCaller) AsyncCall(req proto.Message, callback func(proto.Message, error)) error{
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	caller.seq++
	command := protobuf.Command{
		Type: protobuf.Command_CALL_SERVICE,
		CallService: &protobuf.Command_CallService{
			Service: caller.ServiceName,
			Type:    caller.TypeName,
			Seq:     caller.seq,
			Data: 	 data,
		},
	}
	err = caller.node.Send(&command)
	if err != nil {
		return err
	}
	caller.reqHandle[caller.seq] = func (res *protobuf.Command_ResponseService) {
		if !res.GetSuccess() {
			callback(nil, ros_hybrid_go.NewError("service call failed: " + res.GetErrorMessage()))
		}
		callbackData := proto.Clone(caller.SrvSample)
		err := proto.Unmarshal(res.GetData(), callbackData)
		if err != nil {
			callback(nil, err)
		}
		
		callback(callbackData, nil)
	}
	return nil
}

func (caller *ServiceCaller) Call(req proto.Message) (proto.Message, error) {
	getCallbackFunc := func(res *proto.Message, err *error, resultCh chan struct{}) func (proto.Message, error) {
		return func (callbackRes proto.Message, callbackErr error) {
			*res = callbackRes
			*err = callbackErr
			close(resultCh)
		}
	}
	var res proto.Message
	var err error
	result := make(chan struct{})

	caller.AsyncCall(req, getCallbackFunc(&res, &err, result))
	if err != nil {
		return nil, err
	}

	<-result

	return res, nil
}
