package bootstrap

import (
	"gin-gonic-gorm/configs/app_config"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/routes"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	app := gin.Default()
	database.ConnectDatabase()
	routes.InitRoutes(app)

	app.Run(app_config.PORT)
	// app.Run(":8000")
}
