package routes

import (
	"gin-gonic-gorm/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {

	route := app
	route.GET("/", user_controller.GetAllUser)
}
