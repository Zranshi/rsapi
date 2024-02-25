package logic

import (
	"context"

	"rsapi/ordertable/rpc/item/internal/svc"
	"rsapi/ordertable/rpc/item/item"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *item.Request) (*item.Response, error) {
	// todo: add your logic here and delete this line

	return &item.Response{}, nil
}
