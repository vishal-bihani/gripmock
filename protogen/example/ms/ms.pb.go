// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.24.4
// source: ms.proto

package ms

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V1 [][]byte `protobuf:"bytes,1,rep,name=v1,proto3" json:"v1,omitempty"`
	V2 []string `protobuf:"bytes,2,rep,name=v2,proto3" json:"v2,omitempty"`
	V3 *int64   `protobuf:"varint,3,opt,name=v3,proto3,oneof" json:"v3,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_ms_proto_msgTypes[0]
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
	return file_ms_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetV1() [][]byte {
	if x != nil {
		return x.V1
	}
	return nil
}

func (x *Request) GetV2() []string {
	if x != nil {
		return x.V2
	}
	return nil
}

func (x *Request) GetV3() int64 {
	if x != nil && x.V3 != nil {
		return *x.V3
	}
	return 0
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	V3   *uint64 `protobuf:"varint,2,opt,name=v3,proto3,oneof" json:"v3,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_ms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_ms_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Reply) GetV3() uint64 {
	if x != nil && x.V3 != nil {
		return *x.V3
	}
	return 0
}

var File_ms_proto protoreflect.FileDescriptor

var file_ms_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x6d, 0x73, 0x22, 0x45,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x31, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x02, 0x76, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x32, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x76, 0x32, 0x12, 0x13, 0x0a, 0x02, 0x76, 0x33, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x76, 0x33, 0x88, 0x01, 0x01, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x76, 0x33, 0x22, 0x37, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x76, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00,
	0x52, 0x02, 0x76, 0x33, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x76, 0x33, 0x32, 0x32,
	0x0a, 0x0c, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22,
	0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0b, 0x2e, 0x6d, 0x73, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x61, 0x76, 0x69, 0x78, 0x2f, 0x67, 0x72, 0x69, 0x70, 0x6d, 0x6f, 0x63, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ms_proto_rawDescOnce sync.Once
	file_ms_proto_rawDescData = file_ms_proto_rawDesc
)

func file_ms_proto_rawDescGZIP() []byte {
	file_ms_proto_rawDescOnce.Do(func() {
		file_ms_proto_rawDescData = protoimpl.X.CompressGZIP(file_ms_proto_rawDescData)
	})
	return file_ms_proto_rawDescData
}

var file_ms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ms_proto_goTypes = []interface{}{
	(*Request)(nil), // 0: ms.Request
	(*Reply)(nil),   // 1: ms.Reply
}
var file_ms_proto_depIdxs = []int32{
	0, // 0: ms.MicroService.SayHello:input_type -> ms.Request
	1, // 1: ms.MicroService.SayHello:output_type -> ms.Reply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ms_proto_init() }
func file_ms_proto_init() {
	if File_ms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_ms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
	file_ms_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_ms_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ms_proto_goTypes,
		DependencyIndexes: file_ms_proto_depIdxs,
		MessageInfos:      file_ms_proto_msgTypes,
	}.Build()
	File_ms_proto = out.File
	file_ms_proto_rawDesc = nil
	file_ms_proto_goTypes = nil
	file_ms_proto_depIdxs = nil
}
