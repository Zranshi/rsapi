package logic

import (
	"context"

	"rsapi/mall/rpc/order/internal/svc"
	"rsapi/mall/rpc/order/order"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLogic {
	return &CancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelLogic) Cancel(in *order.CancelReq) (*order.CancelRes, error) {
	resp := &order.CancelRes{Item: new(order.OrderItem)}
	o, err := l.svcCtx.OrderM.FindOne(l.ctx, in.Id)
	if err != nil {
		return resp, util.DbErr(err)
	}
	o.Status = 2
	err = l.svcCtx.OrderM.Update(l.ctx, o)
	if err != nil {
		return resp, util.DbErr(err)
	}
	resp.Item = &order.OrderItem{
		Id:      o.Id,
		UserId:  o.UserId,
		GoodsId: o.GoodsId,
		Count:   o.Count,
		Status:  o.Status,
	}
	return resp, nil
}
