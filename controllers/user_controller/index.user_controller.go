package user_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

	var user models.User

	// err := database.DB.First(&user, id).Error

	err := database.DB.Table("users").Where("id = ?", id).First(&user).Error

	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})
}

func AddUser(ctx *gin.Context) {
	user := new(models.User)
	//var user models.User
	err := ctx.BindJSON(user)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.Create(user).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var input models.User

	errInput := ctx.ShouldBindJSON(&input)

	if errInput != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errInput.Error()})
		return
	}

	var user models.User

	err := database.DB.First(&user, id).Error

	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	user.Name = input.Name
	user.Address = input.Address

	if err := database.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})

}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user models.User

	err := database.DB.Find(&user, id).Error

	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user)

	ctx.JSON(200, gin.H{
		"message": "User deleted",
		"data":    user,
	})

}

func Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}
