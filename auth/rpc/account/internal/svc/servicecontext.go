package svc

import (
	"rsapi/auth/model"
	"rsapi/auth/rpc/account/internal/config"
	"rsapi/auth/rpc/token/tokenclient"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	RdbConn  sqlx.SqlConn
	UserM    model.AuthUserModel
	RoleM    model.AuthRoleModel
	TokenRpc tokenclient.Token
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:   c,
		RdbConn:  conn,
		UserM:    model.NewAuthUserModel(conn, c.MySQL.Cache),
		RoleM:    model.NewAuthRoleModel(conn, c.MySQL.Cache),
		TokenRpc: tokenclient.NewToken(zrpc.MustNewClient(c.TokenRpc)),
	}
}
