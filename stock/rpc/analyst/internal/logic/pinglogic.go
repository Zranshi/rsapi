package logic

import (
	"context"

	"rsapi/stock/rpc/analyst/analyst"
	"rsapi/stock/rpc/analyst/internal/svc"

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

func (l *PingLogic) Ping(in *analyst.Request) (*analyst.Response, error) {
	// todo: add your logic here and delete this line

	return &analyst.Response{}, nil
}
