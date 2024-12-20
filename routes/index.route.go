package routes

import (
	"gin-gonic-gorm/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {

	route := app

	route.GET("/", user_controller.GetAllUser)
	route.GET("/:id", user_controller.GetUserById)
	route.POST("/add", user_controller.Store)
	route.PATCH("/update/:id", user_controller.UpdateUser)
	// route.PUT("/update/data/:id", user_controller.UpdateUser)
	route.DELETE("/delete/:id", user_controller.DeleteUser)
	route.GET("/hello", user_controller.Hello)

}
