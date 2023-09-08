// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: oneof.proto

package one_of

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

// The request message containing the user's name.
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[0]
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
	return file_oneof_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ReplyType:
	//
	//	*Reply_Reply1
	//	*Reply_Reply2
	ReplyType isReply_ReplyType `protobuf_oneof:"replyType"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[1]
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
	return file_oneof_proto_rawDescGZIP(), []int{1}
}

func (m *Reply) GetReplyType() isReply_ReplyType {
	if m != nil {
		return m.ReplyType
	}
	return nil
}

func (x *Reply) GetReply1() *Reply1 {
	if x, ok := x.GetReplyType().(*Reply_Reply1); ok {
		return x.Reply1
	}
	return nil
}

func (x *Reply) GetReply2() *Reply2 {
	if x, ok := x.GetReplyType().(*Reply_Reply2); ok {
		return x.Reply2
	}
	return nil
}

type isReply_ReplyType interface {
	isReply_ReplyType()
}

type Reply_Reply1 struct {
	Reply1 *Reply1 `protobuf:"bytes,1,opt,name=reply1,proto3,oneof"`
}

type Reply_Reply2 struct {
	Reply2 *Reply2 `protobuf:"bytes,2,opt,name=reply2,proto3,oneof"`
}

func (*Reply_Reply1) isReply_ReplyType() {}

func (*Reply_Reply2) isReply_ReplyType() {}

// usual response type
type Reply1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ReturnCode int32  `protobuf:"varint,2,opt,name=return_code,json=returnCode,proto3" json:"return_code,omitempty"`
}

func (x *Reply1) Reset() {
	*x = Reply1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply1) ProtoMessage() {}

func (x *Reply1) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply1.ProtoReflect.Descriptor instead.
func (*Reply1) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{2}
}

func (x *Reply1) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Reply1) GetReturnCode() int32 {
	if x != nil {
		return x.ReturnCode
	}
	return 0
}

// other response type
type Reply2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Other int32 `protobuf:"varint,1,opt,name=other,proto3" json:"other,omitempty"`
}

func (x *Reply2) Reset() {
	*x = Reply2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply2) ProtoMessage() {}

func (x *Reply2) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply2.ProtoReflect.Descriptor instead.
func (*Reply2) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{3}
}

func (x *Reply2) GetOther() int32 {
	if x != nil {
		return x.Other
	}
	return 0
}

var File_oneof_proto protoreflect.FileDescriptor

var file_oneof_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x06,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x31, 0x48, 0x00, 0x52, 0x06, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x31, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x42, 0x0b,
	0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x43, 0x0a, 0x06, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x1e, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x32, 0x34, 0x0a, 0x08, 0x47, 0x72, 0x69, 0x70, 0x6d, 0x6f, 0x63, 0x6b, 0x12, 0x28, 0x0a, 0x08,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0e, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x67,
	0x72, 0x69, 0x70, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6f, 0x6e, 0x65, 0x2d, 0x6f, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oneof_proto_rawDescOnce sync.Once
	file_oneof_proto_rawDescData = file_oneof_proto_rawDesc
)

func file_oneof_proto_rawDescGZIP() []byte {
	file_oneof_proto_rawDescOnce.Do(func() {
		file_oneof_proto_rawDescData = protoimpl.X.CompressGZIP(file_oneof_proto_rawDescData)
	})
	return file_oneof_proto_rawDescData
}

var file_oneof_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_oneof_proto_goTypes = []interface{}{
	(*Request)(nil), // 0: oneof.Request
	(*Reply)(nil),   // 1: oneof.Reply
	(*Reply1)(nil),  // 2: oneof.Reply1
	(*Reply2)(nil),  // 3: oneof.Reply2
}
var file_oneof_proto_depIdxs = []int32{
	2, // 0: oneof.Reply.reply1:type_name -> oneof.Reply1
	3, // 1: oneof.Reply.reply2:type_name -> oneof.Reply2
	0, // 2: oneof.Gripmock.SayHello:input_type -> oneof.Request
	1, // 3: oneof.Gripmock.SayHello:output_type -> oneof.Reply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_oneof_proto_init() }
func file_oneof_proto_init() {
	if File_oneof_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oneof_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_oneof_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_oneof_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply1); i {
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
		file_oneof_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply2); i {
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
	file_oneof_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Reply_Reply1)(nil),
		(*Reply_Reply2)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oneof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oneof_proto_goTypes,
		DependencyIndexes: file_oneof_proto_depIdxs,
		MessageInfos:      file_oneof_proto_msgTypes,
	}.Build()
	File_oneof_proto = out.File
	file_oneof_proto_rawDesc = nil
	file_oneof_proto_goTypes = nil
	file_oneof_proto_depIdxs = nil
}
