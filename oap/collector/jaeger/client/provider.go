// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: jaeger.proto

package client

import (
	fmt "fmt"
	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/oap/collector/jaeger/pb"
	grpc1 "google.golang.org/grpc"
	reflect "reflect"
	strings "strings"
)

var dependencies = []string{
	"grpc-client@erda.oap.collector.jaeger",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType             = reflect.TypeOf((*Client)(nil)).Elem()
	jaegerServiceClientType = reflect.TypeOf((*pb.JaegerServiceClient)(nil)).Elem()
	jaegerServiceServerType = reflect.TypeOf((*pb.JaegerServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.oap.collector.jaeger-client":
		return p.client
	case "erda.oap.collector.jaeger.JaegerService":
		return &jaegerServiceWrapper{client: p.client.JaegerService(), opts: opts}
	case "erda.oap.collector.jaeger.JaegerService.client":
		return p.client.JaegerService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case jaegerServiceClientType:
		return p.client.JaegerService()
	case jaegerServiceServerType:
		return &jaegerServiceWrapper{client: p.client.JaegerService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.oap.collector.jaeger-client", &servicehub.Spec{
		Services: []string{
			"erda.oap.collector.jaeger.JaegerService",
			"erda.oap.collector.jaeger-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			jaegerServiceClientType,
			// server types
			jaegerServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
