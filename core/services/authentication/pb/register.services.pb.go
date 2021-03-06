// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: authentication.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterAuthenticationServiceImp authentication.proto
func RegisterAuthenticationServiceImp(regester transport.Register, srv AuthenticationServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterAuthenticationServiceHandler(regester, AuthenticationServiceHandler(srv), _ops.HTTP...)
	RegisterAuthenticationServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.services.authentication.AuthenticationService",
	)
}

var (
	authenticationServiceClientType  = reflect.TypeOf((*AuthenticationServiceClient)(nil)).Elem()
	authenticationServiceServerType  = reflect.TypeOf((*AuthenticationServiceServer)(nil)).Elem()
	authenticationServiceHandlerType = reflect.TypeOf((*AuthenticationServiceHandler)(nil)).Elem()
)

// AuthenticationServiceClientType .
func AuthenticationServiceClientType() reflect.Type { return authenticationServiceClientType }

// AuthenticationServiceServerType .
func AuthenticationServiceServerType() reflect.Type { return authenticationServiceServerType }

// AuthenticationServiceHandlerType .
func AuthenticationServiceHandlerType() reflect.Type { return authenticationServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		authenticationServiceClientType,
		// server types
		authenticationServiceServerType,
		// handler types
		authenticationServiceHandlerType,
	}
}
