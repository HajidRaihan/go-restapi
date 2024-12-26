package file_controller

import (
	"gin-gonic-gorm/utils"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleUploadFile(ctx *gin.Context) {
	fileHeader, _ := ctx.FormFile("file")

	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "file is required"})
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
