// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: idl/hello.proto

package echo

import (
	context "context"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_idl_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_idl_hello_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_idl_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_idl_hello_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_idl_hello_proto protoreflect.FileDescriptor

var file_idl_hello_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x64, 0x6c, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x22, 0x1b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x1c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x32, 0xbb, 0x01, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x38, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x65, 0x63, 0x68, 0x6f,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x38, 0x0a, 0x13,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x11, 0x42, 0x69, 0x64, 0x69, 0x53, 0x69,
	0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x65, 0x63,
	0x68, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x65, 0x63, 0x68,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x65, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x2d, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x6b, 0x69,
	0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_hello_proto_rawDescOnce sync.Once
	file_idl_hello_proto_rawDescData = file_idl_hello_proto_rawDesc
)

func file_idl_hello_proto_rawDescGZIP() []byte {
	file_idl_hello_proto_rawDescOnce.Do(func() {
		file_idl_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_hello_proto_rawDescData)
	})
	return file_idl_hello_proto_rawDescData
}

var file_idl_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_idl_hello_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: echo.Request
	(*Response)(nil), // 1: echo.Response
}
var file_idl_hello_proto_depIdxs = []int32{
	0, // 0: echo.EchoService.ClientSideStreaming:input_type -> echo.Request
	0, // 1: echo.EchoService.ServerSideStreaming:input_type -> echo.Request
	0, // 2: echo.EchoService.BidiSideStreaming:input_type -> echo.Request
	1, // 3: echo.EchoService.ClientSideStreaming:output_type -> echo.Response
	1, // 4: echo.EchoService.ServerSideStreaming:output_type -> echo.Response
	1, // 5: echo.EchoService.BidiSideStreaming:output_type -> echo.Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_hello_proto_init() }
func file_idl_hello_proto_init() {
	if File_idl_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_idl_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_idl_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_hello_proto_goTypes,
		DependencyIndexes: file_idl_hello_proto_depIdxs,
		MessageInfos:      file_idl_hello_proto_msgTypes,
	}.Build()
	File_idl_hello_proto = out.File
	file_idl_hello_proto_rawDesc = nil
	file_idl_hello_proto_goTypes = nil
	file_idl_hello_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.0.3. DO NOT EDIT.

type EchoService interface {
	ClientSideStreaming(stream EchoService_ClientSideStreamingServer) (err error)
	ServerSideStreaming(req *Request, stream EchoService_ServerSideStreamingServer) (err error)
	BidiSideStreaming(stream EchoService_BidiSideStreamingServer) (err error)
}

type EchoService_ClientSideStreamingServer interface {
	streaming.Stream
	Recv() (*Request, error)
	SendAndClose(*Response) error
}

type EchoService_ServerSideStreamingServer interface {
	streaming.Stream
	Send(*Response) error
}

type EchoService_BidiSideStreamingServer interface {
	streaming.Stream
	Recv() (*Request, error)
	Send(*Response) error
}
