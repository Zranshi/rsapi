package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MallGoodsModel = (*customMallGoodsModel)(nil)

type (
	// MallGoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMallGoodsModel.
	MallGoodsModel interface {
		mallGoodsModel
	}

	customMallGoodsModel struct {
		*defaultMallGoodsModel
	}
)

// NewMallGoodsModel returns a model for the database table.
func NewMallGoodsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MallGoodsModel {
	return &customMallGoodsModel{
		defaultMallGoodsModel: newMallGoodsModel(conn, c, opts...),
	}
}
