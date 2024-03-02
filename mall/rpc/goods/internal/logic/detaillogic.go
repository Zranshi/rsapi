package logic

import (
	"context"
	"rsapi/util"
	"strings"

	"rsapi/mall/rpc/goods/goods"
	"rsapi/mall/rpc/goods/internal/svc"

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

func (l *DetailLogic) Detail(in *goods.DetailReq) (*goods.DetailRes, error) {
	resp := new(goods.DetailRes)
	resp.Item = new(goods.GoodsItem)
	g, err := l.svcCtx.GoodsM.FindOne(l.ctx, in.Id)
	if err != nil {
		return resp, util.DbErr(err)
	}
	resp.Item.Id = g.Id
	resp.Item.Name = g.Name
	if val, _ := g.Description.Value(); val != nil {
		resp.Item.Description = val.(string)
	}
	resp.Item.Images = strings.Split(g.Images, ";")
	return resp, nil
}
