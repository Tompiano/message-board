package api

import "github.com/gin-gonic/gin"

func Entrance() {
	r := gin.Default()
	r.Use(HandlerFunc())
	u := r.Group("/User")
	{

		u.POST("/user", Register) //注册
		u.GET("/user", Login)     //登录
		u.GET("/user", Forget)    //忘记密码可通过保密问题找回
		u.PUT("/user", Modify)    //修改密码
	}
	m := r.Group("Message")
	{
		m.POST("/message", GetMessage) //查看留言
		m.PUT("/message", Update)      //修改留言
		m.DELETE("/message", Delete)   //删除留言
		m.POST("/message", Like)       //对留言点赞
	}
	t := r.Group("Comment")
	{
		t.POST("/comment", SendComment)     //发表评论
		t.GET("/comment", LookComment)      //查看评论
		t.PUT("/comment", ModifyComment)    //修改评论
		t.DELETE("/comment", DeleteComment) //删除评论
	}

	r.Run()
}
