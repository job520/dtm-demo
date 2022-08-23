// Code generated by goctl. DO NOT EDIT!
// Source: test.proto

package server

import (
	"context"

	"demo/go-zero/rpc/internal/logic"
	"demo/go-zero/rpc/internal/svc"
	"demo/go-zero/rpc/types/test"
)

type TestServer struct {
	svcCtx *svc.ServiceContext
	test.UnimplementedTestServer
}

func NewTestServer(svcCtx *svc.ServiceContext) *TestServer {
	return &TestServer{
		svcCtx: svcCtx,
	}
}

func (s *TestServer) TransIn(ctx context.Context, in *test.Req) (*test.Reply, error) {
	l := logic.NewTransInLogic(ctx, s.svcCtx)
	return l.TransIn(in)
}

func (s *TestServer) TransInCompensate(ctx context.Context, in *test.Req) (*test.Reply, error) {
	l := logic.NewTransInCompensateLogic(ctx, s.svcCtx)
	return l.TransInCompensate(in)
}

func (s *TestServer) TransOut(ctx context.Context, in *test.Req) (*test.Reply, error) {
	l := logic.NewTransOutLogic(ctx, s.svcCtx)
	return l.TransOut(in)
}

func (s *TestServer) TransOutCompensate(ctx context.Context, in *test.Req) (*test.Reply, error) {
	l := logic.NewTransOutCompensateLogic(ctx, s.svcCtx)
	return l.TransOutCompensate(in)
}
