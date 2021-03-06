// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: entity.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterEntityServiceImp entity.proto
func RegisterEntityServiceImp(regester transport.Register, srv EntityServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterEntityServiceHandler(regester, EntityServiceHandler(srv), _ops.HTTP...)
	RegisterEntityServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.oap.entity.EntityService",
	)
}

var (
	entityServiceClientType  = reflect.TypeOf((*EntityServiceClient)(nil)).Elem()
	entityServiceServerType  = reflect.TypeOf((*EntityServiceServer)(nil)).Elem()
	entityServiceHandlerType = reflect.TypeOf((*EntityServiceHandler)(nil)).Elem()
)

// EntityServiceClientType .
func EntityServiceClientType() reflect.Type { return entityServiceClientType }

// EntityServiceServerType .
func EntityServiceServerType() reflect.Type { return entityServiceServerType }

// EntityServiceHandlerType .
func EntityServiceHandlerType() reflect.Type { return entityServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		entityServiceClientType,
		// server types
		entityServiceServerType,
		// handler types
		entityServiceHandlerType,
	}
}
