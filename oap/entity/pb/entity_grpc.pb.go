// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: entity.proto

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

// EntityServiceClient is the client API for EntityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntityServiceClient interface {
	SetEntity(ctx context.Context, in *SetEntityRequest, opts ...grpc.CallOption) (*SetEntityResponse, error)
	RemoveEntity(ctx context.Context, in *RemoveEntityRequest, opts ...grpc.CallOption) (*RemoveEntityResponse, error)
	GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*GetEntityResponse, error)
	ListEntities(ctx context.Context, in *ListEntitiesRequest, opts ...grpc.CallOption) (*ListEntitiesResponse, error)
}

type entityServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewEntityServiceClient(cc grpc1.ClientConnInterface) EntityServiceClient {
	return &entityServiceClient{cc}
}

func (c *entityServiceClient) SetEntity(ctx context.Context, in *SetEntityRequest, opts ...grpc.CallOption) (*SetEntityResponse, error) {
	out := new(SetEntityResponse)
	err := c.cc.Invoke(ctx, "/erda.oap.entity.EntityService/SetEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) RemoveEntity(ctx context.Context, in *RemoveEntityRequest, opts ...grpc.CallOption) (*RemoveEntityResponse, error) {
	out := new(RemoveEntityResponse)
	err := c.cc.Invoke(ctx, "/erda.oap.entity.EntityService/RemoveEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*GetEntityResponse, error) {
	out := new(GetEntityResponse)
	err := c.cc.Invoke(ctx, "/erda.oap.entity.EntityService/GetEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityServiceClient) ListEntities(ctx context.Context, in *ListEntitiesRequest, opts ...grpc.CallOption) (*ListEntitiesResponse, error) {
	out := new(ListEntitiesResponse)
	err := c.cc.Invoke(ctx, "/erda.oap.entity.EntityService/ListEntities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityServiceServer is the server API for EntityService service.
// All implementations should embed UnimplementedEntityServiceServer
// for forward compatibility
type EntityServiceServer interface {
	SetEntity(context.Context, *SetEntityRequest) (*SetEntityResponse, error)
	RemoveEntity(context.Context, *RemoveEntityRequest) (*RemoveEntityResponse, error)
	GetEntity(context.Context, *GetEntityRequest) (*GetEntityResponse, error)
	ListEntities(context.Context, *ListEntitiesRequest) (*ListEntitiesResponse, error)
}

// UnimplementedEntityServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEntityServiceServer struct {
}

func (*UnimplementedEntityServiceServer) SetEntity(context.Context, *SetEntityRequest) (*SetEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetEntity not implemented")
}
func (*UnimplementedEntityServiceServer) RemoveEntity(context.Context, *RemoveEntityRequest) (*RemoveEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEntity not implemented")
}
func (*UnimplementedEntityServiceServer) GetEntity(context.Context, *GetEntityRequest) (*GetEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntity not implemented")
}
func (*UnimplementedEntityServiceServer) ListEntities(context.Context, *ListEntitiesRequest) (*ListEntitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntities not implemented")
}

func RegisterEntityServiceServer(s grpc1.ServiceRegistrar, srv EntityServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_EntityService_serviceDesc(srv, opts...), srv)
}

var _EntityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.oap.entity.EntityService",
	HandlerType: (*EntityServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "entity.proto",
}

func _get_EntityService_serviceDesc(srv EntityServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_EntityService_SetEntity_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.SetEntity(ctx, req.(*SetEntityRequest))
	}
	var _EntityService_SetEntity_info transport.ServiceInfo
	if h.Interceptor != nil {
		_EntityService_SetEntity_info = transport.NewServiceInfo("erda.oap.entity.EntityService", "SetEntity", srv)
		_EntityService_SetEntity_Handler = h.Interceptor(_EntityService_SetEntity_Handler)
	}

	_EntityService_RemoveEntity_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.RemoveEntity(ctx, req.(*RemoveEntityRequest))
	}
	var _EntityService_RemoveEntity_info transport.ServiceInfo
	if h.Interceptor != nil {
		_EntityService_RemoveEntity_info = transport.NewServiceInfo("erda.oap.entity.EntityService", "RemoveEntity", srv)
		_EntityService_RemoveEntity_Handler = h.Interceptor(_EntityService_RemoveEntity_Handler)
	}

	_EntityService_GetEntity_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetEntity(ctx, req.(*GetEntityRequest))
	}
	var _EntityService_GetEntity_info transport.ServiceInfo
	if h.Interceptor != nil {
		_EntityService_GetEntity_info = transport.NewServiceInfo("erda.oap.entity.EntityService", "GetEntity", srv)
		_EntityService_GetEntity_Handler = h.Interceptor(_EntityService_GetEntity_Handler)
	}

	_EntityService_ListEntities_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListEntities(ctx, req.(*ListEntitiesRequest))
	}
	var _EntityService_ListEntities_info transport.ServiceInfo
	if h.Interceptor != nil {
		_EntityService_ListEntities_info = transport.NewServiceInfo("erda.oap.entity.EntityService", "ListEntities", srv)
		_EntityService_ListEntities_Handler = h.Interceptor(_EntityService_ListEntities_Handler)
	}

	var serviceDesc = _EntityService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "SetEntity",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(SetEntityRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(EntityServiceServer).SetEntity(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _EntityService_SetEntity_info)
				}
				if interceptor == nil {
					return _EntityService_SetEntity_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.oap.entity.EntityService/SetEntity",
				}
				return interceptor(ctx, in, info, _EntityService_SetEntity_Handler)
			},
		},
		{
			MethodName: "RemoveEntity",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(RemoveEntityRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(EntityServiceServer).RemoveEntity(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _EntityService_RemoveEntity_info)
				}
				if interceptor == nil {
					return _EntityService_RemoveEntity_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.oap.entity.EntityService/RemoveEntity",
				}
				return interceptor(ctx, in, info, _EntityService_RemoveEntity_Handler)
			},
		},
		{
			MethodName: "GetEntity",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetEntityRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(EntityServiceServer).GetEntity(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _EntityService_GetEntity_info)
				}
				if interceptor == nil {
					return _EntityService_GetEntity_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.oap.entity.EntityService/GetEntity",
				}
				return interceptor(ctx, in, info, _EntityService_GetEntity_Handler)
			},
		},
		{
			MethodName: "ListEntities",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListEntitiesRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(EntityServiceServer).ListEntities(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _EntityService_ListEntities_info)
				}
				if interceptor == nil {
					return _EntityService_ListEntities_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.oap.entity.EntityService/ListEntities",
				}
				return interceptor(ctx, in, info, _EntityService_ListEntities_Handler)
			},
		},
	}
	return &serviceDesc
}
