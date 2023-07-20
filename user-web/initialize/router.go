package initialize

import (
	"github.com/gin-gonic/gin"
	"shop_api/user-web/api"
	"shop_api/user-web/middleware"
	"shop_api/user-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", api.HealthHandler)
	ApiGroup := Router.Group("/v1")
	ApiGroup.Use(middleware.Cors())
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	return Router
}
