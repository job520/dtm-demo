package logic

import (
	"context"
	"fmt"

	"demo/go-zero/rpc/internal/svc"
	"demo/go-zero/rpc/types/test"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransOutLogic {
	return &TransOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransOutLogic) TransOut(in *test.Req) (*test.Reply, error) {
	// todo: add your logic here and delete this line
	fmt.Println("TransOut")
	return &test.Reply{}, nil
}
