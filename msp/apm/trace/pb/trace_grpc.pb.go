// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: trace.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TraceServiceClient interface {
	GetSpans(ctx context.Context, in *GetSpansRequest, opts ...grpc.CallOption) (*GetSpansResponse, error)
	GetTraces(ctx context.Context, in *GetTracesRequest, opts ...grpc.CallOption) (*GetTracesResponse, error)
}

type traceServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewTraceServiceClient(cc grpc1.ClientConnInterface) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) GetSpans(ctx context.Context, in *GetSpansRequest, opts ...grpc.CallOption) (*GetSpansResponse, error) {
	out := new(GetSpansResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.trace.TraceService/GetSpans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceServiceClient) GetTraces(ctx context.Context, in *GetTracesRequest, opts ...grpc.CallOption) (*GetTracesResponse, error) {
	out := new(GetTracesResponse)
	err := c.cc.Invoke(ctx, "/erda.msp.apm.trace.TraceService/GetTraces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceServiceServer is the server API for TraceService service.
// All implementations should embed UnimplementedTraceServiceServer
// for forward compatibility
type TraceServiceServer interface {
	GetSpans(context.Context, *GetSpansRequest) (*GetSpansResponse, error)
	GetTraces(context.Context, *GetTracesRequest) (*GetTracesResponse, error)
}

// UnimplementedTraceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (*UnimplementedTraceServiceServer) GetSpans(context.Context, *GetSpansRequest) (*GetSpansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpans not implemented")
}
func (*UnimplementedTraceServiceServer) GetTraces(context.Context, *GetTracesRequest) (*GetTracesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraces not implemented")
}

func RegisterTraceServiceServer(s grpc1.ServiceRegistrar, srv TraceServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_TraceService_serviceDesc(srv, opts...), srv)
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.msp.apm.trace.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "trace.proto",
}

func _get_TraceService_serviceDesc(srv TraceServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_TraceService_GetSpans_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetSpans(ctx, req.(*GetSpansRequest))
	}
	var _TraceService_GetSpans_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TraceService_GetSpans_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetSpans", srv)
		_TraceService_GetSpans_Handler = h.Interceptor(_TraceService_GetSpans_Handler)
	}

	_TraceService_GetTraces_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetTraces(ctx, req.(*GetTracesRequest))
	}
	var _TraceService_GetTraces_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TraceService_GetTraces_info = transport.NewServiceInfo("erda.msp.apm.trace.TraceService", "GetTraces", srv)
		_TraceService_GetTraces_Handler = h.Interceptor(_TraceService_GetTraces_Handler)
	}

	var serviceDesc = _TraceService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "GetSpans",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetSpansRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TraceServiceServer).GetSpans(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TraceService_GetSpans_info)
				}
				if interceptor == nil {
					return _TraceService_GetSpans_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.trace.TraceService/GetSpans",
				}
				return interceptor(ctx, in, info, _TraceService_GetSpans_Handler)
			},
		},
		{
			MethodName: "GetTraces",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetTracesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TraceServiceServer).GetTraces(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TraceService_GetTraces_info)
				}
				if interceptor == nil {
					return _TraceService_GetTraces_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.msp.apm.trace.TraceService/GetTraces",
				}
				return interceptor(ctx, in, info, _TraceService_GetTraces_Handler)
			},
		},
	}
	return &serviceDesc
}
