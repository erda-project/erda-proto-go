// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: tracing.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	pb "github.com/erda-project/erda-proto-go/oap/common/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The span data model
//   of the observability analysis platform.
// Trace is calculated by span, stored as metric.
type Span struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique identifier for a trace. All spans from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
	// is considered invalid.
	//
	// This field is semantically required. Receiver should generate new
	// random trace_id if empty or invalid trace_id was received.
	//
	// This field is required.
	TraceID []byte `protobuf:"bytes,1,opt,name=traceID,proto3" json:"traceID,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes is considered
	// invalid.
	//
	// This field is semantically required. Receiver should generate new
	// random span_id if empty or invalid span_id was received.
	//
	// This field is required.
	SpanID []byte `protobuf:"bytes,2,opt,name=spanID,proto3" json:"spanID,omitempty"`
	// The `span_id` of this span's parent span. If this is a root span, then this
	// field must be empty. The ID is an 8-byte array.
	ParentSpanID      []byte `protobuf:"bytes,3,opt,name=parentSpanID,proto3" json:"parentSpanID,omitempty"`
	StartTimeUnixNano uint64 `protobuf:"fixed64,4,opt,name=startTimeUnixNano,proto3" json:"startTimeUnixNano,omitempty"`
	// end_time_unix_nano is the end time of the span. On the client side, this is the time
	// kept by the local machine where the span execution ends. On the server side, this
	// is the time when the server application handler stops running.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	//
	// This field is semantically required and it is expected that end_time >= start_time.
	EndTimeUnixNano uint64 `protobuf:"fixed64,5,opt,name=endTimeUnixNano,proto3" json:"endTimeUnixNano,omitempty"`
	// A description of the span's operation.
	//
	// For example, the name can be a qualified method name or a file name
	// and a line number where the operation is called. A best practice is to use
	// the same display name at the same call point in an application.
	// This makes it easier to correlate spans in different traces.
	//
	// This field is semantically required to be set to non-empty string.
	// When null or empty string received - receiver may use string "name"
	// as a replacement. There might be smarted algorithms implemented by
	// receiver to fix the empty span name.
	//
	// This field is required.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// Store trace and resource pointer
	Relations *pb.Relation `protobuf:"bytes,7,opt,name=relations,proto3" json:"relations,omitempty"`
	// Including label, attributes and resource
	Attributes map[string]string `protobuf:"bytes,8,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Span) Reset() {
	*x = Span{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Span) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Span) ProtoMessage() {}

func (x *Span) ProtoReflect() protoreflect.Message {
	mi := &file_tracing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Span.ProtoReflect.Descriptor instead.
func (*Span) Descriptor() ([]byte, []int) {
	return file_tracing_proto_rawDescGZIP(), []int{0}
}

func (x *Span) GetTraceID() []byte {
	if x != nil {
		return x.TraceID
	}
	return nil
}

func (x *Span) GetSpanID() []byte {
	if x != nil {
		return x.SpanID
	}
	return nil
}

func (x *Span) GetParentSpanID() []byte {
	if x != nil {
		return x.ParentSpanID
	}
	return nil
}

func (x *Span) GetStartTimeUnixNano() uint64 {
	if x != nil {
		return x.StartTimeUnixNano
	}
	return 0
}

func (x *Span) GetEndTimeUnixNano() uint64 {
	if x != nil {
		return x.EndTimeUnixNano
	}
	return 0
}

func (x *Span) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Span) GetRelations() *pb.Relation {
	if x != nil {
		return x.Relations
	}
	return nil
}

func (x *Span) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_tracing_proto protoreflect.FileDescriptor

var file_tracing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6f, 0x61, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6f,
	0x61, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a, 0x04, 0x53, 0x70, 0x61, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x61,
	0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49,
	0x44, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53,
	0x70, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x4e, 0x61, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x4e,
	0x61, 0x6e, 0x6f, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e,
	0x69, 0x78, 0x4e, 0x61, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0f, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x4e, 0x61, 0x6e, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6f, 0x61, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44, 0x0a, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6f, 0x61, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x53, 0x70, 0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72,
	0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x6f, 0x61, 0x70, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracing_proto_rawDescOnce sync.Once
	file_tracing_proto_rawDescData = file_tracing_proto_rawDesc
)

func file_tracing_proto_rawDescGZIP() []byte {
	file_tracing_proto_rawDescOnce.Do(func() {
		file_tracing_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracing_proto_rawDescData)
	})
	return file_tracing_proto_rawDescData
}

var file_tracing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tracing_proto_goTypes = []interface{}{
	(*Span)(nil),        // 0: erda.oap.trace.Span
	nil,                 // 1: erda.oap.trace.Span.AttributesEntry
	(*pb.Relation)(nil), // 2: erda.oap.common.Relation
}
var file_tracing_proto_depIdxs = []int32{
	2, // 0: erda.oap.trace.Span.relations:type_name -> erda.oap.common.Relation
	1, // 1: erda.oap.trace.Span.attributes:type_name -> erda.oap.trace.Span.AttributesEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tracing_proto_init() }
func file_tracing_proto_init() {
	if File_tracing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Span); i {
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
			RawDescriptor: file_tracing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tracing_proto_goTypes,
		DependencyIndexes: file_tracing_proto_depIdxs,
		MessageInfos:      file_tracing_proto_msgTypes,
	}.Build()
	File_tracing_proto = out.File
	file_tracing_proto_rawDesc = nil
	file_tracing_proto_goTypes = nil
	file_tracing_proto_depIdxs = nil
}
