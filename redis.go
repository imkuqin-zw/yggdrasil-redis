package xredis

import (
	"context"
	"time"

	"github.com/imkuqin-zw/yggdrasil/pkg/config"
	"github.com/imkuqin-zw/yggdrasil/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type Redis interface {
	redis.UniversalClient
}

func NewRedis(name string) Redis {
	cfg := new(Config)
	if err := config.Get("redis." + name).Scan(cfg); err != nil {
		logger.FatalField("fault to load redis config", logger.Err(err))
	}
	redis.SetLogger(newLogging(cfg))
	cli := newUniversalClient(cfg)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()
	if err := cli.Ping(ctx).Err(); err != nil {
		logger.FatalField("fault to connect redis", logger.Err(err))
	}
	return cli
}

func newUniversalClient(cfg *Config) redis.UniversalClient {
	if cfg.Universal.MasterName != "" {
		return redis.NewFailoverClient(cfg.Universal.Failover())
	}
	if cfg.Cluster || len(cfg.Universal.Addrs) > 1 {
		return redis.NewClusterClient(cfg.Universal.Cluster())
	}
	return redis.NewClient(cfg.Universal.Simple())
}
