package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CacheControl(d time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", d/time.Second))
		c.Next()
	}
}
