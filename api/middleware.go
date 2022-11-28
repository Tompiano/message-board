package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HandlerFunc() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("中间件")
	}
}
