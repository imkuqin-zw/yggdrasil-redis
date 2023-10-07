package main

import (
	"context"
	"fmt"
	"time"

	xredis "github.com/imkuqin-zw/yggdrasil-redis"
	"github.com/imkuqin-zw/yggdrasil/pkg/config"
	"github.com/imkuqin-zw/yggdrasil/pkg/config/source/file"
	"github.com/imkuqin-zw/yggdrasil/pkg/logger"
)

func main() {
	if err := config.LoadSource(file.NewSource("./config.yaml", false)); err != nil {
		logger.FatalField("fault to load config file", logger.Err(err))
	}
	cache := xredis.NewRedis("center")
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	if err := cache.SAdd(ctx, "test_fdsaf", "fdsaf", "fdsafsdf").Err(); err != nil {
		panic(err)
	}
	result, err := cache.SMembers(ctx, "test_fdsaf").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
