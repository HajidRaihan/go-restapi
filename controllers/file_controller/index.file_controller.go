package file_controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleUploadFile(ctx *gin.Context) {
	fileHeader, _ := ctx.FormFile("file")

	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "file is required"})
		return
	}

	// file, errFile := fileHeader.Open()
	// if errFile != nil {
	// 	panic(errFile)
	// }

	// extensionFile := filepath.Ext(fileHeader.Filename)

	// filename := "file_" + strconv.FormatInt(time.Now().Unix(), 10) + extensionFile

	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%d-%s", time.Now().Unix(), fileHeader.Filename))

	if errUpload != nil {
		ctx.JSON(500, gin.H{"message": "error uploading file", "error": errUpload.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "file uploaded successfully"})
}
