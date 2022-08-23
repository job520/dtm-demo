package main

import (
	test "demo/grpc/proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

const (
	Address = "127.0.0.1:50052"
)

type testService struct{}

var TestService = testService{}

func (h testService) TransIn(ctx context.Context, in *test.TestRequest) (*test.TestResponse, error) {
	resp := new(test.TestResponse)
	fmt.Println("TransIn")
	return resp, nil
}
func (h testService) TransInCompensate(ctx context.Context, in *test.TestRequest) (*test.TestResponse, error) {
	resp := new(test.TestResponse)
	fmt.Println("TransInCompensate")
	return resp, nil
}
func (h testService) TransOut(ctx context.Context, in *test.TestRequest) (*test.TestResponse, error) {
	resp := new(test.TestResponse)
	fmt.Println("TransOut")
	return resp, nil
}
func (h testService) TransOutCompensate(ctx context.Context, in *test.TestRequest) (*test.TestResponse, error) {
	resp := new(test.TestResponse)
	fmt.Println("TransOutCompensate")
	return resp, nil
}
func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	test.RegisterTestServer(s, TestService)
	fmt.Println("Listen on " + Address)
	s.Serve(listen)
}
