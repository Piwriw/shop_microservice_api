package shop_cart

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"shop_api/order-web/api/common"
	R "shop_api/order-web/api/response"
	"shop_api/order-web/convert"
	"shop_api/order-web/global"
	"shop_api/order-web/model/request"
	"shop_api/order-web/proto"
	"strconv"
)

// GetListHandler 获取购物车列表
func GetListHandler(c *gin.Context) {
	userId, _ := c.Get("userId")
	rsp, err := global.OrderSrvClient.CartItemList(context.Background(), &proto.UserInfo{
		Id: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("[List] 查询 【购物车】列表失败")
		convert.HandleGrpcError2Http(err, c)
		return
	}
	ids := make([]int32, 0)
	for _, item := range rsp.Data {
		ids = append(ids, item.GoodsId)
	}
	if len(ids) == 0 {
		R.RespSuccessMsg(c, "没有商品")
		return
	}

	// 请求商品服务 获取商品详细
	goodsRsp, err := global.GoodSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: ids,
	})
	if err != nil {
		zap.S().Errorw("[List] 查询 【商品服务】详情列表失败")
		convert.HandleGrpcError2Http(err, c)
		return
	}
	goodsList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		for _, good := range goodsRsp.Data {
			if good.Id == item.GoodsId {
				tmpMap := map[string]interface{}{}
				tmpMap["id"] = item.Id
				tmpMap["goods_id"] = item.GoodsId
				tmpMap["good_name"] = good.Name
				tmpMap["good_price"] = good.ShopPrice
				tmpMap["num"] = item.Nums
				tmpMap["checked"] = item.Checked
				goodsList = append(goodsList, tmpMap)
			}
		}
	}

	reMap := map[string]interface{}{
		"total": rsp.Total,
		"data":  goodsList,
	}
	R.RespSuccess(c, reMap)
}
func DeleteHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		R.RespError(c, R.CodeInvalidParam)
		return
	}
	userId, _ := c.Get("userId")
	_, err = global.OrderSrvClient.DeleteCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(id),
	})
	if err != nil {
		zap.S().Errorw("删除购物车失败")
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccess(c, nil)
}
func CreateHandler(c *gin.Context) {
	req := new(request.ShopCartItemReq)
	if err := c.ShouldBindJSON(req); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	// 添加商品到购物车，先检查商品是否存在
	_, err := global.GoodSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: req.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("[List]查询失败，查询【商品信息】失败了")
		convert.HandleGrpcError2Http(err, c)
		return
	}
	// 检查库存
	invDetail, err := global.InventoryClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsID: req.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("[List]查询失败，查询【库存信息】失败了")
		convert.HandleGrpcError2Http(err, c)
		return
	}
	if invDetail.Num < req.Nums {
		R.RespErrorWithMsg(c, "库存不足")
		return
	}
	userId, _ := c.Get("userId")
	cartItem, err := global.OrderSrvClient.CreateCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: req.GoodsId,
		Nums:    req.Nums,
	})
	if err != nil {
		zap.S().Errorw("添加到购物车失败了")
		convert.HandleGrpcError2Http(err, c)
		return
	}
	resMap := map[string]int32{}
	resMap["id"] = cartItem.Id
	R.RespSuccess(c, resMap)
}
func UpdateHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		R.RespError(c, R.CodeInvalidParam)
		return
	}
	req := new(request.ShopCartItemUpdateReq)
	if err := c.ShouldBindJSON(req); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	userId, _ := c.Get("userId")

	request := proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(id),
		Nums:    req.Nums,
	}
	if req.Checked != nil {
		request.Checked = *req.Checked
	}
	_, err = global.OrderSrvClient.UpdateCartItem(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("更新购物车失败")
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccess(c, nil)
}
