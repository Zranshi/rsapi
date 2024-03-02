package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MallOrderModel = (*customMallOrderModel)(nil)

type (
	// MallOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMallOrderModel.
	MallOrderModel interface {
		mallOrderModel
	}

	customMallOrderModel struct {
		*defaultMallOrderModel
	}
)

// NewMallOrderModel returns a model for the database table.
func NewMallOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MallOrderModel {
	return &customMallOrderModel{
		defaultMallOrderModel: newMallOrderModel(conn, c, opts...),
	}
}
