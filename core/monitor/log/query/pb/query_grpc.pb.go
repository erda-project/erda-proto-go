// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: query.proto

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

// LogQueryServiceClient is the client API for LogQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogQueryServiceClient interface {
	// +private
	GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogResponse, error)
	// for runtime log
	GetLogByRuntime(ctx context.Context, in *GetLogByRuntimeRequest, opts ...grpc.CallOption) (*GetLogByRuntimeResponse, error)
	// for organization log
	GetLogByOrganization(ctx context.Context, in *GetLogByOrganizationRequest, opts ...grpc.CallOption) (*GetLogByOrganizationResponse, error)
}

type logQueryServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewLogQueryServiceClient(cc grpc1.ClientConnInterface) LogQueryServiceClient {
	return &logQueryServiceClient{cc}
}

func (c *logQueryServiceClient) GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogResponse, error) {
	out := new(GetLogResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.log.query.LogQueryService/GetLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logQueryServiceClient) GetLogByRuntime(ctx context.Context, in *GetLogByRuntimeRequest, opts ...grpc.CallOption) (*GetLogByRuntimeResponse, error) {
	out := new(GetLogByRuntimeResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.log.query.LogQueryService/GetLogByRuntime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logQueryServiceClient) GetLogByOrganization(ctx context.Context, in *GetLogByOrganizationRequest, opts ...grpc.CallOption) (*GetLogByOrganizationResponse, error) {
	out := new(GetLogByOrganizationResponse)
	err := c.cc.Invoke(ctx, "/erda.core.monitor.log.query.LogQueryService/GetLogByOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogQueryServiceServer is the server API for LogQueryService service.
// All implementations should embed UnimplementedLogQueryServiceServer
// for forward compatibility
type LogQueryServiceServer interface {
	// +private
	GetLog(context.Context, *GetLogRequest) (*GetLogResponse, error)
	// for runtime log
	GetLogByRuntime(context.Context, *GetLogByRuntimeRequest) (*GetLogByRuntimeResponse, error)
	// for organization log
	GetLogByOrganization(context.Context, *GetLogByOrganizationRequest) (*GetLogByOrganizationResponse, error)
}

// UnimplementedLogQueryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLogQueryServiceServer struct {
}

func (*UnimplementedLogQueryServiceServer) GetLog(context.Context, *GetLogRequest) (*GetLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLog not implemented")
}
func (*UnimplementedLogQueryServiceServer) GetLogByRuntime(context.Context, *GetLogByRuntimeRequest) (*GetLogByRuntimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogByRuntime not implemented")
}
func (*UnimplementedLogQueryServiceServer) GetLogByOrganization(context.Context, *GetLogByOrganizationRequest) (*GetLogByOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogByOrganization not implemented")
}

func RegisterLogQueryServiceServer(s grpc1.ServiceRegistrar, srv LogQueryServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_LogQueryService_serviceDesc(srv, opts...), srv)
}

var _LogQueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.monitor.log.query.LogQueryService",
	HandlerType: (*LogQueryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "query.proto",
}

func _get_LogQueryService_serviceDesc(srv LogQueryServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_LogQueryService_GetLog_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetLog(ctx, req.(*GetLogRequest))
	}
	var _LogQueryService_GetLog_info transport.ServiceInfo
	if h.Interceptor != nil {
		_LogQueryService_GetLog_info = transport.NewServiceInfo("erda.core.monitor.log.query.LogQueryService", "GetLog", srv)
		_LogQueryService_GetLog_Handler = h.Interceptor(_LogQueryService_GetLog_Handler)
	}

	_LogQueryService_GetLogByRuntime_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetLogByRuntime(ctx, req.(*GetLogByRuntimeRequest))
	}
	var _LogQueryService_GetLogByRuntime_info transport.ServiceInfo
	if h.Interceptor != nil {
		_LogQueryService_GetLogByRuntime_info = transport.NewServiceInfo("erda.core.monitor.log.query.LogQueryService", "GetLogByRuntime", srv)
		_LogQueryService_GetLogByRuntime_Handler = h.Interceptor(_LogQueryService_GetLogByRuntime_Handler)
	}

	_LogQueryService_GetLogByOrganization_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetLogByOrganization(ctx, req.(*GetLogByOrganizationRequest))
	}
	var _LogQueryService_GetLogByOrganization_info transport.ServiceInfo
	if h.Interceptor != nil {
		_LogQueryService_GetLogByOrganization_info = transport.NewServiceInfo("erda.core.monitor.log.query.LogQueryService", "GetLogByOrganization", srv)
		_LogQueryService_GetLogByOrganization_Handler = h.Interceptor(_LogQueryService_GetLogByOrganization_Handler)
	}

	var serviceDesc = _LogQueryService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "GetLog",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetLogRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(LogQueryServiceServer).GetLog(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _LogQueryService_GetLog_info)
				}
				if interceptor == nil {
					return _LogQueryService_GetLog_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.log.query.LogQueryService/GetLog",
				}
				return interceptor(ctx, in, info, _LogQueryService_GetLog_Handler)
			},
		},
		{
			MethodName: "GetLogByRuntime",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetLogByRuntimeRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(LogQueryServiceServer).GetLogByRuntime(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _LogQueryService_GetLogByRuntime_info)
				}
				if interceptor == nil {
					return _LogQueryService_GetLogByRuntime_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.log.query.LogQueryService/GetLogByRuntime",
				}
				return interceptor(ctx, in, info, _LogQueryService_GetLogByRuntime_Handler)
			},
		},
		{
			MethodName: "GetLogByOrganization",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetLogByOrganizationRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(LogQueryServiceServer).GetLogByOrganization(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _LogQueryService_GetLogByOrganization_info)
				}
				if interceptor == nil {
					return _LogQueryService_GetLogByOrganization_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.monitor.log.query.LogQueryService/GetLogByOrganization",
				}
				return interceptor(ctx, in, info, _LogQueryService_GetLogByOrganization_Handler)
			},
		},
	}
	return &serviceDesc
}
