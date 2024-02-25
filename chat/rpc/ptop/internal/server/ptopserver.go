// Code generated by goctl. DO NOT EDIT.
// Source: ptop.proto

package server

import (
	"context"

	"rsapi/chat/rpc/ptop/internal/logic"
	"rsapi/chat/rpc/ptop/internal/svc"
	"rsapi/chat/rpc/ptop/ptop"
)

type PtopServer struct {
	svcCtx *svc.ServiceContext
	ptop.UnimplementedPtopServer
}

func NewPtopServer(svcCtx *svc.ServiceContext) *PtopServer {
	return &PtopServer{
		svcCtx: svcCtx,
	}
}

func (s *PtopServer) Ping(ctx context.Context, in *ptop.Request) (*ptop.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}