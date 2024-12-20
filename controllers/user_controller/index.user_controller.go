package user_controller

import (
	"errors"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/requests"
	"gin-gonic-gorm/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUser(ctx *gin.Context) {
	users := new([]models.User)

	err := database.DB.Find(&users).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"data": users,
	})
}

func GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user := new(responses.UserResponse)

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	// err := database.DB.First(&user, id).Error

	// err := database.DB.Table("users").Where("id = ?", id).First(&user).Error

	if errDb != nil {
		ctx.JSON(500, gin.H{"message": "internal server error"})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})
}

func Store(ctx *gin.Context) {
	// userReq := new(requests.UserRequest)
	userReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	userEmailExist := new(models.User)

	errUserEmailExist := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist).Error

	if errUserEmailExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server pskjhkj"})
		return
	}

	if userEmailExist.Email != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email already exist"})
		return
	}

	user := new(models.User)

	user.Name = &userReq.Name
	user.Address = &userReq.Address
	user.Email = &userReq.Email
	user.BornDate = &userReq.BornDate

	if err := database.DB.Table("users").Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "data saved successfully", "data": user})
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)
	userReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	// Cari user yang akan diupdate
	if err := database.DB.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Cek email duplikat
	var existingUser models.User
	err := database.DB.Table("users").Where("email = ? AND id != ?", userReq.Email, id).First(&existingUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Update user data
	user.Name = &userReq.Name
	user.Address = &userReq.Address
	user.Email = &userReq.Email
	user.BornDate = &userReq.BornDate

	if err := database.DB.Table("users").Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user := new(models.User)

	errFind := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if errFind != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	errDb := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&models.User{}).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User deleted",
	})
}

func Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}
