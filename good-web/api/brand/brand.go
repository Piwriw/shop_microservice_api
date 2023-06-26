package brand

import (
	"context"
	"github.com/gin-gonic/gin"
	"shop_api/good-web/api/common"
	R "shop_api/good-web/api/response"
	"shop_api/good-web/convert"
	"shop_api/good-web/global"
	"shop_api/good-web/model/request"
	"shop_api/good-web/proto"
	"strconv"
)

func UpdateHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		R.RespErrorWithMsg(c, "获取id失败")
		return
	}
	req := new(request.BrandReq)
	if err = c.ShouldBindJSON(req); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	_, err = global.GoodSrvClient.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   int32(idInt),
		Name: req.Name,
		Logo: req.Logo,
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccessMsg(c, "更新品牌成功")
}
func DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		R.RespErrorWithMsg(c, "获取id失败")
		return
	}
	_, err = global.GoodSrvClient.DeleteBrand(context.Background(), &proto.BrandRequest{
		Id: int32(idInt),
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespErrorWithMsg(c, "删除品牌成功！")
}
func CreateHandler(c *gin.Context) {
	req := new(request.BrandReq)
	if err := c.ShouldBindJSON(req); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	_, err := global.GoodSrvClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: req.Name,
		Logo: req.Logo,
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccessMsg(c, "创建品牌成功")
}
func GetListHandler(c *gin.Context) {
	pn := c.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := c.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	rsp, err := global.GoodSrvClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       int32(pnInt),
		PagePerNums: int32(pSizeInt),
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccess(c, rsp)
}
