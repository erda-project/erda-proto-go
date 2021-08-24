// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: accesskey.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/services/accesskey/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// AccessKeyService accesskey.proto
	AccessKeyService() pb.AccessKeyServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		accessKeyService: pb.NewAccessKeyServiceClient(cc),
	}
}

type serviceClients struct {
	accessKeyService pb.AccessKeyServiceClient
}

func (c *serviceClients) AccessKeyService() pb.AccessKeyServiceClient {
	return c.accessKeyService
}

type accessKeyServiceWrapper struct {
	client pb.AccessKeyServiceClient
	opts   []grpc1.CallOption
}

func (s *accessKeyServiceWrapper) QueryAccessKeys(ctx context.Context, req *pb.QueryAccessKeysRequest) (*pb.QueryAccessKeysResponse, error) {
	return s.client.QueryAccessKeys(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *accessKeyServiceWrapper) GetAccessKey(ctx context.Context, req *pb.GetAccessKeysRequest) (*pb.GetAccessKeysResponse, error) {
	return s.client.GetAccessKey(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *accessKeyServiceWrapper) CreateAccessKeys(ctx context.Context, req *pb.CreateAccessKeysRequest) (*pb.CreateAccessKeysResponse, error) {
	return s.client.CreateAccessKeys(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *accessKeyServiceWrapper) UpdateAccessKeys(ctx context.Context, req *pb.UpdateAccessKeysRequest) (*pb.UpdateAccessKeysResponse, error) {
	return s.client.UpdateAccessKeys(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *accessKeyServiceWrapper) DeleteAccessKeys(ctx context.Context, req *pb.DeleteAccessKeysRequest) (*pb.DeleteAccessKeysResponse, error) {
	return s.client.DeleteAccessKeys(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
