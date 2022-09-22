package router

import (
	"github.com/SourceGlobalCDN/avatar-proxy/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.HomepageHandler)

	return r
}
