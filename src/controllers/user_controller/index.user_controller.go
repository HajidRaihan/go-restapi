package user_controller

import (
	"fmt"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/requests"
	"gin-gonic-gorm/responses"
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

// func Store(ctx *gin.Context) {
// 	user := new(models.User)
// 	//var user models.User
// 	err := ctx.BindJSON(user)

// 	if err != nil {
// 		ctx.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err = database.DB.Create(user).Error

// 	if err != nil {
// 		ctx.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(200, gin.H{
// 		"data": user,
// 	})
// }

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

// func UpdateUser(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	var input models.User

// 	errInput := ctx.ShouldBindJSON(&input)

// 	if errInput != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": errInput.Error()})
// 		return
// 	}

// 	var user models.User

// 	err := database.DB.First(&user, id).Error

// 	if err != nil {
// 		ctx.JSON(404, gin.H{"error": "User not found"})
// 		return
// 	}

// 	user.Name = input.Name
// 	user.Address = input.Address

// 	if err := database.DB.Save(&user).Error; err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"data": user})
// }

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)
	userReq := new(requests.UserRequest)
	userEmailExist := new(models.User)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errReq.Error()})
		return
	}

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if errDb != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// email exist

	errUserEmailExist := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist)

	if errUserEmailExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server error"})
		return
	}

	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email already exist"})
		return
	}

	user.Name = &userReq.Name
	user.Address = &userReq.Address
	user.Email = &userReq.Email
	fmt.Println(user.Email)
	user.BornDate = &userReq.BornDate

	if err := database.DB.Table("users").Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "pantek"})

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
