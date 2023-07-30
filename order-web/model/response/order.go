package response

type OrderListItemRsp struct {
	Id      int32   `json:"id"`
	UserId  int32   `json:"userId"`
	OrderSn string  `json:"orderSn"`
	PayType string  `json:"payType"`
	Status  string  `json:"status"`
	Post    string  `json:"post"`
	Total   float32 `json:"total"`
	Address string  `json:"address"`
	Name    string  `json:"name"`
	Mobile  string  `json:"mobile"`
	AddTime string  `json:"addTime"`
}
