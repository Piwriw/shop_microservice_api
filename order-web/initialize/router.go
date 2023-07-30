package initialize

import (
	"github.com/gin-gonic/gin"
	"shop_api/order-web/api"
	"shop_api/order-web/middleware"
	"shop_api/order-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", api.HealthHandler)
	ApiGroup := Router.Group("/v1")
	ApiGroup.Use(middleware.Cors())
	router.InitRouter(ApiGroup)
	return Router
}
