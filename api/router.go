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
		m.POST("/send", SendMessage) //发表留言
		m.GET("/look", GetMessage)   //查看留言
		m.PUT("/update", Update)     //修改留言
		m.DELETE("/delete", Delete)  //删除留言
		m.POST("/like", Like)        //对留言点赞
	}
	t := r.Group("/comment")
	{
		t.POST("/send", SendComment)       //发表评论
		t.GET("/look", LookComment)        //查看评论
		t.PUT("/modify", ModifyComment)    //修改评论
		t.DELETE("/delete", DeleteComment) //删除评论
	}

	r.Run()
}
