package logic

import (
	"context"
	"rsapi/mall/rpc/goods/goods"
	"rsapi/mall/rpc/goods/internal/svc"
	"rsapi/util"
	"strings"

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

func (l *ListLogic) List(in *goods.ListReq) (*goods.ListRes, error) {
	resp := new(goods.ListRes)
	rows := make([]*struct {
		Id          int64  `db:"id"`
		Name        string `db:"name"`
		Images      string `db:"images"`
		Description string `db:"description"`
	}, 0)
	err := l.svcCtx.RdbConn.QueryRowsCtx(l.ctx, &rows, "SELECT id, name, description, images FROM mall_goods")
	if err != nil {
		return nil, util.DbErr(err)
	}
	resp.Items = make([]*goods.GoodsItem, len(rows))
	for i, row := range rows {
		resp.Items[i] = &goods.GoodsItem{
			Id:          row.Id,
			Name:        row.Name,
			Description: row.Description,
			Images:      strings.Split(row.Images, ";"),
		}
	}
	return resp, nil
}
