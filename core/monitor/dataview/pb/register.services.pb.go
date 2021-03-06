// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: dataview.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterDataViewServiceImp dataview.proto
func RegisterDataViewServiceImp(regester transport.Register, srv DataViewServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterDataViewServiceHandler(regester, DataViewServiceHandler(srv), _ops.HTTP...)
	RegisterDataViewServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.monitor.dataview.DataViewService",
	)
}

var (
	dataViewServiceClientType  = reflect.TypeOf((*DataViewServiceClient)(nil)).Elem()
	dataViewServiceServerType  = reflect.TypeOf((*DataViewServiceServer)(nil)).Elem()
	dataViewServiceHandlerType = reflect.TypeOf((*DataViewServiceHandler)(nil)).Elem()
)

// DataViewServiceClientType .
func DataViewServiceClientType() reflect.Type { return dataViewServiceClientType }

// DataViewServiceServerType .
func DataViewServiceServerType() reflect.Type { return dataViewServiceServerType }

// DataViewServiceHandlerType .
func DataViewServiceHandlerType() reflect.Type { return dataViewServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		dataViewServiceClientType,
		// server types
		dataViewServiceServerType,
		// handler types
		dataViewServiceHandlerType,
	}
}
