package controller

import (
	"fmt"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/blacklist"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/env"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/log"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/serializer"
	"github.com/SourceGlobalCDN/avatar-proxy/service/avatar"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

func AvatarHandler(c *gin.Context) {
	code := c.Param("code")

	if blacklist.CheckGravatar(code) {
		log.Log().Infof("Blocked gravatar code: %s", code)
		c.Status(404)
		return
	}

	var avatarPayload avatar.Payload
	err := c.ShouldBindQuery(&avatarPayload)
	if err != nil {
		log.Log().Errorf("Failed to bind query: %s", err)
		c.JSON(400, serializer.BadRequestError())
		return
	}

	u, _ := url.Parse(env.ProxyConfig.Remote)
	c.Header("X-Avatar-Proxy", u.String())

	u.Path = fmt.Sprintf("%s/%s", strings.TrimSuffix(u.Path, "/"), code)
	c.Header("Link", fmt.Sprintf("<%s>; rel=\"canonical\" as=\"image\"", u.String()))

	c.Header("X-Origin-Url", fmt.Sprintf("https://www.gravatar.com/avatar/%s", code))

	client := avatar.NewFactory()
	avatarCloser, length, err := client.GetAvatar(code, avatarPayload)
	if err != nil {
		log.Log().Errorf("Failed to get avatarCloser: %s", err)
		c.JSON(500, serializer.InternalServerError())
		return
	}

	c.DataFromReader(200, int64(length), "image/png", *avatarCloser, nil)
}
