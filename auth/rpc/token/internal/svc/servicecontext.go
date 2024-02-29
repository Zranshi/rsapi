package svc

import (
	"rsapi/auth/rpc/token/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	KvConn *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		KvConn: redis.MustNewRedis(c.RedisConf),
	}
}
