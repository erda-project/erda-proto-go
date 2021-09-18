// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: queue.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/queue/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.pipeline.queue",
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
	clientsType            = reflect.TypeOf((*Client)(nil)).Elem()
	queueServiceClientType = reflect.TypeOf((*pb.QueueServiceClient)(nil)).Elem()
	queueServiceServerType = reflect.TypeOf((*pb.QueueServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.pipeline.queue-client":
		return p.client
	case "erda.core.pipeline.queue.QueueService":
		return &queueServiceWrapper{client: p.client.QueueService(), opts: opts}
	case "erda.core.pipeline.queue.QueueService.client":
		return p.client.QueueService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case queueServiceClientType:
		return p.client.QueueService()
	case queueServiceServerType:
		return &queueServiceWrapper{client: p.client.QueueService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.pipeline.queue-client", &servicehub.Spec{
		Services: []string{
			"erda.core.pipeline.queue.QueueService",
			"erda.core.pipeline.queue-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			queueServiceClientType,
			// server types
			queueServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
