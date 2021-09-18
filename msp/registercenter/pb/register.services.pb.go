// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: registercenter.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterRegisterCenterServiceImp registercenter.proto
func RegisterRegisterCenterServiceImp(regester transport.Register, srv RegisterCenterServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterRegisterCenterServiceHandler(regester, RegisterCenterServiceHandler(srv), _ops.HTTP...)
	RegisterRegisterCenterServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.msp.registercenter.RegisterCenterService",
	)
}

var (
	registerCenterServiceClientType  = reflect.TypeOf((*RegisterCenterServiceClient)(nil)).Elem()
	registerCenterServiceServerType  = reflect.TypeOf((*RegisterCenterServiceServer)(nil)).Elem()
	registerCenterServiceHandlerType = reflect.TypeOf((*RegisterCenterServiceHandler)(nil)).Elem()
)

// RegisterCenterServiceClientType .
func RegisterCenterServiceClientType() reflect.Type { return registerCenterServiceClientType }

// RegisterCenterServiceServerType .
func RegisterCenterServiceServerType() reflect.Type { return registerCenterServiceServerType }

// RegisterCenterServiceHandlerType .
func RegisterCenterServiceHandlerType() reflect.Type { return registerCenterServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		registerCenterServiceClientType,
		// server types
		registerCenterServiceServerType,
		// handler types
		registerCenterServiceHandlerType,
	}
}
