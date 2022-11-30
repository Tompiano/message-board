package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"message-board/model"
	"message-board/service"
	"message-board/util"
)

func Register(c *gin.Context) {
	Id := c.PostForm("Id")
	userName := c.PostForm("UserName")
	password := c.PostForm("Password")
	//获取设置的保密问题以及答案
	question := c.PostForm("Question")
	answer := c.PostForm("Answer")
	if userName == "" || password == "" {
		util.RespParamErr(c)
		return
	}

	//入参校验：1.保密问题不能为空2.密码长度在10-20之间3.用户名长度在1-10之间
	if question == "" || answer == "" {
		util.RespNormalErr(c, 200, "保密问题不能为空")
		return
	}
	if len(password) < 10 || len(password) > 20 {
		util.RespNormalErr(c, 200, "密码长度不符合规范")
		return
	}
	if len(userName) < 1 || len(userName) > 10 {
		util.RespNormalErr(c, 200, "用户名长度不符合规范")
	}
	//根据用户名先查询用户是否已存在
	u, err := service.SearchUserByUserName(userName, Id)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RespInternalErr(c)
		return
	}
	if u.UserName != "" {
		util.RespNormalErr(c, 300, "账户已存在")
		return
	}

	//实现密码加盐
	hashPassword := service.SecretPassword(password)
	//插入用户信息有关数据
	err = service.CreateUser(model.User{
		UserName: userName,
		Password: hashPassword,
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
	Id := c.PostForm("Id")
	userName := c.PostForm("UserName")
	password := c.PostForm("Password")
	if userName == "" || password == "" {
		util.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByUserName(userName, Id)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespNormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
		}
		return
	}
	//验证密码是否正确
	jud := service.CompareHashPassword(password, u.Password)
	if !jud {
		util.RespNormalErr(c, 20002, "密码错误")
		return
	}
	util.RespOK(c)
	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
}

func Forget(c *gin.Context) {
	Question := c.PostForm("Question")
	Answer := c.PostForm("Answer")
	if Question == "" || Answer == "" {
		util.RespParamErr(c)
		return
	}
	u, err := service.ForgetPassword(Question, Answer)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespNormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user's question and answer error:%v", err)
			util.RespInternalErr(c)
		}
		return
	}
	if u.Question != Question && u.Answer != Answer {
		util.RespNormalErr(c, 20002, "保密问题回答错误")
		return
	}
	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
}

func Modify(c *gin.Context) {
	//我希望的是登录之后才能修改密码
	userName := c.PostForm("UserName")
	password := c.PostForm("Password") //获取修改的密码
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
