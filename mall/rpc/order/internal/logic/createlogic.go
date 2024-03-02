package logic

import (
	"context"

	"rsapi/mall/rpc/order/internal/svc"
	"rsapi/mall/rpc/order/order"
	"rsapi/model"
	"rsapi/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateReq) (*order.CreateRes, error) {
	resp := new(order.CreateRes)
	o := &model.MallOrder{
		UserId:  in.UserId,
		GoodsId: in.GoodsId,
		Count:   in.Count,
		Status:  0,
	}
	result, err := l.svcCtx.OrderM.Insert(l.ctx, o)
	if err != nil {
		return resp, util.DbErr(err)
	}
	resp.Id, err = result.LastInsertId()
	if err != nil {
		return resp, util.DbErr(err)
	}
	return resp, nil
}
