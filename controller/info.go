package controller

import (
	"strings"

	"github.com/SourceGlobalCDN/avatar-proxy/pkg/serializer"
	"github.com/SourceGlobalCDN/avatar-proxy/service/avatar"
	"github.com/gin-gonic/gin"
)

func AvatarInfoChecker(c *gin.Context) {
	code := c.Param("path")
	if path := strings.Split(code, "."); len(path) > 1 {
		code = path[0]
		c.Set("format", path[1])
	}

	code = strings.ToLower(code)
	if len(code) != 32 {
		c.AbortWithStatusJSON(404, serializer.NotFoundError())
		return
	}
	c.Set("code", code)

	c.Next()
}

func AvatarInfoHandler(c *gin.Context) {
	code := c.GetString("code")
	format := c.GetString("format")

	if code == "" || len(code) != 32 {
		c.Next()
		return
	}

	client := avatar.NewFactory()
	avatarInfo, err := client.GetAvatarInfo(code)
	if err != nil {
		c.AbortWithStatusJSON(404, serializer.NotFoundError())
	}

	switch format {
	case "json":
		c.JSON(200, avatarInfo)
	case "xml":
		c.XML(200, avatarInfo)
	case "jsonp":
		c.JSONP(200, avatarInfo)
	case "yaml":
		c.YAML(200, avatarInfo)
	default:
		c.JSON(200, avatarInfo)
	}
}
