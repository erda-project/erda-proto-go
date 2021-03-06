// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: label.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/label/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// LabelService label.proto
	LabelService() pb.LabelServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		labelService: pb.NewLabelServiceClient(cc),
	}
}

type serviceClients struct {
	labelService pb.LabelServiceClient
}

func (c *serviceClients) LabelService() pb.LabelServiceClient {
	return c.labelService
}

type labelServiceWrapper struct {
	client pb.LabelServiceClient
	opts   []grpc1.CallOption
}

func (s *labelServiceWrapper) PipelineLabelBatchInsert(ctx context.Context, req *pb.PipelineLabelBatchInsertRequest) (*pb.PipelineLabelBatchInsertResponse, error) {
	return s.client.PipelineLabelBatchInsert(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *labelServiceWrapper) PipelineLabelList(ctx context.Context, req *pb.PipelineLabelListRequest) (*pb.PipelineLabelListResponse, error) {
	return s.client.PipelineLabelList(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
