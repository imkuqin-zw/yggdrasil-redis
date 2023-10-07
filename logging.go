package xredis

import (
	"context"
	"fmt"

	"github.com/imkuqin-zw/yggdrasil/pkg/logger"
)

type logging struct {
	printf func(ctx context.Context, format string, v ...interface{})
}

func (l *logging) Printf(ctx context.Context, format string, v ...interface{}) {
	l.printf(ctx, format, v...)
}

func newLogging(cfg *Config) *logging {
	var f = func(ctx context.Context, format string, v ...interface{}) {}
	if cfg.Logger.Enable {
		if cfg.Logger.WithTrace {
			f = func(ctx context.Context, format string, v ...interface{}) {
				logger.ErrorField(fmt.Sprintf(format, v...), logger.Context(ctx))
			}
		} else {
			f = func(ctx context.Context, format string, v ...interface{}) {
				logger.ErrorField(fmt.Sprintf(format, v...))
			}
		}
	}
	return &logging{printf: f}
}
