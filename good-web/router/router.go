package router

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.RouterGroup) {
	InitCategoryRouter(router)
	InitGoodRouter(router)
	InitBannerRouter(router)
	InitBrandRouter(router)
}
