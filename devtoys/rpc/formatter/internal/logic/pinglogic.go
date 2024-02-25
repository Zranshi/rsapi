package logic

import (
	"context"

	"rsapi/devtoys/rpc/formatter/formatter"
	"rsapi/devtoys/rpc/formatter/internal/svc"

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

func (l *PingLogic) Ping(in *formatter.Request) (*formatter.Response, error) {
	// todo: add your logic here and delete this line

	return &formatter.Response{}, nil
}
