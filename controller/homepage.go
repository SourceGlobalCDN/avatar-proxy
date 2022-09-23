package controller

import "github.com/gin-gonic/gin"

func HomepageHandler(c *gin.Context) {
	c.Redirect(301, "https://www.sourcegcdn.com")
	c.Abort()
}
