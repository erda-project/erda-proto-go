// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: dingtalk_test.proto

package pb

import (
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type DingTalkTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Webhook string `protobuf:"bytes,1,opt,name=webhook,proto3" json:"webhook,omitempty"`
	Secret  string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *DingTalkTestRequest) Reset() {
	*x = DingTalkTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dingtalk_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DingTalkTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DingTalkTestRequest) ProtoMessage() {}

func (x *DingTalkTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dingtalk_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DingTalkTestRequest.ProtoReflect.Descriptor instead.
func (*DingTalkTestRequest) Descriptor() ([]byte, []int) {
	return file_dingtalk_test_proto_rawDescGZIP(), []int{0}
}

func (x *DingTalkTestRequest) GetWebhook() string {
	if x != nil {
		return x.Webhook
	}
	return ""
}

func (x *DingTalkTestRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type DingTalkTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DingTalkTestResponse) Reset() {
	*x = DingTalkTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dingtalk_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DingTalkTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DingTalkTestResponse) ProtoMessage() {}

func (x *DingTalkTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dingtalk_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DingTalkTestResponse.ProtoReflect.Descriptor instead.
func (*DingTalkTestResponse) Descriptor() ([]byte, []int) {
	return file_dingtalk_test_proto_rawDescGZIP(), []int{1}
}

func (x *DingTalkTestResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DingTalkTestResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_dingtalk_test_proto protoreflect.FileDescriptor

var file_dingtalk_test_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x69, 0x6e, 0x67, 0x74, 0x61, 0x6c, 0x6b, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74,
	0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x13, 0x44, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x6c, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x12, 0x1e, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x22, 0x46, 0x0a, 0x14, 0x44, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c, 0x6b, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x94, 0x01, 0x0a, 0x13, 0x44, 0x69, 0x6e,
	0x67, 0x54, 0x61, 0x6c, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x7d, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x44, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x44, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2f, 0x64, 0x69, 0x6e, 0x67, 0x74, 0x61, 0x6c, 0x6b, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72,
	0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dingtalk_test_proto_rawDescOnce sync.Once
	file_dingtalk_test_proto_rawDescData = file_dingtalk_test_proto_rawDesc
)

func file_dingtalk_test_proto_rawDescGZIP() []byte {
	file_dingtalk_test_proto_rawDescOnce.Do(func() {
		file_dingtalk_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_dingtalk_test_proto_rawDescData)
	})
	return file_dingtalk_test_proto_rawDescData
}

var file_dingtalk_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dingtalk_test_proto_goTypes = []interface{}{
	(*DingTalkTestRequest)(nil),  // 0: erda.admin.DingTalkTestRequest
	(*DingTalkTestResponse)(nil), // 1: erda.admin.DingTalkTestResponse
}
var file_dingtalk_test_proto_depIdxs = []int32{
	0, // 0: erda.admin.DingTalkTestService.SendTestMessage:input_type -> erda.admin.DingTalkTestRequest
	1, // 1: erda.admin.DingTalkTestService.SendTestMessage:output_type -> erda.admin.DingTalkTestResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dingtalk_test_proto_init() }
func file_dingtalk_test_proto_init() {
	if File_dingtalk_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dingtalk_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DingTalkTestRequest); i {
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
		file_dingtalk_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DingTalkTestResponse); i {
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
			RawDescriptor: file_dingtalk_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dingtalk_test_proto_goTypes,
		DependencyIndexes: file_dingtalk_test_proto_depIdxs,
		MessageInfos:      file_dingtalk_test_proto_msgTypes,
	}.Build()
	File_dingtalk_test_proto = out.File
	file_dingtalk_test_proto_rawDesc = nil
	file_dingtalk_test_proto_goTypes = nil
	file_dingtalk_test_proto_depIdxs = nil
}
