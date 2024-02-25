// Code generated by goctl. DO NOT EDIT.
// Source: schedule.proto

package server

import (
	"context"

	"rsapi/task/rpc/schedule/internal/logic"
	"rsapi/task/rpc/schedule/internal/svc"
	"rsapi/task/rpc/schedule/schedule"
)

type ScheduleServer struct {
	svcCtx *svc.ServiceContext
	schedule.UnimplementedScheduleServer
}

func NewScheduleServer(svcCtx *svc.ServiceContext) *ScheduleServer {
	return &ScheduleServer{
		svcCtx: svcCtx,
	}
}

func (s *ScheduleServer) Ping(ctx context.Context, in *schedule.Request) (*schedule.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
