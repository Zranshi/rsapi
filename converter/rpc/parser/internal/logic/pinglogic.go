package logic

import (
	"context"

	"rsapi/converter/rpc/parser/internal/svc"
	"rsapi/converter/rpc/parser/parser"

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

func (l *PingLogic) Ping(in *parser.Request) (*parser.Response, error) {
	// todo: add your logic here and delete this line

	return &parser.Response{}, nil
}
