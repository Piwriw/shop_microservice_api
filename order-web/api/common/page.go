package common

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) (page int32, pageNum int32) {
	pageNumS := c.DefaultQuery("pnum", "0")
	pageNumInt, _ := strconv.Atoi(pageNumS)
	pageS := c.DefaultQuery("p", "0")
	pageInt, _ := strconv.Atoi(pageS)
	return int32(pageInt), int32(pageNumInt)
}
