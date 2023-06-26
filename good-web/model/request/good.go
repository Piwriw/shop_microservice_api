package request

type GoodsReq struct {
	Name        string   `form:"name" json:"name" binding:"required,min=2,max=100"`
	GoodSn      string   `form:"goodSn" json:"goodSn" binding:"required,min=2,lt=20"`
	Stocks      int32    `form:"stocks" json:"stocks" binding:"required,min=1"`
	CategoryId  int32    `form:"category" json:"category" binding:"required"`
	MarketPrice float32  `form:"marketPrice" json:"marketPrice" binding:"required,min=0"`
	ShopPrice   float32  `form:"shopPrice" json:"shopPrice" binding:"required,min=0"`
	GoodBrief   string   `form:"goodBrief" json:"goodBrief" binding:"required,min=3"`
	Images      []string `form:"images" json:"images" binding:"required,min=1"`
	DescImages  []string `form:"descImages" json:"descImages" binding:"required,min=1"`
	ShipFree    *bool    `form:"shipFree" json:"shipFree" binding:"required"`
	FrontImage  string   `form:"frontImage" json:"frontImage" binding:"required,url"`
	Brand       int32    `form:"brand" json:"brand" binding:"required"`
}

type GoodsStatusReq struct {
	IsNew      *bool `form:"isNew" json:"isNew" binding:"required"`
	IsHot      *bool `form:"isHot" json:"isHot" binding:"required"`
	OnSale     *bool `form:"onSale" json:"onSale" binding:"required"`
	CategoryId int32 `form:"category" json:"category" binding:"required"`
	Brand      int32 `form:"brand" json:"brand" binding:"required"`
}
