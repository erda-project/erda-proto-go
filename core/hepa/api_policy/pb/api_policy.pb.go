// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: api_policy.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *structpb.Value `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SetPolicyResponse) Reset() {
	*x = SetPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPolicyResponse) ProtoMessage() {}

func (x *SetPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPolicyResponse.ProtoReflect.Descriptor instead.
func (*SetPolicyResponse) Descriptor() ([]byte, []int) {
	return file_api_policy_proto_rawDescGZIP(), []int{0}
}

func (x *SetPolicyResponse) GetData() *structpb.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

type SetPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category  string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	PackageId string `protobuf:"bytes,2,opt,name=packageId,proto3" json:"packageId,omitempty"`
	ApiId     string `protobuf:"bytes,3,opt,name=apiId,proto3" json:"apiId,omitempty"`
	Body      []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *SetPolicyRequest) Reset() {
	*x = SetPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPolicyRequest) ProtoMessage() {}

func (x *SetPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPolicyRequest.ProtoReflect.Descriptor instead.
func (*SetPolicyRequest) Descriptor() ([]byte, []int) {
	return file_api_policy_proto_rawDescGZIP(), []int{1}
}

func (x *SetPolicyRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *SetPolicyRequest) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *SetPolicyRequest) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

func (x *SetPolicyRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type GetPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category  string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	PackageId string `protobuf:"bytes,2,opt,name=packageId,proto3" json:"packageId,omitempty"`
	ApiId     string `protobuf:"bytes,3,opt,name=apiId,proto3" json:"apiId,omitempty"`
}

func (x *GetPolicyRequest) Reset() {
	*x = GetPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyRequest) ProtoMessage() {}

func (x *GetPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyRequest.ProtoReflect.Descriptor instead.
func (*GetPolicyRequest) Descriptor() ([]byte, []int) {
	return file_api_policy_proto_rawDescGZIP(), []int{2}
}

func (x *GetPolicyRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *GetPolicyRequest) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *GetPolicyRequest) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

type GetPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *structpb.Value `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetPolicyResponse) Reset() {
	*x = GetPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyResponse) ProtoMessage() {}

func (x *GetPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyResponse.ProtoReflect.Descriptor instead.
func (*GetPolicyResponse) Descriptor() ([]byte, []int) {
	return file_api_policy_proto_rawDescGZIP(), []int{3}
}

func (x *GetPolicyResponse) GetData() *structpb.Value {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_policy_proto protoreflect.FileDescriptor

var file_api_policy_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65,
	0x70, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x11, 0x53, 0x65, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x76, 0x0a, 0x10, 0x53, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x69, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0x62, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x86, 0x03, 0x0a, 0x10, 0x41, 0x70, 0x69, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb4, 0x01, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2b, 0x2e, 0x65, 0x72, 0x64,
	0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x46, 0x12, 0x44, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x7d, 0x3f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x3d, 0x7b, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x7d, 0x26, 0x61, 0x70, 0x69, 0x49, 0x64, 0x3d, 0x7b, 0x61, 0x70, 0x69,
	0x49, 0x64, 0x7d, 0x12, 0xba, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x2b, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65,
	0x70, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x68, 0x65, 0x70, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x4c, 0x1a, 0x44, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x7d, 0x3f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x3d, 0x7b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x7d, 0x26, 0x61, 0x70, 0x69,
	0x49, 0x64, 0x3d, 0x7b, 0x61, 0x70, 0x69, 0x49, 0x64, 0x7d, 0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68,
	0x65, 0x70, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_policy_proto_rawDescOnce sync.Once
	file_api_policy_proto_rawDescData = file_api_policy_proto_rawDesc
)

func file_api_policy_proto_rawDescGZIP() []byte {
	file_api_policy_proto_rawDescOnce.Do(func() {
		file_api_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_policy_proto_rawDescData)
	})
	return file_api_policy_proto_rawDescData
}

var file_api_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_policy_proto_goTypes = []interface{}{
	(*SetPolicyResponse)(nil), // 0: erda.core.hepa.api_policy.SetPolicyResponse
	(*SetPolicyRequest)(nil),  // 1: erda.core.hepa.api_policy.SetPolicyRequest
	(*GetPolicyRequest)(nil),  // 2: erda.core.hepa.api_policy.GetPolicyRequest
	(*GetPolicyResponse)(nil), // 3: erda.core.hepa.api_policy.GetPolicyResponse
	(*structpb.Value)(nil),    // 4: google.protobuf.Value
}
var file_api_policy_proto_depIdxs = []int32{
	4, // 0: erda.core.hepa.api_policy.SetPolicyResponse.data:type_name -> google.protobuf.Value
	4, // 1: erda.core.hepa.api_policy.GetPolicyResponse.data:type_name -> google.protobuf.Value
	2, // 2: erda.core.hepa.api_policy.ApiPolicyService.GetPolicy:input_type -> erda.core.hepa.api_policy.GetPolicyRequest
	1, // 3: erda.core.hepa.api_policy.ApiPolicyService.SetPolicy:input_type -> erda.core.hepa.api_policy.SetPolicyRequest
	3, // 4: erda.core.hepa.api_policy.ApiPolicyService.GetPolicy:output_type -> erda.core.hepa.api_policy.GetPolicyResponse
	0, // 5: erda.core.hepa.api_policy.ApiPolicyService.SetPolicy:output_type -> erda.core.hepa.api_policy.SetPolicyResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_policy_proto_init() }
func file_api_policy_proto_init() {
	if File_api_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPolicyResponse); i {
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
		file_api_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPolicyRequest); i {
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
		file_api_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyRequest); i {
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
		file_api_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyResponse); i {
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
			RawDescriptor: file_api_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_policy_proto_goTypes,
		DependencyIndexes: file_api_policy_proto_depIdxs,
		MessageInfos:      file_api_policy_proto_msgTypes,
	}.Build()
	File_api_policy_proto = out.File
	file_api_policy_proto_rawDesc = nil
	file_api_policy_proto_goTypes = nil
	file_api_policy_proto_depIdxs = nil
}
