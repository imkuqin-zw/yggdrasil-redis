package xredis

import (
	"github.com/go-redis/redis"
	"github.com/imkuqin-zw/yggdrasil/pkg/config"
	"github.com/imkuqin-zw/yggdrasil/pkg/logger"
)

type Redis interface {
	redis.UniversalClient
}

func NewRedis(name string) Redis {
	cfg := new(Config)
	if err := config.Get("redis." + name).Scan(cfg); err != nil {
		logger.FatalFiled("fault to load redis config", logger.Err(err))
	}
	cli := redis.NewUniversalClient(&cfg.Universal)
	if err := cli.Ping().Err(); err != nil {
		logger.FatalFiled("fault to connect redis", logger.Err(err))
	}
	return cli
}
