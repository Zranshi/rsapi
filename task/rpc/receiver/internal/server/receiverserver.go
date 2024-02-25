// Code generated by goctl. DO NOT EDIT.
// Source: receiver.proto

package server

import (
	"context"

	"rsapi/task/rpc/receiver/internal/logic"
	"rsapi/task/rpc/receiver/internal/svc"
	"rsapi/task/rpc/receiver/receiver"
)

type ReceiverServer struct {
	svcCtx *svc.ServiceContext
	receiver.UnimplementedReceiverServer
}

func NewReceiverServer(svcCtx *svc.ServiceContext) *ReceiverServer {
	return &ReceiverServer{
		svcCtx: svcCtx,
	}
}

func (s *ReceiverServer) Ping(ctx context.Context, in *receiver.Request) (*receiver.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
