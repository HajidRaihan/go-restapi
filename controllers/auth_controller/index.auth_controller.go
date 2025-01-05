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

	passIsValid := utils.VerifyPassword(loginReq.Password, *user.Password)

	if !passIsValid {
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

func Register(ctx *gin.Context) {
	registerReq := new(requests.RegisterRequest)

	errReq := ctx.ShouldBind(&registerReq)

	if errReq != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": errReq.Error()})
		return
	}

	user := new(models.User)

	password := registerReq.Password

	hashedPass, errHash := utils.HashPassword(password)

	if errHash != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": "Failed to hash password"})
		return
	}

	user.Name = &registerReq.Name
	user.Email = &registerReq.Email
	user.Password = &hashedPass

	if err := database.DB.Table("users").Create(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": "Failed to create user"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User created successfully",
	})
}
