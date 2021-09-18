// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: build.proto

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

// BuildServiceClient is the client API for BuildService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuildServiceClient interface {
	QueryBuildArtifact(ctx context.Context, in *BuildArtifactQueryRequest, opts ...grpc.CallOption) (*BuildArtifactQueryResponse, error)
	RegisterBuildArtifact(ctx context.Context, in *BuildArtifactRegisterRequest, opts ...grpc.CallOption) (*BuildArtifactRegisterResponse, error)
	ReportBuildCache(ctx context.Context, in *BuildCacheReportRequest, opts ...grpc.CallOption) (*BuildCacheReportResponse, error)
}

type buildServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewBuildServiceClient(cc grpc1.ClientConnInterface) BuildServiceClient {
	return &buildServiceClient{cc}
}

func (c *buildServiceClient) QueryBuildArtifact(ctx context.Context, in *BuildArtifactQueryRequest, opts ...grpc.CallOption) (*BuildArtifactQueryResponse, error) {
	out := new(BuildArtifactQueryResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.build.BuildService/QueryBuildArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) RegisterBuildArtifact(ctx context.Context, in *BuildArtifactRegisterRequest, opts ...grpc.CallOption) (*BuildArtifactRegisterResponse, error) {
	out := new(BuildArtifactRegisterResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.build.BuildService/RegisterBuildArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildServiceClient) ReportBuildCache(ctx context.Context, in *BuildCacheReportRequest, opts ...grpc.CallOption) (*BuildCacheReportResponse, error) {
	out := new(BuildCacheReportResponse)
	err := c.cc.Invoke(ctx, "/erda.core.pipeline.build.BuildService/ReportBuildCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildServiceServer is the server API for BuildService service.
// All implementations should embed UnimplementedBuildServiceServer
// for forward compatibility
type BuildServiceServer interface {
	QueryBuildArtifact(context.Context, *BuildArtifactQueryRequest) (*BuildArtifactQueryResponse, error)
	RegisterBuildArtifact(context.Context, *BuildArtifactRegisterRequest) (*BuildArtifactRegisterResponse, error)
	ReportBuildCache(context.Context, *BuildCacheReportRequest) (*BuildCacheReportResponse, error)
}

// UnimplementedBuildServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBuildServiceServer struct {
}

func (*UnimplementedBuildServiceServer) QueryBuildArtifact(context.Context, *BuildArtifactQueryRequest) (*BuildArtifactQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBuildArtifact not implemented")
}
func (*UnimplementedBuildServiceServer) RegisterBuildArtifact(context.Context, *BuildArtifactRegisterRequest) (*BuildArtifactRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterBuildArtifact not implemented")
}
func (*UnimplementedBuildServiceServer) ReportBuildCache(context.Context, *BuildCacheReportRequest) (*BuildCacheReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportBuildCache not implemented")
}

func RegisterBuildServiceServer(s grpc1.ServiceRegistrar, srv BuildServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_BuildService_serviceDesc(srv, opts...), srv)
}

var _BuildService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.pipeline.build.BuildService",
	HandlerType: (*BuildServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "build.proto",
}

func _get_BuildService_serviceDesc(srv BuildServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_BuildService_QueryBuildArtifact_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.QueryBuildArtifact(ctx, req.(*BuildArtifactQueryRequest))
	}
	var _BuildService_QueryBuildArtifact_info transport.ServiceInfo
	if h.Interceptor != nil {
		_BuildService_QueryBuildArtifact_info = transport.NewServiceInfo("erda.core.pipeline.build.BuildService", "QueryBuildArtifact", srv)
		_BuildService_QueryBuildArtifact_Handler = h.Interceptor(_BuildService_QueryBuildArtifact_Handler)
	}

	_BuildService_RegisterBuildArtifact_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.RegisterBuildArtifact(ctx, req.(*BuildArtifactRegisterRequest))
	}
	var _BuildService_RegisterBuildArtifact_info transport.ServiceInfo
	if h.Interceptor != nil {
		_BuildService_RegisterBuildArtifact_info = transport.NewServiceInfo("erda.core.pipeline.build.BuildService", "RegisterBuildArtifact", srv)
		_BuildService_RegisterBuildArtifact_Handler = h.Interceptor(_BuildService_RegisterBuildArtifact_Handler)
	}

	_BuildService_ReportBuildCache_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ReportBuildCache(ctx, req.(*BuildCacheReportRequest))
	}
	var _BuildService_ReportBuildCache_info transport.ServiceInfo
	if h.Interceptor != nil {
		_BuildService_ReportBuildCache_info = transport.NewServiceInfo("erda.core.pipeline.build.BuildService", "ReportBuildCache", srv)
		_BuildService_ReportBuildCache_Handler = h.Interceptor(_BuildService_ReportBuildCache_Handler)
	}

	var serviceDesc = _BuildService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "QueryBuildArtifact",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(BuildArtifactQueryRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(BuildServiceServer).QueryBuildArtifact(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _BuildService_QueryBuildArtifact_info)
				}
				if interceptor == nil {
					return _BuildService_QueryBuildArtifact_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.build.BuildService/QueryBuildArtifact",
				}
				return interceptor(ctx, in, info, _BuildService_QueryBuildArtifact_Handler)
			},
		},
		{
			MethodName: "RegisterBuildArtifact",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(BuildArtifactRegisterRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(BuildServiceServer).RegisterBuildArtifact(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _BuildService_RegisterBuildArtifact_info)
				}
				if interceptor == nil {
					return _BuildService_RegisterBuildArtifact_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.build.BuildService/RegisterBuildArtifact",
				}
				return interceptor(ctx, in, info, _BuildService_RegisterBuildArtifact_Handler)
			},
		},
		{
			MethodName: "ReportBuildCache",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(BuildCacheReportRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(BuildServiceServer).ReportBuildCache(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _BuildService_ReportBuildCache_info)
				}
				if interceptor == nil {
					return _BuildService_ReportBuildCache_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.pipeline.build.BuildService/ReportBuildCache",
				}
				return interceptor(ctx, in, info, _BuildService_ReportBuildCache_Handler)
			},
		},
	}
	return &serviceDesc
}
