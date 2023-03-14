package ros_hybrid_go

import (
	"sync"

	"github.com/HWZen/ros_hybrid_go/pkg/protobuf"
	"github.com/HWZen/ros_hybrid_go/src/ros_hybrid_error"
	"google.golang.org/protobuf/proto"

	set "github.com/deckarep/golang-set"
)

type ServiceCaller struct {
	node 			*Node
	ServiceName 	string
	TypeName 		string
	SrvSample 		proto.Message
	seq 			uint64
	channel 		chan *protobuf.Command_ResponseService
	reqChanPool 	sync.Pool
	seqReqChanMap 	map[uint64]chan *protobuf.Command_ResponseService

}

func NewServiceCaller(node *Node, serviceName string, typeName string, srvSample proto.Message) *ServiceCaller {
	serviceCaller := &ServiceCaller{
		node: node,
		ServiceName: serviceName,
		TypeName: typeName,
		SrvSample: srvSample,
		channel: make(chan *protobuf.Command_ResponseService),
		reqChanPool: sync.Pool{
			New: func() interface{} {
				return make(chan *protobuf.Command_ResponseService)
			},
		},
		seqReqChanMap: make(map[uint64]chan *protobuf.Command_ResponseService),
		
	}
	if _, exist := node.service_client_channels[serviceName]; !exist {
		node.service_client_channels[serviceName] = set.NewSet()
	}
	node.service_client_channels[serviceName].Add(&serviceCaller.channel)
	go func (caller *ServiceCaller) {
		for responseMsg, ok := <-caller.channel; ok; responseMsg, ok = <-caller.channel {
			callChannel, exist := caller.seqReqChanMap[responseMsg.Seq]
			if !exist {
				continue
			}
			callChannel <- responseMsg
		}
	}(serviceCaller)
	return serviceCaller
}

func (caller *ServiceCaller) Call(req proto.Message) (proto.Message, error) {
	dataBuf, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	command := protobuf.Command{
		Type: protobuf.Command_CALL_SERVICE,
		CallService: &protobuf.Command_CallService{
			Service: caller.ServiceName,
			Type: caller.TypeName,
			Seq: caller.seq,
			Data: dataBuf,
		},
	}
	callChannel := caller.reqChanPool.Get().(chan *protobuf.Command_ResponseService)
	caller.seqReqChanMap[caller.seq] = callChannel
	defer func(caller *ServiceCaller) {
		delete(caller.seqReqChanMap, caller.seq)
		caller.reqChanPool.Put(callChannel)
	}(caller)
	caller.seq++
	err = caller.node.Send(&command)
	if err != nil {
		return nil, err
	}
	responseMsg, ok := <-callChannel
	if !ok {
		return nil, ros_hybrid_go.NewError("service caller channel closed")
	}
	if !responseMsg.Success {
		return nil, ros_hybrid_go.NewError(responseMsg.GetErrorMessage())
	}
	response := proto.Clone(caller.SrvSample)
	err = proto.Unmarshal(responseMsg.Data, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (caller *ServiceCaller) AsyncCall(req proto.Message, callback func(proto.Message, error)) {
	go func() {
		response, err := caller.Call(req)
		callback(response, err)
	}()
}

func (caller *ServiceCaller) Close() {
	if _, exist := caller.node.service_client_channels[caller.ServiceName]; exist {
		delete(caller.node.service_client_channels, caller.ServiceName)
		close(caller.channel)
	}
}

