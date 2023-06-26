package response

import "shop_api/good-web/proto"

type CategoryDetailRes struct {
	ID             int32                         `json:"id"`
	Name           string                        `json:"name"`
	Level          int32                         `json:"level"`
	ParentCategory int32                         `json:"parentCategory"`
	IsTab          bool                          `json:"isTab"`
	SubCategorys   []*proto.CategoryInfoResponse `json:"subCategorys"`
}
