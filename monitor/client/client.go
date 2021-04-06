// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Source: todo.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/monitor/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// TODOService todo.proto
	TODOService() pb.TODOServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		todoservice: pb.NewTODOServiceClient(cc),
	}
}

type serviceClients struct {
	todoservice pb.TODOServiceClient
}

func (c *serviceClients) TODOService() pb.TODOServiceClient {
	return c.todoservice
}

type todoserviceWrapper struct {
	client pb.TODOServiceClient
	opts   []grpc1.CallOption
}

func (s *todoserviceWrapper) Do(ctx context.Context, req *pb.TODORequst) (*pb.TODOResponse, error) {
	return s.client.Do(ctx, req, s.opts...)
}
