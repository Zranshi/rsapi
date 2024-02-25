package logic

import (
	"context"

	"rsapi/oss/rpc/usage/internal/svc"
	"rsapi/oss/rpc/usage/usage"

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

func (l *PingLogic) Ping(in *usage.Request) (*usage.Response, error) {
	// todo: add your logic here and delete this line

	return &usage.Response{}, nil
}
