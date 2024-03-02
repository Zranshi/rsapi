package logic

import (
	"context"
	"rsapi/model"
	"rsapi/util"
	"strings"

	"rsapi/mall/rpc/goods/goods"
	"rsapi/mall/rpc/goods/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *goods.AddReq) (*goods.AddRes, error) {
	resp := new(goods.AddRes)
	g := &model.MallGoods{}
	g.Name = in.Name
	err := g.Description.Scan(in.Description)
	if err != nil {
		return nil, util.IllegalArgumentErr("description")
	}
	g.Images = strings.Join(in.Images, ";")
	result, err := l.svcCtx.GoodsM.Insert(l.ctx, g)
	if err != nil {
		return nil, util.DbErr(err)
	} else if id, err := result.LastInsertId(); err != nil {
		return nil, util.DbErr(err)
	} else {
		resp.Id = id
	}
	return resp, nil
}
