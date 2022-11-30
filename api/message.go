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
	//查看留言
	MessageId, _ := strconv.ParseInt(c.PostForm("MessageId"), 10, 64)
	AuthorId, _ := strconv.ParseInt(c.PostForm("AuthorId"), 10, 64)
	ReceiveId, _ := strconv.ParseInt(c.PostForm("ReceiveId"), 10, 64)
	Detail := c.PostForm("Detail")
	//判断各个数据是否为空
	if Detail == "" || MessageId == 0 || AuthorId == 0 || ReceiveId == 0 {
		util.RespParamErr(c)
		return
	}
	//查询留言id对应的detail
	m, err := service.SearchDetail(MessageId, AuthorId, ReceiveId)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	if m.MessageId != MessageId && m.AuthorId != AuthorId && m.ReceiveId != ReceiveId {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)

}
func Update(c *gin.Context) {
	//获取用户原来的各个id和其修改后的留言内容即detail
	MessageId, _ := strconv.ParseInt(c.PostForm("MessageId"), 10, 64)
	AuthorId, _ := strconv.ParseInt(c.PostForm("AuthorId"), 10, 64)
	ReceiveId, _ := strconv.ParseInt(c.PostForm("ReceiveId"), 10, 64)
	Detail := c.PostForm("Detail")
	//判断各个数据是否为空
	if Detail == "" || MessageId == 0 || AuthorId == 0 || ReceiveId == 0 {
		util.RespParamErr(c)
		return
	}
	//更新留言id对应的detail
	err := service.UpdateMessage(model.Message{
		Detail: Detail,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)

}
func Delete(c *gin.Context) {
	//获取用户的留言以及各个id
	Detail := "该留言已删除"
	MessageId, _ := strconv.ParseInt(c.PostForm("MessageId"), 10, 64)
	if MessageId == 0 {
		util.RespParamErr(c)
		return
	}
	/*我觉得想要仅仅删除该留言而不删除其评论则将各个id都利用起来
	那么删除的就是各项id指向的唯一的那条留言
	至于留下”已删除“字样，我觉得就不用delete，而用update
	将原来的数据修改为”该留言已删除“*/
	err := service.DeleteMessage(model.Message{
		Detail: Detail,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}

}
func Like(c *gin.Context) {
	LikeNumber, _ := strconv.ParseInt(c.PostForm("LikeNumber"), 10, 64)
	userName := c.PostForm("userName")
	MessageId, _ := strconv.ParseInt(c.PostForm("MessageId"), 10, 64)
	AuthorId, _ := strconv.ParseInt(c.PostForm("AuthorId"), 10, 64)
	ReceiveId, _ := strconv.ParseInt(c.PostForm("ReceiveId"), 10, 64)
	if userName == "" || MessageId == 0 || AuthorId == 0 || ReceiveId == 0 {
		util.RespParamErr(c)
		return
	}
	err := service.LikeIncrease(model.Message{
		MessageId:  MessageId,
		AuthorId:   AuthorId,
		ReceiveId:  ReceiveId,
		LikeNumber: LikeNumber,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
