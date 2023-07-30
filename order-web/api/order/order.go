package order

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	R "shop_api/good-web/api/response"
	"shop_api/order-web/api/common"
	"shop_api/order-web/convert"
	"shop_api/order-web/global"
	"shop_api/order-web/model/request"
	"shop_api/order-web/model/response"
	"shop_api/order-web/proto"
	"strconv"
)

func GetListHandler(c *gin.Context) {
	userId, _ := c.Get("userId")
	claims, _ := c.Get("claims")
	models := claims.(*request.CustomClaims)
	req := proto.OrderFilterRequest{}
	// 如果用户是管理员，反馈所有订单
	if models.AuthorityId == 1 {
		req.UserId = int32(userId.(uint))
	}
	page, num := common.GetPage(c)
	req.Pages = page
	req.PagePerNums = num
	rsp, err := global.OrderSrvClient.OrderList(context.Background(), &req)
	if err != nil {
		zap.S().Errorw("获取订单列表失败")
		convert.HandleGrpcError2Http(err, c)
		return
	}
	res := make([]response.OrderListItemRsp, 0)
	for _, item := range rsp.Data {
		res = append(res, response.OrderListItemRsp{
			Id:      item.Id,
			UserId:  item.UserId,
			OrderSn: item.OrderSn,
			PayType: item.PayType,
			Status:  item.Status,
			Post:    item.Post,
			Total:   item.Total,
			Address: item.Address,
			Name:    item.Name,
			Mobile:  item.Mobile,
			AddTime: item.AddTime,
		})
	}
	R.RespSuccess(c, res)
}
func CreateHandler(c *gin.Context) {
	req := new(request.CreateOrderReq)
	if err := c.ShouldBindJSON(req); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	userId, _ := c.Get("userId")
	rsp, err := global.OrderSrvClient.CreateOrder(context.Background(), &proto.OrderRequest{
		UserId:  int32(userId.(uint)),
		Address: req.Address,
		Name:    req.Name,
		Mobile:  req.Mobile,
		Post:    req.Mobile,
	})
	if err != nil {
		zap.S().Errorw("新建订单失败")
		convert.HandleGrpcError2Http(err, c)
		return
	}
	resMap := map[string]int32{"id": rsp.Id}
	R.RespSuccess(c, resMap)
}
func GetDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	res, err := global.OrderSrvClient.OrderDetail(context.Background(), &proto.OrderRequest{Id: int32(id)})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccess(c, res)
}
