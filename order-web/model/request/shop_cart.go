package request

type ShopCartItemReq struct {
	GoodsId int32 `json:"goodsId" binding:"required"`
	Nums    int32 `json:"nums" binding:"required,min=1"`
}
type ShopCartItemUpdateReq struct {
	Nums    int32 `json:"nums" binding:"required,min=1"`
	Checked *bool `json:"checked"`
}
