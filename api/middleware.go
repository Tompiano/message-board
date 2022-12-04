package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/util"
)

func HandlerFunc() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("中间件")
	}
}
func AuthorHandler() gin.HandlerFunc {
	return middle
}
func middle(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	//验证token格式是否正确
	if tokenString == "" || !string.HasPrefix(tokenString, "Bearer") {
		util.RespNormalErr(c, 400, "权限不足")
		return
	}
	//验证通过，提取有效部分，除去Bearer
	tokenString = tokenString[7:]
	c.Next() //进行下一步
}
