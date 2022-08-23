package logic

import (
	"context"
	"fmt"

	"demo/go-zero/rpc/internal/svc"
	"demo/go-zero/rpc/types/test"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransInLogic {
	return &TransInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransInLogic) TransIn(in *test.Req) (*test.Reply, error) {
	// todo: add your logic here and delete this line
	fmt.Println("TransIn")
	return &test.Reply{}, nil
}
