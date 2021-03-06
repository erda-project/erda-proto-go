// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: build.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterBuildServiceImp build.proto
func RegisterBuildServiceImp(regester transport.Register, srv BuildServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterBuildServiceHandler(regester, BuildServiceHandler(srv), _ops.HTTP...)
	RegisterBuildServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.pipeline.build.BuildService",
	)
}

var (
	buildServiceClientType  = reflect.TypeOf((*BuildServiceClient)(nil)).Elem()
	buildServiceServerType  = reflect.TypeOf((*BuildServiceServer)(nil)).Elem()
	buildServiceHandlerType = reflect.TypeOf((*BuildServiceHandler)(nil)).Elem()
)

// BuildServiceClientType .
func BuildServiceClientType() reflect.Type { return buildServiceClientType }

// BuildServiceServerType .
func BuildServiceServerType() reflect.Type { return buildServiceServerType }

// BuildServiceHandlerType .
func BuildServiceHandlerType() reflect.Type { return buildServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		buildServiceClientType,
		// server types
		buildServiceServerType,
		// handler types
		buildServiceHandlerType,
	}
}
