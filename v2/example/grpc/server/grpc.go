package grpc

import (
	"context"
	example "github.com/codingXiang/cxgateway/v2/example/grpc/pb"
)

type ExampleService struct{}

func NewExampleService() example.ExampleServiceServer {
	server := new(ExampleService)
	return server
}

func (t *ExampleService) Add(context context.Context, request *example.Request) (*example.Reply, error) {
	return new(example.Reply), nil
}

func (t *ExampleService) Remove(context context.Context, request *example.Request) (*example.Reply, error) {
	return new(example.Reply), nil
}
