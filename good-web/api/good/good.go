package good

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"shop_api/good-web/api/common"
	R "shop_api/good-web/api/response"
	"shop_api/good-web/convert"
	"shop_api/good-web/global"
	"shop_api/good-web/model/request"
	"shop_api/good-web/proto"
	"strconv"
)

func UpdateHandler(c *gin.Context) {
	req := new(request.GoodsReq)
	if err := c.ShouldBindJSON(req); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		R.FailWithDetailed(c, "参数解析错误", err.Error())
		return
	}
	_, err = global.GoodSrvClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:              int32(idInt),
		Name:            req.Name,
		GoodsSn:         req.GoodSn,
		Stocks:          req.Stocks,
		MarketPrice:     req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodBrief,
		ShipFree:        *req.ShipFree,
		Images:          req.Images,
		DescImages:      req.DescImages,
		GoodsFrontImage: req.FrontImage,
		CategoryId:      req.CategoryId,
		BrandId:         req.Brand,
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccessMsg(c, "更新商品信息成功！")
}

// UpdateStatusHandler 只更新商品状态
func UpdateStatusHandler(c *gin.Context) {
	req := new(request.GoodsStatusReq)
	if err := c.ShouldBindJSON(req); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		R.FailWithDetailed(c, "参数解析错误", err.Error())
		return
	}
	_, err = global.GoodSrvClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:         int32(idInt),
		IsHot:      *req.IsHot,
		IsNew:      *req.IsNew,
		OnSale:     *req.OnSale,
		CategoryId: req.CategoryId,
		BrandId:    req.Brand,
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccess(c, nil)
}
func GetStocksHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		R.FailWithDetailed(c, "参数解析错误", err.Error())
		return
	}
	//todo 获取商品库存
	return
}
func DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		R.FailWithDetailed(c, "参数解析错误", err.Error())
		return
	}
	_, err = global.GoodSrvClient.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{
		Id: int32(idInt),
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccess(c, "")
}
func GetDetailHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		R.FailWithDetailed(c, "参数解析错误", err.Error())
		return
	}
	r, err := global.GoodSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: int32(idInt),
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	rsp := map[string]interface{}{
		"id":          r.Id,
		"name":        r.Name,
		"goods_brief": r.GoodsBrief,
		"desc":        r.GoodsDesc,
		"ship_free":   r.ShipFree,
		"images":      r.Images,
		"desc_images": r.DescImages,
		"front_image": r.GoodsFrontImage,
		"shop_price":  r.ShopPrice,
		"category": map[string]interface{}{
			"id":   r.Category.Id,
			"name": r.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   r.Brand.Id,
			"name": r.Brand.Name,
			"logo": r.Brand.Logo,
		},
		"is_hot":  r.IsHot,
		"is_new":  r.IsNew,
		"on_sale": r.OnSale,
	}
	R.RespSuccess(c, rsp)
}
func CreateHandler(c *gin.Context) {
	goodReq := new(request.GoodsReq)
	if err := c.ShouldBindJSON(goodReq); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	_, err := global.GoodSrvClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Name:            goodReq.Name,
		GoodsSn:         goodReq.GoodSn,
		Stocks:          goodReq.Stocks,
		MarketPrice:     goodReq.MarketPrice,
		ShopPrice:       goodReq.ShopPrice,
		GoodsBrief:      goodReq.GoodBrief,
		ShipFree:        *goodReq.ShipFree,
		Images:          goodReq.Images,
		DescImages:      goodReq.DescImages,
		GoodsFrontImage: goodReq.FrontImage,
		CategoryId:      goodReq.CategoryId,
		BrandId:         goodReq.Brand,
	})
	if err != nil {
		convert.HandleGrpcError2Http(err, c)
		return
	}
	R.RespSuccess(c, "")
}
func GetListHandler(c *gin.Context) {
	request := &proto.GoodsFilterRequest{}
	priceMin := c.DefaultQuery("pmin", "0")
	priceMinInt, err := strconv.Atoi(priceMin)

	priceMax := c.DefaultQuery("pman", "0")
	priceMaxInt, err := strconv.Atoi(priceMax)
	if err != nil {
		R.FailWithDetailed(c, "参数解析错误", err.Error())
		return
	}
	request.PriceMin = int32(priceMinInt)
	request.PriceMax = int32(priceMaxInt)

	isHot := c.DefaultQuery("ih", "0")
	if isHot == "1" {
		request.IsHot = true
	}
	isNew := c.DefaultQuery("in", "0")
	if isNew == "1" {
		request.IsNew = true
	}
	isTab := c.DefaultQuery("it", "0")
	if isTab == "1" {
		request.IsTab = true
	}
	categoryId := c.DefaultQuery("c", "0")
	categoryIdInt, err := strconv.Atoi(categoryId)
	request.TopCategory = int32(categoryIdInt)

	page, pageNum := common.GetPage(c)
	request.Pages = page
	request.PagePerNums = pageNum

	keywords := c.DefaultQuery("keywords", "")
	request.KeyWords = keywords

	brandId := c.DefaultQuery("bid", "0")
	brandIdInt, _ := strconv.Atoi(brandId)
	request.Brand = int32(brandIdInt)

	// 请求商品服务
	res, err := global.GoodSrvClient.GoodsList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("[GoodsList] 查询 商品列表失败")
		convert.HandleGrpcError2Http(err, c)
		return
	}
	reMap := map[string]interface{}{
		"total": res.Total,
	}
	goodsList := make([]interface{}, 0)
	for _, value := range res.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":          value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"category": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"is_hot":  value.IsHot,
			"is_new":  value.IsNew,
			"on_sale": value.OnSale,
		})
	}
	reMap["data"] = goodsList
	R.RespSuccess(c, reMap)
}
