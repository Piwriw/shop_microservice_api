package request

type CategoryReq struct {
	Name           string `form:"name" json:"name" binding:"required,min=3,max=20"`
	ParentCategory int32  `form:"parent" json:"parent"`
	Level          int32  `form:"level" json:"level" binding:"required,oneof=1 2 3"`
	IsTab          *bool  `form:"isTab" json:"isTab" binding:"required"`
}
type UpdateCategoryReq struct {
	Name  string `form:"name" json:"name" binding:"required,min=3,max=20"`
	IsTab *bool  `form:"isTab" json:"isTab" binding:"required"`
}
