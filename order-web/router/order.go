package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/order-web/api/order"
	"shop_api/order-web/middleware"
)

func InitOrderRouter(router *gin.RouterGroup) {
	OrderRouter := router.Group("orders").Use(middleware.JWTAuth())
	{
		//OrderRouter.GET("", middleware.JWTAuth(), middleware.IsAdminAuth(), order.GetListHandler)
		OrderRouter.GET("", order.GetListHandler)
		OrderRouter.POST("", order.CreateHandler)
		OrderRouter.GET("/:id/", order.GetDetailHandler)
	}
}
