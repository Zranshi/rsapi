package logic

import (
	"context"

	"rsapi/converter/rpc/adapter/adapter"
	"rsapi/converter/rpc/adapter/internal/svc"

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

func (l *PingLogic) Ping(in *adapter.Request) (*adapter.Response, error) {
	// todo: add your logic here and delete this line

	return &adapter.Response{}, nil
}
