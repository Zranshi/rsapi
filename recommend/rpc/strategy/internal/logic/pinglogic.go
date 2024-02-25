package logic

import (
	"context"

	"rsapi/recommend/rpc/strategy/internal/svc"
	"rsapi/recommend/rpc/strategy/strategy"

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

func (l *PingLogic) Ping(in *strategy.Request) (*strategy.Response, error) {
	// todo: add your logic here and delete this line

	return &strategy.Response{}, nil
}
