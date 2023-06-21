package response

type JwtResponse struct {
	UserId    int32  `json:"userId"`
	NickName  string `json:"nickName"`
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expiredAt"`
}
