// Code generated by goctl. DO NOT EDIT.
// Source: formatter.proto

package server

import (
	"context"

	"rsapi/devtoys/rpc/formatter/formatter"
	"rsapi/devtoys/rpc/formatter/internal/logic"
	"rsapi/devtoys/rpc/formatter/internal/svc"
)

type FormatterServer struct {
	svcCtx *svc.ServiceContext
	formatter.UnimplementedFormatterServer
}

func NewFormatterServer(svcCtx *svc.ServiceContext) *FormatterServer {
	return &FormatterServer{
		svcCtx: svcCtx,
	}
}

func (s *FormatterServer) Ping(ctx context.Context, in *formatter.Request) (*formatter.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
