package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
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
	parentUserId, _ := strconv.ParseInt(c.PostForm("parentUserId"), 10, 64)
	//检查各个数据是否为空
	if content == "" || userName == "" || MessageId == 0 || parentUserId == 0 {
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
	if parentId == 0 {
		//直接向ParentComment中间插入数据
		err := service.CreateParentComment(model.Comment{
			Content:      content,
			UserName:     userName,
			MessageId:    MessageId,
			ParentId:     parentId,
			ParentUserId: parentUserId,
		})

		if err != nil {
			util.RespInternalErr(c)
			return
		}

	} else {
		ChildId, _ := strconv.ParseInt(c.PostForm("ChildId"), 10, 64)
		parentId = parentUserId
		//将回复的parentId改成评论的parentUserId后向ChildComment中间插入数据
		err := service.CreateChildComment(model.Comment{
			Content:   content,
			UserName:  userName,
			MessageId: MessageId,
			ParentId:  parentId,
			ChildId:   ChildId,
		})

		if err != nil {
			util.RespInternalErr(c)
			return
		}

	}
	util.RespOK(c)

}
func LookComment(c *gin.Context) {
	/*步骤：
	1.先要找到message，（message根据MessageId来查找）
	2.找到评论，即在ParentComment表里面找
	3.找到回复，根据ParentId和ChildId在ChildComment表里面找
	4.找到相应的评论或者回复以后就进行修改
	*/
	Content := c.PostForm("content")
	MessageId, _ := strconv.ParseInt(c.PostForm("messageId"), 10, 64)
	parentId, _ := strconv.ParseInt(c.PostForm("parentId"), 10, 64)
	ParentUserId, _ := strconv.ParseInt(c.PostForm("parentUserId"), 10, 64)
	ChildId, _ := strconv.ParseInt(c.PostForm("childId"), 10, 64)
	if MessageId == 0 || Content == "" {
		util.RespParamErr(c)
		return
	}
	if parentId == 0 {
		_, err := service.SearchParentComment(MessageId, ParentUserId)
		if err != nil {
			if err == sql.ErrNoRows {
				util.RespNormalErr(c, 300, "该评论不存在")
			} else {
				log.Printf("search parentComment error:%v", err)
				util.RespInternalErr(c)
			}
			return
		}
	} else {
		_, err := service.SearchChildComment(MessageId, parentId, ChildId)
		if err != nil {
			if err == sql.ErrNoRows {
				util.RespNormalErr(c, 300, "该回复不存在")
			} else {
				log.Printf("search childComment error:%v", err)
				util.RespInternalErr(c)
			}
			return
		}
	}
	util.RespOK(c)
}
func ModifyComment(c *gin.Context) {
	Content := c.PostForm("content")
	MessageId, _ := strconv.ParseInt(c.PostForm("messageId"), 10, 64)
	parentId, _ := strconv.ParseInt(c.PostForm("parentId"), 10, 64)
	ParentUserId, _ := strconv.ParseInt(c.PostForm("parentUserId"), 10, 64)
	ChildId, _ := strconv.ParseInt(c.PostForm("childId"), 10, 64)
	if Content == "" || MessageId == 0 || ParentUserId == 0 {
		util.RespParamErr(c)
		return
	}
	if parentId == 0 {
		err := service.ModifyParentComment(model.Comment{
			ParentId:     parentId,
			ParentUserId: ParentUserId,
			MessageId:    MessageId,
			Content:      Content,
		})
		if err != nil {
			util.RespInternalErr(c)
			return
		}
	} else {
		err := service.ModifyChildComment(model.Comment{
			ParentId:  parentId,
			ChildId:   ChildId,
			MessageId: MessageId,
			Content:   Content,
		})
		if err != nil {
			util.RespInternalErr(c)
			return
		}
	}
	util.RespOK(c)
}
func DeleteComment(c *gin.Context) {
	Content := "该评论已删除"
	MessageId, _ := strconv.ParseInt(c.PostForm("MessageId"), 10, 64)
	parentId, _ := strconv.ParseInt(c.PostForm("parentId"), 10, 64)
	ParentUserId, _ := strconv.ParseInt(c.PostForm("parentUserId"), 10, 64)
	ChildId, _ := strconv.ParseInt(c.PostForm("childId"), 10, 64)
	if MessageId == 0 || ParentUserId == 0 {
		util.RespParamErr(c)
		return
	}
	if parentId == 0 {
		err := service.ModifyParentComment(model.Comment{
			ParentId:     parentId,
			ParentUserId: ParentUserId,
			MessageId:    MessageId,
			Content:      Content,
		})
		if err != nil {
			util.RespInternalErr(c)
			return
		}
	} else {
		err := service.ModifyChildComment(model.Comment{
			ParentId:  parentId,
			ChildId:   ChildId,
			MessageId: MessageId,
			Content:   Content,
		})
		if err != nil {
			util.RespInternalErr(c)
			return
		}
	}
	util.RespOK(c)
}
