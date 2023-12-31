// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pb/helloworld.proto

package pb

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

// The request message containing a name(Client-side)
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_pb_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing a message(Server-side)
type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_pb_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FactorialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *FactorialRequest) Reset() {
	*x = FactorialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FactorialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FactorialRequest) ProtoMessage() {}

func (x *FactorialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FactorialRequest.ProtoReflect.Descriptor instead.
func (*FactorialRequest) Descriptor() ([]byte, []int) {
	return file_pb_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *FactorialRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FactorialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *FactorialResponse) Reset() {
	*x = FactorialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FactorialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FactorialResponse) ProtoMessage() {}

func (x *FactorialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FactorialResponse.ProtoReflect.Descriptor instead.
func (*FactorialResponse) Descriptor() ([]byte, []int) {
	return file_pb_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *FactorialResponse) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_pb_helloworld_proto protoreflect.FileDescriptor

var file_pb_helloworld_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a,
	0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x11, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x32, 0x7b, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12,
	0x31, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x12,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31,
	0x5a, 0x2f, 0x6c, 0x65, 0x74, 0x73, 0x2d, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2d, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x2d, 0x30, 0x30, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_helloworld_proto_rawDescOnce sync.Once
	file_pb_helloworld_proto_rawDescData = file_pb_helloworld_proto_rawDesc
)

func file_pb_helloworld_proto_rawDescGZIP() []byte {
	file_pb_helloworld_proto_rawDescOnce.Do(func() {
		file_pb_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_helloworld_proto_rawDescData)
	})
	return file_pb_helloworld_proto_rawDescData
}

var file_pb_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_helloworld_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),      // 0: pb.HelloRequest
	(*HelloResponse)(nil),     // 1: pb.HelloResponse
	(*FactorialRequest)(nil),  // 2: pb.FactorialRequest
	(*FactorialResponse)(nil), // 3: pb.FactorialResponse
}
var file_pb_helloworld_proto_depIdxs = []int32{
	0, // 0: pb.HelloWorld.SayHello:input_type -> pb.HelloRequest
	2, // 1: pb.HelloWorld.Factorial:input_type -> pb.FactorialRequest
	1, // 2: pb.HelloWorld.SayHello:output_type -> pb.HelloResponse
	3, // 3: pb.HelloWorld.Factorial:output_type -> pb.FactorialResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_helloworld_proto_init() }
func file_pb_helloworld_proto_init() {
	if File_pb_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_pb_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_pb_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FactorialRequest); i {
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
		file_pb_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FactorialResponse); i {
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
			RawDescriptor: file_pb_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_helloworld_proto_goTypes,
		DependencyIndexes: file_pb_helloworld_proto_depIdxs,
		MessageInfos:      file_pb_helloworld_proto_msgTypes,
	}.Build()
	File_pb_helloworld_proto = out.File
	file_pb_helloworld_proto_rawDesc = nil
	file_pb_helloworld_proto_goTypes = nil
	file_pb_helloworld_proto_depIdxs = nil
}
