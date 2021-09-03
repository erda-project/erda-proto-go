// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: jaeger.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/common/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// JaegerServiceClient is the client API for JaegerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JaegerServiceClient interface {
	SpansWithThrift(ctx context.Context, in *pb.VoidRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error)
}

type jaegerServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewJaegerServiceClient(cc grpc1.ClientConnInterface) JaegerServiceClient {
	return &jaegerServiceClient{cc}
}

func (c *jaegerServiceClient) SpansWithThrift(ctx context.Context, in *pb.VoidRequest, opts ...grpc.CallOption) (*pb.VoidResponse, error) {
	out := new(pb.VoidResponse)
	err := c.cc.Invoke(ctx, "/erda.oap.collector.jaeger.JaegerService/SpansWithThrift", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JaegerServiceServer is the server API for JaegerService service.
// All implementations should embed UnimplementedJaegerServiceServer
// for forward compatibility
type JaegerServiceServer interface {
	SpansWithThrift(context.Context, *pb.VoidRequest) (*pb.VoidResponse, error)
}

// UnimplementedJaegerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedJaegerServiceServer struct {
}

func (*UnimplementedJaegerServiceServer) SpansWithThrift(context.Context, *pb.VoidRequest) (*pb.VoidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpansWithThrift not implemented")
}

func RegisterJaegerServiceServer(s grpc1.ServiceRegistrar, srv JaegerServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_JaegerService_serviceDesc(srv, opts...), srv)
}

var _JaegerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.oap.collector.jaeger.JaegerService",
	HandlerType: (*JaegerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "jaeger.proto",
}

func _get_JaegerService_serviceDesc(srv JaegerServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_JaegerService_SpansWithThrift_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.SpansWithThrift(ctx, req.(*pb.VoidRequest))
	}
	var _JaegerService_SpansWithThrift_info transport.ServiceInfo
	if h.Interceptor != nil {
		_JaegerService_SpansWithThrift_info = transport.NewServiceInfo("erda.oap.collector.jaeger.JaegerService", "SpansWithThrift", srv)
		_JaegerService_SpansWithThrift_Handler = h.Interceptor(_JaegerService_SpansWithThrift_Handler)
	}

	var serviceDesc = _JaegerService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "SpansWithThrift",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(pb.VoidRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(JaegerServiceServer).SpansWithThrift(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _JaegerService_SpansWithThrift_info)
				}
				if interceptor == nil {
					return _JaegerService_SpansWithThrift_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.oap.collector.jaeger.JaegerService/SpansWithThrift",
				}
				return interceptor(ctx, in, info, _JaegerService_SpansWithThrift_Handler)
			},
		},
	}
	return &serviceDesc
}
