package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AuthUserModel = (*customAuthUserModel)(nil)

type (
	// AuthUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthUserModel.
	AuthUserModel interface {
		authUserModel
	}

	customAuthUserModel struct {
		*defaultAuthUserModel
	}
)

// NewAuthUserModel returns a model for the database table.
func NewAuthUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AuthUserModel {
	return &customAuthUserModel{
		defaultAuthUserModel: newAuthUserModel(conn, c, opts...),
	}
}
