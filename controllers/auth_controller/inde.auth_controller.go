package auth_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/requests"
	"gin-gonic-gorm/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	// ...
)

func Login(ctx *gin.Context) {
	loginReq := new(requests.LoginRequest)
	if errReq := ctx.ShouldBind(&loginReq); errReq != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": errReq.Error()})
		return
	}

	user := new(models.User)
	err := database.DB.Table("users").Where("email = ?", loginReq.Email).Find(&user).Error
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Invalid email or password"})
		return
	}

	if loginReq.Password != "12345" {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Invalid email or password"})
		return
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token, errToken := utils.GenerateToken(&claims)
	if errToken != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": "Failed to generate token"})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "login successfully",
		"token":   token,
	})
}
