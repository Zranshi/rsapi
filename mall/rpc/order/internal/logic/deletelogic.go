package logic

import (
	"context"

	"rsapi/mall/rpc/order/internal/svc"
	"rsapi/mall/rpc/order/order"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *order.DeleteReq) (*order.DeleteRes, error) {
	resp := &order.DeleteRes{Item: new(order.OrderItem)}
	o, err := l.svcCtx.OrderM.FindOne(l.ctx, in.Id)
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
	err = l.svcCtx.OrderM.Delete(l.ctx, in.Id)
	if err != nil {
		return resp, util.DbErr(err)
	}
	return resp, nil
}
