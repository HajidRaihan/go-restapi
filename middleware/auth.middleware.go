package middleware

import (
	"gin-gonic-gorm/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	if !strings.Contains(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "invalid token"})
		return
	}

	token := strings.Replace(bearerToken, "Bearer ", "", 1)

	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized sdssd"})
		return
	}

	claimsData, err := utils.DecodeToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	ctx.Set("claimsData", claimsData)
	ctx.Set("user_id", claimsData["id"])
	ctx.Set("user_name", claimsData["name"])
	ctx.Set("user_email", claimsData["email"])

	// if token != "123" {
	// 	ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
	// }
	ctx.Next()
}

func TokenMiddleware(ctx *gin.Context) {

}
