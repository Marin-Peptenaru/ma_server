package middlewares

import (
	"books/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		requestToken := ctx.GetHeader("Authorization")[len(BEARER_SCHEMA):]
		verifiedToken, err := service.NewJwtService().ValidateToken(requestToken)

		if !verifiedToken.Valid || err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}