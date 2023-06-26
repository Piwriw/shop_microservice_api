package category

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"shop_api/good-web/api/common"
	R "shop_api/good-web/api/response"
	"shop_api/good-web/convert"
	"shop_api/good-web/global"
	"shop_api/good-web/model/request"
	"shop_api/good-web/model/response"
	"shop_api/good-web/proto"
	"strconv"
)

func UpdateHandler(c *gin.Context) {
	req := new(request.UpdateCategoryReq)
	if err := c.ShouldBindJSON(req); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	_, err := global.GoodSrvClient.UpdateCategory(context.Background(), &proto.CategoryInfoRequest{
		Name:  req.Name,
		IsTab: *req.IsTab,
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccessMsg(c, "更新分类成功！")
}
func DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		R.RespErrorWithMsg(c, "错误的获取参数")
		return
	}
	_, err = global.GoodSrvClient.DeleteCategory(context.Background(), &proto.DeleteCategoryRequest{
		Id: int32(idInt),
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccessMsg(c, "删除分类成功！")
}
func GetDetailHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		R.RespErrorWithMsg(c, "错误的获取参数")
		return
	}
	res, err := global.GoodSrvClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: int32(idInt),
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccess(c, &response.CategoryDetailRes{
		ID:             res.Info.Id,
		Name:           res.Info.Name,
		Level:          res.Info.Level,
		ParentCategory: res.Info.ParentCategory,
		IsTab:          res.Info.IsTab,
		SubCategorys:   res.SubCategorys,
	})
}
func CreateHandler(c *gin.Context) {
	req := new(request.CategoryReq)
	if err := c.ShouldBindJSON(req); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	_, err := global.GoodSrvClient.CreateCategory(context.Background(), &proto.CategoryInfoRequest{
		Name:           req.Name,
		ParentCategory: req.ParentCategory,
		Level:          req.Level,
		IsTab:          *req.IsTab,
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccessMsg(c, "创建分类成功！")
}
func GetListHandler(c *gin.Context) {

	res, err := global.GoodSrvClient.GetAllCategorysList(context.Background(), &empty.Empty{})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	data := make([]interface{}, 0)
	err = json.Unmarshal([]byte(res.JsonData), &data)
	if err != nil {
		zap.S().Errorw("[List] 查询 【分类列表】失败： ", err.Error())
	}
	R.RespSuccess(c, res)
}
