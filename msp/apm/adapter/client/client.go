// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: adapter.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/adapter/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// AdapterService adapter.proto
	AdapterService() pb.AdapterServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		adapterService: pb.NewAdapterServiceClient(cc),
	}
}

type serviceClients struct {
	adapterService pb.AdapterServiceClient
}

func (c *serviceClients) AdapterService() pb.AdapterServiceClient {
	return c.adapterService
}

type adapterServiceWrapper struct {
	client pb.AdapterServiceClient
	opts   []grpc1.CallOption
}

func (s *adapterServiceWrapper) GetAdapters(ctx context.Context, req *pb.GetAdapterRequest) (*pb.GetAdapterResponse, error) {
	return s.client.GetAdapters(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}