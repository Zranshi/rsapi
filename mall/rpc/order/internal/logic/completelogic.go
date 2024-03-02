package logic

import (
	"context"

	"rsapi/mall/rpc/order/internal/svc"
	"rsapi/mall/rpc/order/order"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteLogic {
	return &CompleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompleteLogic) Complete(in *order.CompleteReq) (*order.CompleteRes, error) {
	resp := &order.CompleteRes{Item: new(order.OrderItem)}
	o, err := l.svcCtx.OrderM.FindOne(l.ctx, in.Id)
	if err != nil {
		return resp, util.DbErr(err)
	}
	o.Status = 1
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
