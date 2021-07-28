// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: extension.proto

package pb

import (
	transport "github.com/erda-project/erda-infra/pkg/transport"
	reflect "reflect"
)

// RegisterExtensionServiceImp extension.proto
func RegisterExtensionServiceImp(regester transport.Register, srv ExtensionServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterExtensionServiceHandler(regester, ExtensionServiceHandler(srv), _ops.HTTP...)
	RegisterExtensionServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.dicehub.extension.ExtensionService",
	)
}

var (
	extensionServiceClientType  = reflect.TypeOf((*ExtensionServiceClient)(nil)).Elem()
	extensionServiceServerType  = reflect.TypeOf((*ExtensionServiceServer)(nil)).Elem()
	extensionServiceHandlerType = reflect.TypeOf((*ExtensionServiceHandler)(nil)).Elem()
)

// ExtensionServiceClientType .
func ExtensionServiceClientType() reflect.Type { return extensionServiceClientType }

// ExtensionServiceServerType .
func ExtensionServiceServerType() reflect.Type { return extensionServiceServerType }

// ExtensionServiceHandlerType .
func ExtensionServiceHandlerType() reflect.Type { return extensionServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		extensionServiceClientType,
		// server types
		extensionServiceServerType,
		// handler types
		extensionServiceHandlerType,
	}
}
