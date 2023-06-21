package middleware

import (
	"github.com/gin-gonic/gin"
	R "shop_api/user-web/api/response"
	jwt_request "shop_api/user-web/model/request"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*jwt_request.CustomClaims)
		if currentUser.AuthorityId != 2 {
			R.ResponseError(ctx, R.CodeInvalidPermission)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
