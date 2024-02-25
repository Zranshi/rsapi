package logic

import (
	"context"

	"rsapi/grab/rpc/storer/internal/svc"
	"rsapi/grab/rpc/storer/storer"

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

func (l *PingLogic) Ping(in *storer.Request) (*storer.Response, error) {
	// todo: add your logic here and delete this line

	return &storer.Response{}, nil
}
