package good

import (
	"github.com/gin-gonic/gin"
	R "shop_api/good-web/api/response"
)

func HealthHandler(c *gin.Context) {
	R.RespSuccess(c, "success")
}
