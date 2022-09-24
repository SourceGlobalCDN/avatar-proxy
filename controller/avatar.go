package controller

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/SourceGlobalCDN/avatar-proxy/pkg/blacklist"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/env"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/log"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/serializer"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/util"
	"github.com/SourceGlobalCDN/avatar-proxy/service/avatar"
	"github.com/gin-gonic/gin"
)

func AvatarParser(c *gin.Context) {
	code := c.Param("code")
	code = strings.ToLower(code)
	if len(code) != 32 {
		c.Set("code", "")
	} else {
		c.Set("code", code)
	}

	c.Next()
}

func AvatarHandler(c *gin.Context) {
	code := c.GetString("code")

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

	if avatarPayload.Size > 2048 {
		u := c.Request.URL
		q := u.Query()
		q.Set("s", "2048")
		u.RawQuery = q.Encode()

		c.Redirect(302, fmt.Sprintf("/avatar/%s?%s", code, u.RawQuery))
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
		if errors.Is(err, avatar.ErrNotFound) {
			c.JSON(404, serializer.NotFoundError())
			return
		}

		log.Log().Errorf("Failed to get avatar: %s", err)
		c.JSON(500, serializer.InternalServerError())
		return
	}

	data, err := io.ReadAll(*avatarCloser)
	if err != nil {
		log.Log().Errorf("Failed to read avatar: %s", err)
		c.JSON(500, serializer.InternalServerError())
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", code))
	c.Header("Content-Length", fmt.Sprintf("%d", length))

	md5 := util.MD5Hex(data)
	c.Header("ETag", fmt.Sprintf("W/\"%s\"", util.MD5Hex(data)))
	c.Header("Content-MD5", md5)

	c.Data(200, "image/jpeg", data)
}

func RedirectToAvatar(c *gin.Context) {
	code := c.GetString("code")
	c.Redirect(302, fmt.Sprintf("/avatar/%s?%s", code, c.Request.URL.RawQuery))
}
