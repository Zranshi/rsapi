package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"rsapi/mall/rpc/cart/internal/config"
	"rsapi/model"
)

type ServiceContext struct {
	Config  config.Config
	RdbConn sqlx.SqlConn
	CartM   model.MallCartModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:  c,
		RdbConn: conn,
		CartM:   model.NewMallCartModel(conn, c.MySQL.Cache),
	}
}
