package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var OK = respTemplate{
	Status: 200,
	Info:   "success",
}
var ParamError = respTemplate{
	Status: 300,
	Info:   "params error",
}
var InternalError = respTemplate{
	Status: 300,
	Info:   "internal error",
}

func RespNormalErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})

}

func RespOK(c *gin.Context) {
	c.JSON(http.StatusOK, OK)
}
func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
}
func RespInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalError)
}

func RespLike(c *gin.Context) {
	c.JSON(http.StatusOK, "还没有人点赞哦，快来点赞吧")
}
