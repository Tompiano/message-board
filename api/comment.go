package api

import (
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/util"
	"strconv"
	"strings"
)

func SendComment(c *gin.Context) {
	//获取用户写的评论和用户信息
	content := c.PostForm("content")
	userName := c.PostForm("userName")
	MessageId, _ := strconv.ParseInt(c.PostForm("messageId"), 10, 64)
	parentId, _ := strconv.ParseInt(c.PostForm("parentId"), 10, 64)
	ChildId, _ := strconv.ParseInt(c.PostForm("ChildId"), 10, 64)
	parentUserId, _ := strconv.ParseInt(c.PostForm("parentUserId"), 10, 64)
	//检查各个数据是否为空
	if content == "" || userName == "" || MessageId == 0 || parentId == 0 || ChildId == 0 || parentUserId == 0 {
		util.RespParamErr(c)
		return
	}
	//评论内容小于500字，且不能包含敏感词汇
	if len(content) > 500 {
		util.RespNormalErr(c, 200, "内容不超过500字")
		return
	}
	jud := strings.Contains(content, "傻逼")
	if !jud {
		util.RespNormalErr(c, 200, "内容不能包含敏感词汇")
	}
	//插入评论或回复
	err := service.CreateComment(model.Comment{
		Content:      content,
		UserName:     userName,
		MessageId:    MessageId,
		ParentId:     parentId,
		ChildId:      ChildId,
		ParentUserId: parentUserId,
	})

	if err != nil {
		util.RespInternalErr(c)
		return
	}

}
func LookComment(c *gin.Context) {
	//因为在查看留言和评论的时候是可以将整个留言板上的全部看完的，所以需要全部查询出来
	/*步骤：
	1.先要找到message，（message根据MessageId来查找）
	2.找ParentId为0的点
	3.根据不同的ParentUserId和ParentId对应来找到他们下方的评论
	4.再根据不同的childId来找到它们的回复
	*/
	MessageId, _ := strconv.ParseInt(c.PostForm("messageId"), 10, 64)
	if MessageId == 0 {
		util.RespParamErr(c)
		return
	}

}
func ModifyComment(c *gin.Context) {

}
func DeleteComment(c *gin.Context) {

}
