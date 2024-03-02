package logic

import (
	"context"
	"rsapi/util"

	"rsapi/mall/rpc/cart/cart"
	"rsapi/mall/rpc/cart/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PopLogic {
	return &PopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PopLogic) Pop(in *cart.PopReq) (*cart.PopRes, error) {
	resp := new(cart.PopRes)
	if in.Count <= 0 {
		err := l.svcCtx.CartM.Delete(l.ctx, in.Id)
		if err != nil {
			return resp, util.DbErr(err)
		}
		resp.Count = 0
	} else {
		m, err := l.svcCtx.CartM.FindOne(l.ctx, in.Id)
		if err != nil {
			return resp, util.DbErr(err)
		}
		m.Count = in.Count
		err = l.svcCtx.CartM.Update(l.ctx, m)
		if err != nil {
			return resp, util.DbErr(err)
		}
		resp.Count = in.Count
	}
	return resp, nil
}
