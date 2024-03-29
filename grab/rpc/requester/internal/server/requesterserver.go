// Code generated by goctl. DO NOT EDIT.
// Source: requester.proto

package server

import (
	"context"

	"rsapi/grab/rpc/requester/internal/logic"
	"rsapi/grab/rpc/requester/internal/svc"
	"rsapi/grab/rpc/requester/requester"
)

type RequesterServer struct {
	svcCtx *svc.ServiceContext
	requester.UnimplementedRequesterServer
}

func NewRequesterServer(svcCtx *svc.ServiceContext) *RequesterServer {
	return &RequesterServer{
		svcCtx: svcCtx,
	}
}

func (s *RequesterServer) Ping(ctx context.Context, in *requester.Request) (*requester.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
