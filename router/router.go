package router

import (
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/log"
	"time"

	"github.com/SourceGlobalCDN/avatar-proxy/controller"
	"github.com/SourceGlobalCDN/avatar-proxy/middleware"
	cache "github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.HomepageHandler)

	// Gravatar
	{
		r.GET(":path", middleware.CacheControl(24*time.Hour), controller.AvatarInfoChecker, controller.AvatarInfoHandler)

		// Redirect requests from /gravatar/ to /avatar/
		r.GET("/gravatar/:code", middleware.CacheControl(24*365*time.Hour), controller.AvatarParser, controller.RedirectToAvatar)

		avatars := r.Group("/avatar", middleware.CacheControl(time.Hour*24*365))

		if store != nil {
			avatars.Use(cache.CacheByRequestURI(store, time.Hour*24*365, cache.WithPrefixKey("avatar"), cache.WithOnHitCache(func(c *gin.Context) {
				log.Log().Infof("Cache hit: %s", c.Request.RequestURI)
			})))
		}

		avatars.GET("", controller.AvatarHandler)
		avatars.GET(":code", controller.AvatarParser, controller.AvatarHandler)
	}

	return r
}
