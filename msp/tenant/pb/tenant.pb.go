// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: tenant.proto

package pb

import (
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Workspace int32

const (
	Workspace_DEV     Workspace = 0
	Workspace_TEST    Workspace = 1
	Workspace_STAGING Workspace = 2
	Workspace_PROD    Workspace = 3
	Workspace_DEFAULT Workspace = 4
)

// Enum value maps for Workspace.
var (
	Workspace_name = map[int32]string{
		0: "DEV",
		1: "TEST",
		2: "STAGING",
		3: "PROD",
		4: "DEFAULT",
	}
	Workspace_value = map[string]int32{
		"DEV":     0,
		"TEST":    1,
		"STAGING": 2,
		"PROD":    3,
		"DEFAULT": 4,
	}
)

func (x Workspace) Enum() *Workspace {
	p := new(Workspace)
	*p = x
	return p
}

func (x Workspace) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Workspace) Descriptor() protoreflect.EnumDescriptor {
	return file_tenant_proto_enumTypes[0].Descriptor()
}

func (Workspace) Type() protoreflect.EnumType {
	return &file_tenant_proto_enumTypes[0]
}

func (x Workspace) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Workspace.Descriptor instead.
func (Workspace) EnumDescriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{0}
}

type Type int32

const (
	Type_DOP Type = 0
	Type_MSP Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "DOP",
		1: "MSP",
	}
	Type_value = map[string]int32{
		"DOP": 0,
		"MSP": 1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_tenant_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_tenant_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{1}
}

type CreateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID  string   `protobuf:"bytes,1,opt,name=projectID,json=projectId,proto3" json:"projectID,omitempty"`
	TenantType string   `protobuf:"bytes,2,opt,name=tenantType,proto3" json:"tenantType,omitempty"`
	Workspaces []string `protobuf:"bytes,3,rep,name=workspaces,proto3" json:"workspaces,omitempty"`
}

func (x *CreateTenantRequest) Reset() {
	*x = CreateTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantRequest) ProtoMessage() {}

func (x *CreateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantRequest.ProtoReflect.Descriptor instead.
func (*CreateTenantRequest) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTenantRequest) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *CreateTenantRequest) GetTenantType() string {
	if x != nil {
		return x.TenantType
	}
	return ""
}

func (x *CreateTenantRequest) GetWorkspaces() []string {
	if x != nil {
		return x.Workspaces
	}
	return nil
}

type CreateTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Tenant `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateTenantResponse) Reset() {
	*x = CreateTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenantResponse) ProtoMessage() {}

func (x *CreateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenantResponse.ProtoReflect.Descriptor instead.
func (*CreateTenantResponse) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTenantResponse) GetData() []*Tenant {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID  string `protobuf:"bytes,1,opt,name=projectID,json=projectId,proto3" json:"projectID,omitempty"`
	TenantType string `protobuf:"bytes,2,opt,name=tenantType,proto3" json:"tenantType,omitempty"`
	Workspace  string `protobuf:"bytes,3,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *GetTenantRequest) Reset() {
	*x = GetTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequest) ProtoMessage() {}

func (x *GetTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequest.ProtoReflect.Descriptor instead.
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{2}
}

func (x *GetTenantRequest) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *GetTenantRequest) GetTenantType() string {
	if x != nil {
		return x.TenantType
	}
	return ""
}

func (x *GetTenantRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

type GetTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Tenant `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetTenantResponse) Reset() {
	*x = GetTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantResponse) ProtoMessage() {}

func (x *GetTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantResponse.ProtoReflect.Descriptor instead.
func (*GetTenantResponse) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{3}
}

func (x *GetTenantResponse) GetData() *Tenant {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID  string `protobuf:"bytes,1,opt,name=projectID,json=projectId,proto3" json:"projectID,omitempty"`
	TenantType string `protobuf:"bytes,2,opt,name=tenantType,proto3" json:"tenantType,omitempty"`
	Workspace  string `protobuf:"bytes,3,opt,name=workspace,proto3" json:"workspace,omitempty"`
}

func (x *DeleteTenantRequest) Reset() {
	*x = DeleteTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantRequest) ProtoMessage() {}

func (x *DeleteTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantRequest.ProtoReflect.Descriptor instead.
func (*DeleteTenantRequest) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTenantRequest) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *DeleteTenantRequest) GetTenantType() string {
	if x != nil {
		return x.TenantType
	}
	return ""
}

func (x *DeleteTenantRequest) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

type DeleteTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Tenant `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DeleteTenantResponse) Reset() {
	*x = DeleteTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantResponse) ProtoMessage() {}

func (x *DeleteTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantResponse.ProtoReflect.Descriptor instead.
func (*DeleteTenantResponse) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTenantResponse) GetData() *Tenant {
	if x != nil {
		return x.Data
	}
	return nil
}

// Differentiate the actual deployment environment of the service and addon
type Tenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tenant id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Tenant type
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Related project id
	ProjectID string `protobuf:"bytes,3,opt,name=projectID,json=projectId,proto3" json:"projectID,omitempty"`
	// Related deployment workspace
	Workspace string `protobuf:"bytes,4,opt,name=workspace,proto3" json:"workspace,omitempty"`
	// Create timestamp
	CreateTime int64 `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// Update timestamp
	UpdateTime int64 `protobuf:"varint,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	// Delete or not
	IsDeleted bool `protobuf:"varint,7,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{6}
}

func (x *Tenant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tenant) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Tenant) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *Tenant) GetWorkspace() string {
	if x != nil {
		return x.Workspace
	}
	return ""
}

func (x *Tenant) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Tenant) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *Tenant) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

var File_tenant_proto protoreflect.FileDescriptor

var file_tenant_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f,
	0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58,
	0x01, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a,
	0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x72,
	0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58,
	0x01, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a,
	0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52,
	0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x22, 0x43, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc6, 0x01, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2a,
	0x42, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x07, 0x0a, 0x03,
	0x44, 0x45, 0x56, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x50, 0x52, 0x4f, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c,
	0x54, 0x10, 0x04, 0x2a, 0x18, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x44,
	0x4f, 0x50, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x53, 0x50, 0x10, 0x01, 0x32, 0xad, 0x03,
	0x0a, 0x0d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x8a, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x12, 0x24, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x27, 0x22, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x3d, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x7d, 0x12, 0x81, 0x01, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x65, 0x72, 0x64,
	0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x73, 0x70, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x3d, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x7d,
	0x12, 0x8a, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x24, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x2a, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70,
	0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x3d, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x7d, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61,
	0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tenant_proto_rawDescOnce sync.Once
	file_tenant_proto_rawDescData = file_tenant_proto_rawDesc
)

func file_tenant_proto_rawDescGZIP() []byte {
	file_tenant_proto_rawDescOnce.Do(func() {
		file_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_tenant_proto_rawDescData)
	})
	return file_tenant_proto_rawDescData
}

var file_tenant_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tenant_proto_goTypes = []interface{}{
	(Workspace)(0),               // 0: erda.msp.tenant.Workspace
	(Type)(0),                    // 1: erda.msp.tenant.Type
	(*CreateTenantRequest)(nil),  // 2: erda.msp.tenant.CreateTenantRequest
	(*CreateTenantResponse)(nil), // 3: erda.msp.tenant.CreateTenantResponse
	(*GetTenantRequest)(nil),     // 4: erda.msp.tenant.GetTenantRequest
	(*GetTenantResponse)(nil),    // 5: erda.msp.tenant.GetTenantResponse
	(*DeleteTenantRequest)(nil),  // 6: erda.msp.tenant.DeleteTenantRequest
	(*DeleteTenantResponse)(nil), // 7: erda.msp.tenant.DeleteTenantResponse
	(*Tenant)(nil),               // 8: erda.msp.tenant.Tenant
}
var file_tenant_proto_depIdxs = []int32{
	8, // 0: erda.msp.tenant.CreateTenantResponse.data:type_name -> erda.msp.tenant.Tenant
	8, // 1: erda.msp.tenant.GetTenantResponse.data:type_name -> erda.msp.tenant.Tenant
	8, // 2: erda.msp.tenant.DeleteTenantResponse.data:type_name -> erda.msp.tenant.Tenant
	2, // 3: erda.msp.tenant.TenantService.CreateTenant:input_type -> erda.msp.tenant.CreateTenantRequest
	4, // 4: erda.msp.tenant.TenantService.GetTenant:input_type -> erda.msp.tenant.GetTenantRequest
	6, // 5: erda.msp.tenant.TenantService.DeleteTenant:input_type -> erda.msp.tenant.DeleteTenantRequest
	3, // 6: erda.msp.tenant.TenantService.CreateTenant:output_type -> erda.msp.tenant.CreateTenantResponse
	5, // 7: erda.msp.tenant.TenantService.GetTenant:output_type -> erda.msp.tenant.GetTenantResponse
	7, // 8: erda.msp.tenant.TenantService.DeleteTenant:output_type -> erda.msp.tenant.DeleteTenantResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tenant_proto_init() }
func file_tenant_proto_init() {
	if File_tenant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTenantRequest); i {
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
		file_tenant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTenantResponse); i {
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
		file_tenant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantRequest); i {
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
		file_tenant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantResponse); i {
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
		file_tenant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTenantRequest); i {
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
		file_tenant_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTenantResponse); i {
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
		file_tenant_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tenant); i {
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
			RawDescriptor: file_tenant_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tenant_proto_goTypes,
		DependencyIndexes: file_tenant_proto_depIdxs,
		EnumInfos:         file_tenant_proto_enumTypes,
		MessageInfos:      file_tenant_proto_msgTypes,
	}.Build()
	File_tenant_proto = out.File
	file_tenant_proto_rawDesc = nil
	file_tenant_proto_goTypes = nil
	file_tenant_proto_depIdxs = nil
}
