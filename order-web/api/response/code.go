package R

type ResCode int64

const (
	CodeInvalidParam ResCode = 1000 + iota
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServeBusy

	CodeNeedLogin
	CodeInvalidToken
	CodeCreateToken
	CodeInvalidPermission
	CodeCreateVerificationCode
	CodeInvalidVerificationCode
	CodeInvalidService
)

var codeMsgMap = map[ResCode]string{
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeServeBusy:       "服务繁忙",

	CodeNeedLogin:               "需要登录",
	CodeInvalidToken:            "无效Token",
	CodeCreateToken:             "生成Token失败",
	CodeInvalidPermission:       "无权限操作",
	CodeCreateVerificationCode:  "生成验证码错误",
	CodeInvalidVerificationCode: "验证码错误",
	CodeInvalidService:          "服务连不上了",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServeBusy]
	}
	return msg
}
