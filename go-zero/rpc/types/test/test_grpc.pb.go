// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: test.proto

package test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestClient interface {
	TransIn(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Reply, error)
	TransInCompensate(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Reply, error)
	TransOut(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Reply, error)
	TransOutCompensate(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Reply, error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	return &testClient{cc}
}

func (c *testClient) TransIn(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/test.Test/TransIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) TransInCompensate(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/test.Test/TransInCompensate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) TransOut(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/test.Test/TransOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) TransOutCompensate(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/test.Test/TransOutCompensate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
// All implementations must embed UnimplementedTestServer
// for forward compatibility
type TestServer interface {
	TransIn(context.Context, *Req) (*Reply, error)
	TransInCompensate(context.Context, *Req) (*Reply, error)
	TransOut(context.Context, *Req) (*Reply, error)
	TransOutCompensate(context.Context, *Req) (*Reply, error)
	mustEmbedUnimplementedTestServer()
}

// UnimplementedTestServer must be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (UnimplementedTestServer) TransIn(context.Context, *Req) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransIn not implemented")
}
func (UnimplementedTestServer) TransInCompensate(context.Context, *Req) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransInCompensate not implemented")
}
func (UnimplementedTestServer) TransOut(context.Context, *Req) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransOut not implemented")
}
func (UnimplementedTestServer) TransOutCompensate(context.Context, *Req) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransOutCompensate not implemented")
}
func (UnimplementedTestServer) mustEmbedUnimplementedTestServer() {}

// UnsafeTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServer will
// result in compilation errors.
type UnsafeTestServer interface {
	mustEmbedUnimplementedTestServer()
}

func RegisterTestServer(s grpc.ServiceRegistrar, srv TestServer) {
	s.RegisterService(&Test_ServiceDesc, srv)
}

func _Test_TransIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).TransIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Test/TransIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).TransIn(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_TransInCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).TransInCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Test/TransInCompensate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).TransInCompensate(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_TransOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).TransOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Test/TransOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).TransOut(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_TransOutCompensate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).TransOutCompensate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Test/TransOutCompensate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).TransOutCompensate(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// Test_ServiceDesc is the grpc.ServiceDesc for Test service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Test_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransIn",
			Handler:    _Test_TransIn_Handler,
		},
		{
			MethodName: "TransInCompensate",
			Handler:    _Test_TransInCompensate_Handler,
		},
		{
			MethodName: "TransOut",
			Handler:    _Test_TransOut_Handler,
		},
		{
			MethodName: "TransOutCompensate",
			Handler:    _Test_TransOutCompensate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
