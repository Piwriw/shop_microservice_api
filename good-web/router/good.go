package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/good-web/api/good"
)

func InitGoodRouter(router *gin.RouterGroup) {
	GoodRouter := router.Group("good")
	{
		GoodRouter.GET("/list", good.GetListHandler)
		GoodRouter.POST("/add", good.CreateHandler)
		GoodRouter.GET("/:id", good.GetDetailHandler)
		GoodRouter.DELETE("/:id", good.DeleteHandler)
		GoodRouter.GET("/:id/stocks", good.GetStocksHandler)
		GoodRouter.PATCH("/:id", good.UpdateStatusHandler)
		GoodRouter.PUT("/:id", good.UpdateHandler)

	}
}
