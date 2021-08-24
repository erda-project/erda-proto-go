// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: entity.proto

package pb

import (
	_ "github.com/erda-project/erda-proto-go/oap/common/pb"
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

// The Entity data model
//   of the observability analysis platform.
// Unlike time series data, entity data has a unique ID,
//   and data can be inserted, updated, and deleted according to the unique ID.
type EntityRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unique
	EntityID string                     `protobuf:"bytes,1,opt,name=entityID,proto3" json:"entityID,omitempty"`
	Table    string                     `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	RowID    []byte                     `protobuf:"bytes,3,opt,name=rowID,proto3" json:"rowID,omitempty"`
	RowData  map[string]*structpb.Value `protobuf:"bytes,4,rep,name=rowData,proto3" json:"rowData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Including label, attributes and resource
	Attributes         map[string]string `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreateTimeUnixNano uint64            `protobuf:"fixed64,6,opt,name=createTimeUnixNano,proto3" json:"createTimeUnixNano,omitempty"`
	UpdateTimeUnixNano uint64            `protobuf:"fixed64,7,opt,name=updateTimeUnixNano,proto3" json:"updateTimeUnixNano,omitempty"`
}

func (x *EntityRow) Reset() {
	*x = EntityRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityRow) ProtoMessage() {}

func (x *EntityRow) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityRow.ProtoReflect.Descriptor instead.
func (*EntityRow) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{0}
}

func (x *EntityRow) GetEntityID() string {
	if x != nil {
		return x.EntityID
	}
	return ""
}

func (x *EntityRow) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *EntityRow) GetRowID() []byte {
	if x != nil {
		return x.RowID
	}
	return nil
}

func (x *EntityRow) GetRowData() map[string]*structpb.Value {
	if x != nil {
		return x.RowData
	}
	return nil
}

func (x *EntityRow) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *EntityRow) GetCreateTimeUnixNano() uint64 {
	if x != nil {
		return x.CreateTimeUnixNano
	}
	return 0
}

func (x *EntityRow) GetUpdateTimeUnixNano() uint64 {
	if x != nil {
		return x.UpdateTimeUnixNano
	}
	return 0
}

var File_entity_proto protoreflect.FileDescriptor

var file_entity_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x6f, 0x61, 0x70, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6f,
	0x61, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x03, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x6f, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x77, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x6f, 0x77, 0x49, 0x44, 0x12, 0x41, 0x0a, 0x07,
	0x72, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x6f, 0x61, 0x70, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x6f, 0x77, 0x2e, 0x52, 0x6f, 0x77, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x4a, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6f, 0x61, 0x70, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x6f, 0x77, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x4e, 0x61, 0x6e,
	0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x06, 0x52, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x4e, 0x61, 0x6e, 0x6f, 0x12, 0x2e, 0x0a, 0x12, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x4e, 0x61, 0x6e,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x06, 0x52, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x4e, 0x61, 0x6e, 0x6f, 0x1a, 0x52, 0x0a, 0x0c, 0x52,
	0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64,
	0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x6f, 0x61, 0x70, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_proto_rawDescOnce sync.Once
	file_entity_proto_rawDescData = file_entity_proto_rawDesc
)

func file_entity_proto_rawDescGZIP() []byte {
	file_entity_proto_rawDescOnce.Do(func() {
		file_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_proto_rawDescData)
	})
	return file_entity_proto_rawDescData
}

var file_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_entity_proto_goTypes = []interface{}{
	(*EntityRow)(nil),      // 0: erda.oap.entity.EntityRow
	nil,                    // 1: erda.oap.entity.EntityRow.RowDataEntry
	nil,                    // 2: erda.oap.entity.EntityRow.AttributesEntry
	(*structpb.Value)(nil), // 3: google.protobuf.Value
}
var file_entity_proto_depIdxs = []int32{
	1, // 0: erda.oap.entity.EntityRow.rowData:type_name -> erda.oap.entity.EntityRow.RowDataEntry
	2, // 1: erda.oap.entity.EntityRow.attributes:type_name -> erda.oap.entity.EntityRow.AttributesEntry
	3, // 2: erda.oap.entity.EntityRow.RowDataEntry.value:type_name -> google.protobuf.Value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_entity_proto_init() }
func file_entity_proto_init() {
	if File_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityRow); i {
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
			RawDescriptor: file_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_proto_goTypes,
		DependencyIndexes: file_entity_proto_depIdxs,
		MessageInfos:      file_entity_proto_msgTypes,
	}.Build()
	File_entity_proto = out.File
	file_entity_proto_rawDesc = nil
	file_entity_proto_goTypes = nil
	file_entity_proto_depIdxs = nil
}
