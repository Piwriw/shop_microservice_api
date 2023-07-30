package router

import (
	"github.com/gin-gonic/gin"
	"shop_api/order-web/api/shop_cart"
	"shop_api/order-web/middleware"
)

func InitShopCartRouter(router *gin.RouterGroup) {
	ShopCartRouter := router.Group("shopcarts").Use(middleware.JWTAuth())
	{
		ShopCartRouter.GET("", shop_cart.GetListHandler)
		ShopCartRouter.DELETE("/:id/", shop_cart.DeleteHandler)
		ShopCartRouter.POST("", shop_cart.CreateHandler)
		ShopCartRouter.PATCH("/:id/", shop_cart.UpdateHandler)
	}
}
