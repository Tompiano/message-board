package api

import "github.com/gin-gonic/gin"

func Entrance() {
	r := gin.Default()
	r.Use(HandlerFunc())
	u := r.Group("/user")
	{
		u.POST("/register", Register) //注册
		u.GET("/login", Login)        //登录
		u.GET("/forget", Forget)      //忘记密码可通过保密问题找回
		u.PUT("/modify", Modify)      //修改密码
	}
	m := r.Group("/message")
	{
		m.POST("/send", SendMessage) //发表留言或评论
		m.GET("/look", GetMessage)   //查看留言或评论
		m.PUT("/update", Update)     //修改留言或评论
		m.DELETE("/delete", Delete)  //删除留言或评论
		m.POST("/like", Like)        //对留言点赞
	}

	r.Run()
}
