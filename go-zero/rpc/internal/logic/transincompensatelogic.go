package logic

import (
	"context"
	"fmt"

	"demo/go-zero/rpc/internal/svc"
	"demo/go-zero/rpc/types/test"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransInCompensateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransInCompensateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransInCompensateLogic {
	return &TransInCompensateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransInCompensateLogic) TransInCompensate(in *test.Req) (*test.Reply, error) {
	// todo: add your logic here and delete this line
	fmt.Println("TransInCompensate")
	return &test.Reply{}, nil
}
