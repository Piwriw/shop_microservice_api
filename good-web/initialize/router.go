package initialize

import (
	"github.com/gin-gonic/gin"
	"shop_api/good-web/api/good"
	"shop_api/good-web/middleware"
	"shop_api/good-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", good.HealthHandler)
	ApiGroup := Router.Group("/v1")
	ApiGroup.Use(middleware.Cors())
	router.InitRouter(ApiGroup)
	return Router
}
