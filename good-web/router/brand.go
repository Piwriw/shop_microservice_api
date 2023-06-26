package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/good-web/api/brand"
)

func InitBrandRouter(router *gin.RouterGroup) {
	BrandRouter := router.Group("brand")
	{
		BrandRouter.GET("/list", brand.GetListHandler)
		BrandRouter.POST("/add", brand.CreateHandler)
		BrandRouter.DELETE("/:id", brand.DeleteHandler)
		BrandRouter.PUT("/:id", brand.UpdateHandler)
	}
}
