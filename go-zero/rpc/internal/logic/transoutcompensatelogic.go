package logic

import (
	"context"
	"fmt"

	"demo/go-zero/rpc/internal/svc"
	"demo/go-zero/rpc/types/test"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransOutCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransOutCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransOutCompensateLogic {
	return &TransOutCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransOutCompensateLogic) TransOutCompensate(in *test.Req) (*test.Reply, error) {
	// todo: add your logic here and delete this line
	fmt.Println("TransOutCompensate")
	return &test.Reply{}, nil
}
