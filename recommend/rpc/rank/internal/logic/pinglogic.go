package logic

import (
	"context"

	"rsapi/recommend/rpc/rank/internal/svc"
	"rsapi/recommend/rpc/rank/rank"

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

func (l *PingLogic) Ping(in *rank.Request) (*rank.Response, error) {
	// todo: add your logic here and delete this line

	return &rank.Response{}, nil
}
