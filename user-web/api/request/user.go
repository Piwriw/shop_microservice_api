package request

type PassWordLoginForm struct {
	Mobile      string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Password    string `json:"password" form:"password" binding:"required,min=3,max=10"`
	Captchcha   string `json:"captchcha" form:"captchcha" binding:"required,min=5,max=5"`
	CaptchchaID string `json:"captchchaID" form:"captchchaID" binding:"required"`
}
type RegisterUserForm struct {
	Mobile   string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Password string `json:"password" form:"password" binding:"required,min=3,max=10"`
	NickName string `json:"nickName" form:"nickName" binding:"required,min=3,max=10"`
}
