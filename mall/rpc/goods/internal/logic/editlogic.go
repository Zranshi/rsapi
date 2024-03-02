package logic

import (
	"context"
	"rsapi/util"
	"strings"

	"rsapi/mall/rpc/goods/goods"
	"rsapi/mall/rpc/goods/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditLogic) Edit(in *goods.EditReq) (*goods.EditRes, error) {
	resp := new(goods.EditRes)
	g, err := l.svcCtx.GoodsM.FindOne(l.ctx, in.Id)
	if err != nil {
		return resp, util.DbErr(err)
	}
	if in.Name != nil {
		g.Name = *in.Name
	}
	if in.Description != nil {
		err := g.Description.Scan(*in.Description)
		if err != nil {
			return resp, util.IllegalArgumentErr("description")
		}
	}
	if in.Images != nil {
		g.Images = strings.Join(in.Images, ";")
	}

	err = l.svcCtx.GoodsM.Update(l.ctx, g)
	if err != nil {
		return resp, util.DbErr(err)
	}
	return resp, nil
}
