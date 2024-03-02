package logic

import (
	"context"

	"rsapi/mall/rpc/order/internal/svc"
	"rsapi/mall/rpc/order/order"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *order.ListReq) (*order.ListRes, error) {
	resp := new(order.ListRes)
	resp.Items = make([]*order.OrderItem, 0)
	rows := make([]struct {
		Id      int64 `db:"id"`
		UserId  int64 `db:"user_id"`
		GoodsId int64 `db:"goods_id"`
		Count   int64 `db:"count"`
		Status  int64 `db:"status"`
	}, 0)
	err := l.svcCtx.RdbConn.QueryRowsCtx(l.ctx, &rows, "SELECT * FROM mall_order WHERE user_id = ?", in.UserId)
	if err != nil {
		return resp, util.DbErr(err)
	}
	resp.Items = make([]*order.OrderItem, len(rows))
	for i, row := range rows {
		resp.Items[i] = &order.OrderItem{
			Id:      row.Id,
			UserId:  row.UserId,
			GoodsId: row.GoodsId,
			Count:   row.Count,
			Status:  row.Status,
		}
	}
	return resp, nil
}
