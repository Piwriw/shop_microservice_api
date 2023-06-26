package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/good-web/api/category"
)

func InitCategoryRouter(router *gin.RouterGroup) {
	CategoryRouter := router.Group("category")
	{
		CategoryRouter.GET("/list", category.GetListHandler)
		CategoryRouter.POST("/add", category.CreateHandler)
		CategoryRouter.GET("/:id", category.GetDetailHandler)
		CategoryRouter.DELETE("/:id", category.DeleteHandler)
		CategoryRouter.PUT("/:id", category.UpdateHandler)

	}
}
