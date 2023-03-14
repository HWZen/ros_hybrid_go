// generated automatically by ros_hybrid_protoc on  Sun Feb 19 23:34:16 2023
// Do not Edit!
// wrapping message: ros_hybrid_sdk/MyService

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: srv/ros_hybrid_sdk/MyService.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MyService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *MyService_Request  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response *MyService_Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *MyService) Reset() {
	*x = MyService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_srv_ros_hybrid_sdk_MyService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyService) ProtoMessage() {}

func (x *MyService) ProtoReflect() protoreflect.Message {
	mi := &file_srv_ros_hybrid_sdk_MyService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyService.ProtoReflect.Descriptor instead.
func (*MyService) Descriptor() ([]byte, []int) {
	return file_srv_ros_hybrid_sdk_MyService_proto_rawDescGZIP(), []int{0}
}

func (x *MyService) GetRequest() *MyService_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *MyService) GetResponse() *MyService_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type MyService_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyMsg *MyMsg `protobuf:"bytes,1,opt,name=myMsg,proto3" json:"myMsg,omitempty"`
}

func (x *MyService_Request) Reset() {
	*x = MyService_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_srv_ros_hybrid_sdk_MyService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyService_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyService_Request) ProtoMessage() {}

func (x *MyService_Request) ProtoReflect() protoreflect.Message {
	mi := &file_srv_ros_hybrid_sdk_MyService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyService_Request.ProtoReflect.Descriptor instead.
func (*MyService_Request) Descriptor() ([]byte, []int) {
	return file_srv_ros_hybrid_sdk_MyService_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MyService_Request) GetMyMsg() *MyMsg {
	if x != nil {
		return x.MyMsg
	}
	return nil
}

type MyService_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *MyService_Response) Reset() {
	*x = MyService_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_srv_ros_hybrid_sdk_MyService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyService_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyService_Response) ProtoMessage() {}

func (x *MyService_Response) ProtoReflect() protoreflect.Message {
	mi := &file_srv_ros_hybrid_sdk_MyService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyService_Response.ProtoReflect.Descriptor instead.
func (*MyService_Response) Descriptor() ([]byte, []int) {
	return file_srv_ros_hybrid_sdk_MyService_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MyService_Response) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_srv_ros_hybrid_sdk_MyService_proto protoreflect.FileDescriptor

var file_srv_ros_hybrid_sdk_MyService_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x72, 0x76, 0x2f, 0x72, 0x6f, 0x73, 0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64,
	0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x72, 0x6f, 0x73,
	0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x5f, 0x73, 0x64, 0x6b, 0x1a, 0x1a, 0x6d, 0x73, 0x67,
	0x73, 0x2f, 0x73, 0x74, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2f, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6d, 0x73, 0x67, 0x73, 0x2f, 0x72, 0x6f,
	0x73, 0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x5f, 0x73, 0x64, 0x6b, 0x2f, 0x4d, 0x79, 0x4d,
	0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x09, 0x4d, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64,
	0x2e, 0x72, 0x6f, 0x73, 0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x5f, 0x73, 0x64, 0x6b, 0x2e,
	0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68,
	0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x72, 0x6f, 0x73, 0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64,
	0x5f, 0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x1a, 0x3d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05,
	0x6d, 0x79, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x68, 0x79,
	0x62, 0x72, 0x69, 0x64, 0x2e, 0x72, 0x6f, 0x73, 0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x5f,
	0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x79, 0x4d, 0x73, 0x67, 0x52, 0x05, 0x6d, 0x79, 0x4d, 0x73, 0x67,
	0x1a, 0x3b, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68,
	0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x73, 0x74, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_srv_ros_hybrid_sdk_MyService_proto_rawDescOnce sync.Once
	file_srv_ros_hybrid_sdk_MyService_proto_rawDescData = file_srv_ros_hybrid_sdk_MyService_proto_rawDesc
)

func file_srv_ros_hybrid_sdk_MyService_proto_rawDescGZIP() []byte {
	file_srv_ros_hybrid_sdk_MyService_proto_rawDescOnce.Do(func() {
		file_srv_ros_hybrid_sdk_MyService_proto_rawDescData = protoimpl.X.CompressGZIP(file_srv_ros_hybrid_sdk_MyService_proto_rawDescData)
	})
	return file_srv_ros_hybrid_sdk_MyService_proto_rawDescData
}

var file_srv_ros_hybrid_sdk_MyService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_srv_ros_hybrid_sdk_MyService_proto_goTypes = []interface{}{
	(*MyService)(nil),          // 0: hybrid.ros_hybrid_sdk.MyService
	(*MyService_Request)(nil),  // 1: hybrid.ros_hybrid_sdk.MyService.Request
	(*MyService_Response)(nil), // 2: hybrid.ros_hybrid_sdk.MyService.Response
	(*MyMsg)(nil),              // 3: hybrid.ros_hybrid_sdk.MyMsg
	(*Header)(nil),             // 4: hybrid.std_msgs.Header
}
var file_srv_ros_hybrid_sdk_MyService_proto_depIdxs = []int32{
	1, // 0: hybrid.ros_hybrid_sdk.MyService.request:type_name -> hybrid.ros_hybrid_sdk.MyService.Request
	2, // 1: hybrid.ros_hybrid_sdk.MyService.response:type_name -> hybrid.ros_hybrid_sdk.MyService.Response
	3, // 2: hybrid.ros_hybrid_sdk.MyService.Request.myMsg:type_name -> hybrid.ros_hybrid_sdk.MyMsg
	4, // 3: hybrid.ros_hybrid_sdk.MyService.Response.header:type_name -> hybrid.std_msgs.Header
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_srv_ros_hybrid_sdk_MyService_proto_init() }
func file_srv_ros_hybrid_sdk_MyService_proto_init() {
	if File_srv_ros_hybrid_sdk_MyService_proto != nil {
		return
	}
	file_msgs_std_msgs_Header_proto_init()
	file_msgs_ros_hybrid_sdk_MyMsg_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_srv_ros_hybrid_sdk_MyService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyService); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_srv_ros_hybrid_sdk_MyService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyService_Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_srv_ros_hybrid_sdk_MyService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyService_Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_srv_ros_hybrid_sdk_MyService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_srv_ros_hybrid_sdk_MyService_proto_goTypes,
		DependencyIndexes: file_srv_ros_hybrid_sdk_MyService_proto_depIdxs,
		MessageInfos:      file_srv_ros_hybrid_sdk_MyService_proto_msgTypes,
	}.Build()
	File_srv_ros_hybrid_sdk_MyService_proto = out.File
	file_srv_ros_hybrid_sdk_MyService_proto_rawDesc = nil
	file_srv_ros_hybrid_sdk_MyService_proto_goTypes = nil
	file_srv_ros_hybrid_sdk_MyService_proto_depIdxs = nil
}
