// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: event_query.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterEventQueryServiceImp event_query.proto
func RegisterEventQueryServiceImp(regester transport.Register, srv EventQueryServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterEventQueryServiceHandler(regester, EventQueryServiceHandler(srv), _ops.HTTP...)
	RegisterEventQueryServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.monitor.event.EventQueryService",
	)
}

var (
	eventQueryServiceClientType  = reflect.TypeOf((*EventQueryServiceClient)(nil)).Elem()
	eventQueryServiceServerType  = reflect.TypeOf((*EventQueryServiceServer)(nil)).Elem()
	eventQueryServiceHandlerType = reflect.TypeOf((*EventQueryServiceHandler)(nil)).Elem()
)

// EventQueryServiceClientType .
func EventQueryServiceClientType() reflect.Type { return eventQueryServiceClientType }

// EventQueryServiceServerType .
func EventQueryServiceServerType() reflect.Type { return eventQueryServiceServerType }

// EventQueryServiceHandlerType .
func EventQueryServiceHandlerType() reflect.Type { return eventQueryServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		eventQueryServiceClientType,
		// server types
		eventQueryServiceServerType,
		// handler types
		eventQueryServiceHandlerType,
	}
}
