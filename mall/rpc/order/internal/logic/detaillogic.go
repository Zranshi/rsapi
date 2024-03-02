package logic

import (
	"context"

	"rsapi/mall/rpc/order/internal/svc"
	"rsapi/mall/rpc/order/order"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *order.DetailReq) (*order.DetailRes, error) {
	resp := &order.DetailRes{Item: &order.OrderItem{}}
	o, err := l.svcCtx.OrderM.FindOne(l.ctx, in.Id)
	if err != nil {
		return resp, util.DbErr(err)
	}
	resp.Item.Id = o.Id
	resp.Item.Count = o.Count
	resp.Item.GoodsId = o.GoodsId
	resp.Item.UserId = o.UserId
	resp.Item.Status = o.Status
	return resp, nil
}
