package routes

import (
	"gin-gonic-gorm/configs/app_config"
	"gin-gonic-gorm/controllers/auth_controller"
	"gin-gonic-gorm/controllers/file_controller"
	"gin-gonic-gorm/controllers/user_controller"
	"gin-gonic-gorm/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {

	route := app.Group("/api")

	userRoute := route.Group("user", middleware.AuthMiddleware)
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)

	//AUTH ROUTE
	route.POST("/login", auth_controller.Login)

	userRoute.GET("/", user_controller.GetAllUser)
	userRoute.GET("/paginate", user_controller.GetUserPaginated)
	userRoute.GET("/:id", user_controller.GetUserById)
	userRoute.POST("/add", user_controller.Store)
	userRoute.PATCH("/update/:id", user_controller.UpdateUser)
	// userRoute.PUT("/update/data/:id", user_controller.UpdateUser)
	userRoute.DELETE("/delete/:id", user_controller.DeleteUser)
	userRoute.GET("/hello", user_controller.Hello)

	// ROUTE File

	authRoute := route.Group("file", middleware.AuthMiddleware)

	authRoute.POST("/", file_controller.HandleUploadFile)
	authRoute.POST("/middleware", middleware.UploadFile, file_controller.SendStatus)
	authRoute.DELETE("/:filename", file_controller.HandleRemoveFile)

}
