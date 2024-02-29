package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AuthRoleModel = (*customAuthRoleModel)(nil)

type (
	// AuthRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthRoleModel.
	AuthRoleModel interface {
		authRoleModel
	}

	customAuthRoleModel struct {
		*defaultAuthRoleModel
	}
)

// NewAuthRoleModel returns a model for the database table.
func NewAuthRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AuthRoleModel {
	return &customAuthRoleModel{
		defaultAuthRoleModel: newAuthRoleModel(conn, c, opts...),
	}
}
