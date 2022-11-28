package api

import (
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/util"
	"strconv"
)

func SendMessage(c *gin.Context) {
	//额直接用PostForm获取到的是字符串。。所以我就把它转成了int类型
	MessageId, _ := strconv.ParseInt(c.PostForm("MessageId"), 10, 64)
	AuthorId, _ := strconv.ParseInt(c.PostForm("AuthorId"), 10, 64)
	ReceiveId, _ := strconv.ParseInt(c.PostForm("ReceiveId"), 10, 64)
	Detail := c.PostForm("Detail")
	//判断各个数据是否为空
	if Detail == "" || MessageId == 0 || AuthorId == 0 || ReceiveId == 0 {
		util.RespParamErr(c)
		return
	}
	//将留言写进数据库
	err := service.CreateMessage(model.Message{
		Detail:    Detail,
		MessageId: MessageId,
		AuthorId:  AuthorId,
		ReceiveId: ReceiveId,
	})
	if err != nil {
		util.RespParamErr(c)
		return
	}
	util.RespOK(c)
}
func GetMessage(c *gin.Context) {
	MessageId, _ := strconv.ParseInt(c.PostForm("MessageId"), 10, 64)
	AuthorId, _ := strconv.ParseInt(c.PostForm("AuthorId"), 10, 64)
	ReceiveId, _ := strconv.ParseInt(c.PostForm("ReceiveId"), 10, 64)
	Detail := c.PostForm("Detail")
	//判断各个数据是否为空
	if Detail == "" || MessageId == 0 || AuthorId == 0 || ReceiveId == 0 {
		util.RespParamErr(c)
		return
	}
	//查询留言
	m, err := service.SearchDetail(MessageId, AuthorId, ReceiveId)
	if err != nil {
		util.RespInternalErr(m)
		return
	}
	util.RespOK(c)

}
func Update(c *gin.Context) {
	
}
