package api

import "github.com/gin-gonic/gin"

func Register(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")
	if len(userName) > 6 || len(userName) <= 0 {
		c.JSON(200, "用户名长度不符合规范")
	} else {

	}
}
