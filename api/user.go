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
	if userName == "" || password == "" {
		util.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByUserName(userName)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RespInternalErr(c)
		return
	}
	if u.UserName != "" {
		util.NormalErr(c, 300, "账户已存在")
		return
	}
	err = service.CreateUser(model.User{
		UserName: userName,
		Password: password,
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
			util.NormalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	if u.Password == password {
		util.NormalErr(c, 20001, "密码错误")
		return
	}
	c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
}
