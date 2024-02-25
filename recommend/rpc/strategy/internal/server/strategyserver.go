// Code generated by goctl. DO NOT EDIT.
// Source: strategy.proto

package server

import (
	"context"

	"rsapi/recommend/rpc/strategy/internal/logic"
	"rsapi/recommend/rpc/strategy/internal/svc"
	"rsapi/recommend/rpc/strategy/strategy"
)

type StrategyServer struct {
	svcCtx *svc.ServiceContext
	strategy.UnimplementedStrategyServer
}

func NewStrategyServer(svcCtx *svc.ServiceContext) *StrategyServer {
	return &StrategyServer{
		svcCtx: svcCtx,
	}
}

func (s *StrategyServer) Ping(ctx context.Context, in *strategy.Request) (*strategy.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
