package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/user-web/api"
	"shop_api/user-web/middleware"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user")
	{
		UserRouter.GET("/list", middleware.JWTAuth(), middleware.IsAdminAuth(), api.GetUserListHandler)
		UserRouter.POST("/pwd_login", api.PassWordLoginHandler)
		UserRouter.POST("/add", api.RegisterUserHandler)
	}
}
