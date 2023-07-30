package router

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.RouterGroup) {
	InitOrderRouter(router)
	InitShopCartRouter(router)
}
