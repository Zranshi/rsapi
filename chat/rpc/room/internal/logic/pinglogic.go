package logic

import (
	"context"

	"rsapi/chat/rpc/room/internal/svc"
	"rsapi/chat/rpc/room/room"

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

func (l *PingLogic) Ping(in *room.Request) (*room.Response, error) {
	// todo: add your logic here and delete this line

	return &room.Response{}, nil
}
