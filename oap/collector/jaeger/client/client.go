// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: jaeger.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb1 "github.com/erda-project/erda-proto-go/common/pb"
	pb "github.com/erda-project/erda-proto-go/oap/collector/jaeger/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// JaegerService jaeger.proto
	JaegerService() pb.JaegerServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		jaegerService: pb.NewJaegerServiceClient(cc),
	}
}

type serviceClients struct {
	jaegerService pb.JaegerServiceClient
}

func (c *serviceClients) JaegerService() pb.JaegerServiceClient {
	return c.jaegerService
}

type jaegerServiceWrapper struct {
	client pb.JaegerServiceClient
	opts   []grpc1.CallOption
}

func (s *jaegerServiceWrapper) SpansWithThrift(ctx context.Context, req *pb1.VoidRequest) (*pb1.VoidResponse, error) {
	return s.client.SpansWithThrift(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
