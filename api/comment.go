package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"message-board/model"
	"message-board/service"
	"message-board/util"
	"strings"
)

func SendComment(c *gin.Context) {
	//获取用户写的评论和用户信息
	content := c.PostForm("content")
	userName := c.PostForm("userName")
	pId := c.PostForm("pId")
	//检查各个数据是否为空
	if content == "" || userName == "" || pId == "" {
		log.Print("有参数为空")
		util.RespParamErr(c)
		return
	}
	//评论内容小于500字，且不能包含敏感词汇
	if len(content) > 500 {
		util.RespNormalErr(c, 200, "内容不超过500字")
		return
	}
	jud := strings.Contains(content, "傻逼")
	if jud == true {
		util.RespNormalErr(c, 200, "内容不能包含敏感词汇")
	}
	err := service.CreateComment(model.Comment{
		UserName: userName,
		PId:      pId,
		Content:  content,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)

}
func LookComment(c *gin.Context) {

	userName := c.PostForm("userName")
	pId := c.PostForm("pId")
	if userName == "" {
		util.RespParamErr(c)
		return
	}
	err := service.SearchComment(pId, userName)
	if err != nil {
		util.RespInternalErr(c)
	}

	util.RespOK(c)
}
func ModifyComment(c *gin.Context) {
	content := c.PostForm("content")
	userName := c.PostForm("userName")
	pId := c.PostForm("pId")

	if content == "" || userName == "" || pId == "" {
		util.RespParamErr(c)
		return
	}
	err := service.ModifyComment(model.Comment{
		UserName: userName,
		PId:      pId,
		Content:  content,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
func DeleteComment(c *gin.Context) {
	content := "该评论已删除"
	userName := c.PostForm("userName")
	pId := c.PostForm("pId")

	if userName == "" || pId == "" {
		util.RespParamErr(c)
		return
	}
	err := service.ModifyComment(model.Comment{
		UserName: userName,
		PId:      pId,
		Content:  content,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
