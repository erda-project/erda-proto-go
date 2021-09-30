// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: legacy_consumer.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/consumer/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// LegacyConsumerService legacy_consumer.proto
	LegacyConsumerService() pb.LegacyConsumerServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		legacyConsumerService: pb.NewLegacyConsumerServiceClient(cc),
	}
}

type serviceClients struct {
	legacyConsumerService pb.LegacyConsumerServiceClient
}

func (c *serviceClients) LegacyConsumerService() pb.LegacyConsumerServiceClient {
	return c.legacyConsumerService
}

type legacyConsumerServiceWrapper struct {
	client pb.LegacyConsumerServiceClient
	opts   []grpc1.CallOption
}

func (s *legacyConsumerServiceWrapper) GetConsumer(ctx context.Context, req *pb.GetConsumerRequest) (*pb.GetConsumerResponse, error) {
	return s.client.GetConsumer(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
