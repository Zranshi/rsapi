package logic

import (
	"context"
	"rsapi/model"
	"rsapi/util"

	"rsapi/mall/rpc/cart/cart"
	"rsapi/mall/rpc/cart/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushLogic {
	return &PushLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushLogic) Push(in *cart.PushReq) (*cart.PushRes, error) {
	resp := new(cart.PushRes)
	m := &model.MallCart{
		UserId:  in.UserId,
		GoodsId: in.GoodsId,
		Count:   in.Count,
	}
	if result, err := l.svcCtx.CartM.Insert(l.ctx, m); err != nil {
		return resp, util.DbErr(err)
	} else {
		resp.Id, err = result.LastInsertId()
		if err != nil {
			return resp, util.DbErr(err)
		}
	}
	return resp, nil
}
