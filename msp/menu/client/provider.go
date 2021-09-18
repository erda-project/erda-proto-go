// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: menu.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/menu/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.msp.menu",
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
	clientsType           = reflect.TypeOf((*Client)(nil)).Elem()
	menuServiceClientType = reflect.TypeOf((*pb.MenuServiceClient)(nil)).Elem()
	menuServiceServerType = reflect.TypeOf((*pb.MenuServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.msp.menu-client":
		return p.client
	case "erda.msp.menu.MenuService":
		return &menuServiceWrapper{client: p.client.MenuService(), opts: opts}
	case "erda.msp.menu.MenuService.client":
		return p.client.MenuService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case menuServiceClientType:
		return p.client.MenuService()
	case menuServiceServerType:
		return &menuServiceWrapper{client: p.client.MenuService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.msp.menu-client", &servicehub.Spec{
		Services: []string{
			"erda.msp.menu.MenuService",
			"erda.msp.menu-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			menuServiceClientType,
			// server types
			menuServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
