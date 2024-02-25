package logic

import (
	"context"

	"rsapi/task/rpc/schedule/internal/svc"
	"rsapi/task/rpc/schedule/schedule"

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

func (l *PingLogic) Ping(in *schedule.Request) (*schedule.Response, error) {
	// todo: add your logic here and delete this line

	return &schedule.Response{}, nil
}
