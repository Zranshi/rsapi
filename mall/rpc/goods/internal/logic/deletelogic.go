package logic

import (
	"context"

	"rsapi/mall/rpc/goods/goods"
	"rsapi/mall/rpc/goods/internal/svc"
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

func (l *DeleteLogic) Delete(in *goods.DeleteReq) (*goods.DeleteRes, error) {
	resp := new(goods.DeleteRes)
	err := l.svcCtx.GoodsM.Delete(l.ctx, in.Id)
	if err != nil {
		return resp, util.DbErr(err)
	}
	resp.Id = in.Id
	return resp, nil
}
