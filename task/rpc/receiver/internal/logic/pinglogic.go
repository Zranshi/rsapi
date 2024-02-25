package logic

import (
	"context"

	"rsapi/task/rpc/receiver/internal/svc"
	"rsapi/task/rpc/receiver/receiver"

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

func (l *PingLogic) Ping(in *receiver.Request) (*receiver.Response, error) {
	// todo: add your logic here and delete this line

	return &receiver.Response{}, nil
}
