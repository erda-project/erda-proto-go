// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: runtime_service.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/runtime_service/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// RuntimeService runtime_service.proto
	RuntimeService() pb.RuntimeServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		runtimeService: pb.NewRuntimeServiceClient(cc),
	}
}

type serviceClients struct {
	runtimeService pb.RuntimeServiceClient
}

func (c *serviceClients) RuntimeService() pb.RuntimeServiceClient {
	return c.runtimeService
}

type runtimeServiceWrapper struct {
	client pb.RuntimeServiceClient
	opts   []grpc1.CallOption
}

func (s *runtimeServiceWrapper) ChangeRuntime(ctx context.Context, req *pb.ChangeRuntimeRequest) (*pb.ChangeRuntimeResponse, error) {
	return s.client.ChangeRuntime(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *runtimeServiceWrapper) DeleteRuntime(ctx context.Context, req *pb.DeleteRuntimeRequest) (*pb.DeleteRuntimeResponse, error) {
	return s.client.DeleteRuntime(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *runtimeServiceWrapper) GetApps(ctx context.Context, req *pb.GetAppsRequest) (*pb.GetAppsResponse, error) {
	return s.client.GetApps(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *runtimeServiceWrapper) GetServiceRuntimes(ctx context.Context, req *pb.GetServiceRuntimesRequest) (*pb.GetServiceRuntimesResponse, error) {
	return s.client.GetServiceRuntimes(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *runtimeServiceWrapper) GetServiceApiPrefix(ctx context.Context, req *pb.GetServiceApiPrefixRequest) (*pb.GetServiceApiPrefixResponse, error) {
	return s.client.GetServiceApiPrefix(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
