package banner

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
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
	req := new(request.BannerReq)
	if err := c.ShouldBindJSON(req); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	_, err = global.GoodSrvClient.UpdateBanner(context.Background(), &proto.BannerRequest{
		Id:    int32(idInt),
		Index: int32(req.Index),
		Image: req.Image,
		Url:   req.Url,
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccessMsg(c, "更新Banner成功!")
}
func DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		R.RespErrorWithMsg(c, "获取id失败")
		return
	}
	_, err = global.GoodSrvClient.DeleteBanner(context.Background(), &proto.BannerRequest{
		Id: int32(idInt),
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccessMsg(c, "删除Banner成功！")
}
func CreateHandler(c *gin.Context) {
	req := new(request.BannerReq)
	if err := c.ShouldBindJSON(req); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	_, err := global.GoodSrvClient.CreateBanner(context.Background(), &proto.BannerRequest{
		Index: int32(req.Index),
		Image: req.Image,
		Url:   req.Url,
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccessMsg(c, "创建Banner成功!")
}
func GetListHandler(c *gin.Context) {
	res, err := global.GoodSrvClient.BannerList(context.Background(), &empty.Empty{})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccess(c, res)
}
