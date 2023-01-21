package main

import (
	"moex/config"

	"moex/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var token string

var (
	db *gorm.DB = config.ConnectDB()
)

func init() {
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file")
	}
}

func main() {
	defer config.DisconnectDB(db)

	router := gin.Default()
	// Вызываем бота
	router.GET("api/event/edit/:ticker/:url", controllers.GetNew)
	router.POST("api/event/save", controllers.SaveNews)
	//	router.GET("/api/chart/events", controllers.GetChart)

	router.Run("localhost:8083")
}
