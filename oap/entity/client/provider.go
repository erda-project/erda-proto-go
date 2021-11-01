// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: entity.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/oap/entity/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.oap.entity",
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
	entityServiceClientType = reflect.TypeOf((*pb.EntityServiceClient)(nil)).Elem()
	entityServiceServerType = reflect.TypeOf((*pb.EntityServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.oap.entity-client":
		return p.client
	case "erda.oap.entity.EntityService":
		return &entityServiceWrapper{client: p.client.EntityService(), opts: opts}
	case "erda.oap.entity.EntityService.client":
		return p.client.EntityService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case entityServiceClientType:
		return p.client.EntityService()
	case entityServiceServerType:
		return &entityServiceWrapper{client: p.client.EntityService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.oap.entity-client", &servicehub.Spec{
		Services: []string{
			"erda.oap.entity.EntityService",
			"erda.oap.entity-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			entityServiceClientType,
			// server types
			entityServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
