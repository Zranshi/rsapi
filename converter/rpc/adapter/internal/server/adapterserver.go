// Code generated by goctl. DO NOT EDIT.
// Source: adapter.proto

package server

import (
	"context"

	"rsapi/converter/rpc/adapter/adapter"
	"rsapi/converter/rpc/adapter/internal/logic"
	"rsapi/converter/rpc/adapter/internal/svc"
)

type AdapterServer struct {
	svcCtx *svc.ServiceContext
	adapter.UnimplementedAdapterServer
}

func NewAdapterServer(svcCtx *svc.ServiceContext) *AdapterServer {
	return &AdapterServer{
		svcCtx: svcCtx,
	}
}

func (s *AdapterServer) Ping(ctx context.Context, in *adapter.Request) (*adapter.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}