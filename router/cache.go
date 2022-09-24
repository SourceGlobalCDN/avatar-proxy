package router

import (
	"fmt"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/env"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/log"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/go-redis/redis/v8"
	"time"
)

var store persist.CacheStore

func InitCache() {
	if !env.CacheConfig.Enabled {
		log.Log().Info("Gin cache is disabled")
		return
	}

	switch env.CacheConfig.Mode {
	case "redis":
		store = persist.NewRedisStore(redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", env.CacheConfig.Host, env.CacheConfig.Port),
			Password: env.CacheConfig.Password,
			DB:       env.CacheConfig.Database,
		}))
	case "memory":
		store = persist.NewMemoryStore(time.Minute * 10)
	default:
		log.Log().Fatal("Invalid cache mode")
	}
}
