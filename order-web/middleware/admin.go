package middleware

import (
	"github.com/gin-gonic/gin"
	R "shop_api/order-web/api/response"
	jwt_request "shop_api/order-web/model/request"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*jwt_request.CustomClaims)
		if currentUser.AuthorityId != 2 {
			R.RespError(ctx, R.CodeInvalidPermission)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
