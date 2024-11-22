package user_controller

import "github.com/gin-gonic/gin"

func GetAllUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Hello World!"})
}
