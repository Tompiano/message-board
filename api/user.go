package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/model"
	"message-board/service"
	"message-board/util"
)

func Register(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")
	question := c.PostForm("question")
	answer := c.PostForm("answer")
	if userName == "" || password == "" {
		util.RespParamErr(c)
		return
	}
	if question == "" || answer == "" {
		util.RespNormalErr(c, 200, "保密问题不能为空")
		return
	}
	u, err := service.SearchUserByUserName(userName)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RespInternalErr(c)
		return
	}
	if u.UserName != "" {
		util.RespNormalErr(c, 300, "账户已存在")
		return
	}
	err = service.CreateUser(model.User{
		UserName: userName,
		Password: password,
		Question: question,
		Answer:   answer,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
func Login(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")
	if userName == "" || password == "" {
		util.RespParamErr(c)
		return
	}

	u, err := service.SearchUserByUserName(userName)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespNormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
		}
		return
	}
	if u.Password != password {
		util.RespNormalErr(c, 20001, "密码错误")
		return
	}
	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
}
func Forget(c *gin.Context) {
	Question := c.PostForm("question")
	Answer := c.PostForm("answer")
	if Question == "" || Answer == "" {
		util.RespParamErr(c)
		return
	}
	_, err := service.ForgetPassword(Question, Answer)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespNormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user's question and answer error:%v", err)
			util.RespInternalErr(c)
		}
		return
	}
	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
}

func Modify(c *gin.Context) {
	//我希望的是登录之后才能修改密码
	userName := c.PostForm("username")
	password := c.PostForm("password") //获取修改的密码
	//对输入的密码有一定的规范
	if password == "" {
		util.RespParamErr(c)
		return
	}
	//根据姓名的位置将新密码插入数据库
	err := service.ModifyUser(model.User{
		UserName: userName,
		Password: password,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
