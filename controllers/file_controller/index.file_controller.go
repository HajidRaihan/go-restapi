package file_controller

import (
	"fmt"
	"gin-gonic-gorm/constanta"
	"gin-gonic-gorm/utils"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HandleUploadFile(ctx *gin.Context) {
	claimsData := ctx.MustGet("claimsData").(jwt.MapClaims)
	fmt.Println("claimsData => email => ", claimsData["email"])

	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userId => ", userId)

	fileHeader, _ := ctx.FormFile("file")

	if fileHeader == nil {
		ctx.JSON(400, gin.H{"message": "file is required"})
		return
	}

	//validation file by extention
	fileExtension := []string{".jpg", ".png", ".PNG", ".jpeg", ".pdf"}
	isFileValidated := utils.FileValidationByExtension(fileHeader, fileExtension)
	if !isFileValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file not allowed",
		})
		return
	}

	//validation file by content-type
	// fileType := []string{"image/jpg"}
	// isFileValidated := utils.FileValidation(fileHeader, fileType)
	// if !isFileValidated {
	// 	ctx.AbortWithStatusJSON(400, gin.H{
	// 		"message": "file not allowed",
	// 	})
	// 	return
	// }

	extensionFile := filepath.Ext(fileHeader.Filename)

	filename := utils.RandomFileName(extensionFile)

	isSaved := utils.SaveFile(ctx, fileHeader, filename)

	if !isSaved {
		ctx.JSON(500, gin.H{"message": "error uploading file"})
		return
	}

	ctx.JSON(200, gin.H{"message": "file uploaded successfully"})
}

func HandleRemoveFile(ctx *gin.Context) {
	filename := ctx.Param("filename")

	if filename == "" {
		ctx.JSON(400, gin.H{"message": "filename is required"})
	}

	err := utils.RemoveFile(constanta.DIR_FILE + filename)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "error removing file",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "file deleted",
	})
}

func SendStatus(ctx *gin.Context) {
	claimsData := ctx.MustGet("claimsData").(jwt.MapClaims)
	fmt.Println("claimsData => email => ", claimsData["email"])

	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userId => ", userId)

	filename := ctx.MustGet("filename").(string)

	ctx.JSON(200, gin.H{
		"message":   "ok",
		"file_name": filename,
	})
}
