package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MallCartModel = (*customMallCartModel)(nil)

type (
	// MallCartModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMallCartModel.
	MallCartModel interface {
		mallCartModel
	}

	customMallCartModel struct {
		*defaultMallCartModel
	}
)

// NewMallCartModel returns a model for the database table.
func NewMallCartModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MallCartModel {
	return &customMallCartModel{
		defaultMallCartModel: newMallCartModel(conn, c, opts...),
	}
}
