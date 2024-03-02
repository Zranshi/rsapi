package svc

import (
	"rsapi/mall/rpc/order/internal/config"
	"rsapi/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config  config.Config
	RdbConn sqlx.SqlConn
	OrderM  model.MallOrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:  c,
		RdbConn: conn,
		OrderM:  model.NewMallOrderModel(conn, c.MySQL.Cache),
	}
}
