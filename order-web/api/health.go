package api

import (
	"github.com/gin-gonic/gin"
	R "shop_api/order-web/api/response"
)

func HealthHandler(c *gin.Context) {
	R.RespSuccess(c, "success")
}
