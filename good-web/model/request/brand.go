package request

type BrandReq struct {
	Name string `form:"name" json:"name" binding:"required,min=3,max=10"`
	Logo string `form:"logo" json:"logo" binding:"url"`
}

type CategoryBrandReq struct {
	CategoryId int `form:"categoryId" json:"categoryId" binding:"required"`
	BrandId    int `form:"brandId" json:"brandId" binding:"required"`
}
