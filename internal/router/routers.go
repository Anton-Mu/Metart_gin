package router

import (
	"Metart/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	// v1版本api接口
	v1 := r.Group("api/v1")
	{
		v1.POST("/users", controller.CreateUser)
		v1.GET("/users", controller.GetUserList)
	}

	return r
}
