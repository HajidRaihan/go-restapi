package bootstrap

import (
	"gin-gonic-gorm/configs"
	"gin-gonic-gorm/configs/app_config"
	"gin-gonic-gorm/configs/cors_config"
	"gin-gonic-gorm/configs/log_config"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	// Load .env file
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	configs.InitConfig()

	// Connect to database
	database.ConnectDatabase()

	log_config.DefaultLogging("logs/file/app.log")

	// init gin engine
	app := gin.Default()

	//CORS
	app.Use(cors_config.CorsConfigContrib())

	// init routes
	routes.InitRoutes(app)

	// Run the server
	app.Run(app_config.PORT)
	// app.Run(":8080")
}
