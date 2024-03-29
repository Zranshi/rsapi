// Code generated by goctl. DO NOT EDIT.
// Source: displayer.proto

package server

import (
	"context"

	"rsapi/stock/rpc/displayer/displayer"
	"rsapi/stock/rpc/displayer/internal/logic"
	"rsapi/stock/rpc/displayer/internal/svc"
)

type DisplayerServer struct {
	svcCtx *svc.ServiceContext
	displayer.UnimplementedDisplayerServer
}

func NewDisplayerServer(svcCtx *svc.ServiceContext) *DisplayerServer {
	return &DisplayerServer{
		svcCtx: svcCtx,
	}
}

func (s *DisplayerServer) Ping(ctx context.Context, in *displayer.Request) (*displayer.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
