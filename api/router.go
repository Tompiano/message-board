package api

import "github.com/gin-gonic/gin"

func Entrance() {
	r := gin.Default()
	u := r.Group("/user")
	{
		u.POST("/register", Register)
		u.GET("/login")
		u.PUT("/password")
	}
	m := r.Group("/message")
	{
		m.GET("/write", GetMessage)  //写留言
		m.POST("/send", SendMessage) //发送留言
		m.PUT("/Modify")             //更新留言
		m.DELETE("/delete")          //删除留言
	}
	r.Run()
}
