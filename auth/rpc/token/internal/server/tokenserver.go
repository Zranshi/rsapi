// Code generated by goctl. DO NOT EDIT.
// Source: token.proto

package server

import (
	"context"

	"rsapi/auth/rpc/token/internal/logic"
	"rsapi/auth/rpc/token/internal/svc"
	"rsapi/auth/rpc/token/token"
)

type TokenServer struct {
	svcCtx *svc.ServiceContext
	token.UnimplementedTokenServer
}

func NewTokenServer(svcCtx *svc.ServiceContext) *TokenServer {
	return &TokenServer{
		svcCtx: svcCtx,
	}
}

func (s *TokenServer) Ping(ctx context.Context, in *token.Request) (*token.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}