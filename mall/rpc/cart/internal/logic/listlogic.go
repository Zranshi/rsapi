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

	err := l.svcCtx.RdbConn.QueryRowsCtx(l.ctx, &resp.Items, "SELECT * FROM mall_cart WHERE user_id = ?", in.UserId)
	if err != nil {
		return resp, util.DbErr(err)
	}
	return resp, nil
}
