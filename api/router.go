package api

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	u := r.Group("/user")
	{
		u.POST("/register", Register)
		u.GET("/login")
		u.PUT("/password")
	}
	r.Run()
}
