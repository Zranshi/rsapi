package logic

import (
	"context"
	"rsapi/util"

	"rsapi/mall/rpc/cart/cart"
	"rsapi/mall/rpc/cart/internal/svc"

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

func (l *ListLogic) List(in *cart.ListReq) (*cart.ListRes, error) {
	resp := new(cart.ListRes)
	resp.Items = make([]*cart.CartItem, 0)

	rows := make([]struct {
		Id      int64 `db:"id"`
		GoodsId int64 `db:"goods_id"`
		Count   int64 `db:"count"`
	}, 0)

	err := l.svcCtx.RdbConn.QueryRowsCtx(l.ctx, &rows, "SELECT * FROM mall_cart WHERE user_id = ?", in.UserId)
	if err != nil {
		return resp, util.DbErr(err)
	}
	resp.Items = make([]*cart.CartItem, len(rows))
	for i, row := range rows {
		resp.Items[i] = &cart.CartItem{
			Id:      row.Id,
			GoodsId: row.GoodsId,
			Count:   row.Count,
		}
	}
	return resp, nil
}
