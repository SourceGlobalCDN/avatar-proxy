package router

import (
	"time"

	"github.com/SourceGlobalCDN/avatar-proxy/controller"
	"github.com/SourceGlobalCDN/avatar-proxy/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.HomepageHandler)

	{
		avatars := r.Group("/avatar", middleware.CacheControl(time.Hour*24*365))

		avatars.GET("", controller.AvatarHandler)
		avatars.GET(":code", controller.AvatarHandler)
	}

	return r
}
