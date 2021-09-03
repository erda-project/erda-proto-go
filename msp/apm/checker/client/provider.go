// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: checker.proto, checker_v1.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/checker/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.msp.apm.checker",
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
	clientsType                = reflect.TypeOf((*Client)(nil)).Elem()
	checkerServiceClientType   = reflect.TypeOf((*pb.CheckerServiceClient)(nil)).Elem()
	checkerServiceServerType   = reflect.TypeOf((*pb.CheckerServiceServer)(nil)).Elem()
	checkerV1ServiceClientType = reflect.TypeOf((*pb.CheckerV1ServiceClient)(nil)).Elem()
	checkerV1ServiceServerType = reflect.TypeOf((*pb.CheckerV1ServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.msp.apm.checker-client":
		return p.client
	case "erda.msp.apm.checker.CheckerService":
		return &checkerServiceWrapper{client: p.client.CheckerService(), opts: opts}
	case "erda.msp.apm.checker.CheckerService.client":
		return p.client.CheckerService()
	case "erda.msp.apm.checker.CheckerV1Service":
		return &checkerV1ServiceWrapper{client: p.client.CheckerV1Service(), opts: opts}
	case "erda.msp.apm.checker.CheckerV1Service.client":
		return p.client.CheckerV1Service()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case checkerServiceClientType:
		return p.client.CheckerService()
	case checkerServiceServerType:
		return &checkerServiceWrapper{client: p.client.CheckerService(), opts: opts}
	case checkerV1ServiceClientType:
		return p.client.CheckerV1Service()
	case checkerV1ServiceServerType:
		return &checkerV1ServiceWrapper{client: p.client.CheckerV1Service(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.msp.apm.checker-client", &servicehub.Spec{
		Services: []string{
			"erda.msp.apm.checker.CheckerService",
			"erda.msp.apm.checker.CheckerV1Service",
			"erda.msp.apm.checker-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			checkerServiceClientType,
			checkerV1ServiceClientType,
			// server types
			checkerServiceServerType,
			checkerV1ServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
