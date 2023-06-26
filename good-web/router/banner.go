package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/good-web/api/banner"
)

func InitBannerRouter(router *gin.RouterGroup) {
	BannerRouter := router.Group("banner")
	{
		BannerRouter.GET("/list", banner.GetListHandler)
		BannerRouter.POST("/add", banner.CreateHandler)
		BannerRouter.DELETE("/:id", banner.DeleteHandler)
		BannerRouter.PUT("/:id", banner.UpdateHandler)
	}
}
