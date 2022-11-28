package api

import "github.com/gin-gonic/gin"

func Entrance() {
	r := gin.Default()
	u := r.Group("/user")
	{
		u.POST("/register", Register) //注册
		u.GET("/login", Login)        //登录
		u.GET("/forget", Forget)      //忘记密码可通过保密问题找回
		u.PUT("/modify", Modify)      //修改密码
	}
	m := r.Group("/message")
	{
		m.POST("/send", SendMessage) //发表留言
		m.GET("/write", GetMessage)  //查看留言
		m.PUT("update", Update)      //更新留言
		m.DELETE("/delete")          //删除留言
	}
	r.Run()
}
