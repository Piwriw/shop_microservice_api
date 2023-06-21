package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop_api/user-web/api/common"
	"shop_api/user-web/api/request"
	"shop_api/user-web/api/response"
	"shop_api/user-web/convert"
	"shop_api/user-web/global"
	"shop_api/user-web/middleware"
	jwt_request "shop_api/user-web/model/request"
	"shop_api/user-web/model/response"
	"shop_api/user-web/proto"
	"strconv"
	"time"
)

func RegisterUserHandler(c *gin.Context) {
	r := new(request.RegisterUserForm)
	if err := c.ShouldBindJSON(&r); err != nil {
		common.HandleValidatorError(c, err)
		return
	}

	_, grpcErr := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: r.NickName,
		PassWord: r.Password,
		Mobile:   r.Mobile,
	})
	if grpcErr != nil {
		//zap.S().Errorf("[Register] 【新建用户失败】：%s", err.Error())
		convert.HandleGrpcError2Http(grpcErr, c)
		return
	}
	R.Success(c)
}

func PassWordLoginHandler(c *gin.Context) {
	passWordLoginForm := request.PassWordLoginForm{}
	if err := c.ShouldBindJSON(&passWordLoginForm); err != nil {
		common.HandleValidatorError(c, err)
		return
	}
	if !store.Verify(passWordLoginForm.CaptchchaID, passWordLoginForm.Captchcha, true) {
		R.ResponseError(c, R.CodeInvalidVerificationCode)
		return
	}

	rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passWordLoginForm.Mobile,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				R.ResponseError(c, R.CodeUserNotExist)
			default:
				R.ResponseError(c, R.CodeServeBusy)
			}
			return
		}
	} else {
		if passRsp, passErr := global.UserSrvClient.CheckPassWord(context.Background(), &proto.CheckPasswordInfo{
			Password:          passWordLoginForm.Password,
			EncryptedPassword: rsp.Password,
		}); passErr != nil {
			R.ResponseError(c, R.CodeInvalidPassword)
			return
		} else {
			if passRsp.IsSuccess {
				//生成Token
				j := middleware.NewJWT()
				claims := jwt_request.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.Nickname,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),
						ExpiresAt: time.Now().Unix() + global.AppConf.Jwt.ExpiresAt,
						Issuer:    "piwriw",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					R.ResponseError(c, R.CodeCreateToken)
				}
				R.SuccessWithDetailed(c, "获取token success", &response.JwtResponse{
					UserId:    rsp.Id,
					NickName:  rsp.Nickname,
					Token:     token,
					ExpiredAt: (time.Now().Unix() + global.AppConf.Jwt.ExpiresAt) * 1000,
				})
				return
			}
			R.ResponseError(c, R.CodeInvalidPassword)
		}
	}
}

func GetUserListHandler(c *gin.Context) {

	claims, _ := c.Get("claims")
	currentUser := claims.(*jwt_request.CustomClaims)
	zap.S().Infof("访问用户：%d", currentUser.ID)
	pn := c.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := c.DefaultQuery("psize", "0")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 【用户列表失败】")
		convert.HandleGrpcError2Http(err, c)
		return
	}
	result := make([]response.UserResponse, 0)
	for _, value := range rsp.Data {
		result = append(result, response.UserResponse{
			Id:       value.Id,
			NickName: value.Nickname,
			Birthday: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		})
	}
	R.SuccessWithDetailed(c, "获取用户列表成功", result)
}
