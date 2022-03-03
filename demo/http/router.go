package http

import (
	"demo/http/inner"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {

	innerGroup := router.Group("/inner")
	{
		userGroup := innerGroup.Group("/user")
		{

			userGroup.GET("/create", new(inner.UserController).Create)
			//userGroup.DELETE("/del", new(inner.UserController).Delete)
			//userGroup.PUT("/update", new(inner.UserController).Update)
			//userGroup.GET("/userinfo", new(inner.UserController).Userinfo)
			//userGroup.GET("/list", new(inner.UserController).List)

		}
	}

}
