package logic

import (
	"context"

	"rsapi/devtoys/rpc/codec/codec"
	"rsapi/devtoys/rpc/codec/internal/svc"

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

func (l *PingLogic) Ping(in *codec.Request) (*codec.Response, error) {
	// todo: add your logic here and delete this line

	return &codec.Response{}, nil
}
