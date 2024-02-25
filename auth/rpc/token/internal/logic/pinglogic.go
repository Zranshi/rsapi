package logic

import (
	"context"

	"rsapi/auth/rpc/token/internal/svc"
	"rsapi/auth/rpc/token/token"

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

func (l *PingLogic) Ping(in *token.Request) (*token.Response, error) {
	// todo: add your logic here and delete this line

	return &token.Response{}, nil
}
