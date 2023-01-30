package main

import (
	"moex/config"

	"moex/controllers"

	"github.com/gin-contrib/cors"
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

	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost:8080"}
	//config.AddAllowMethods("OPTIONS")

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "POST", "GET", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// config.AllowCredentials = true
	// config.AddAllowHeaders("authorization")
	// router.Use(cors.New(config))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,

		//	AllowAllOrigins: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 1,
		//		MaxAge:           86400,
	}))

	// Получаем все новости для просмотра
	router.GET("api/events/:ticker", controllers.GetNews)

	//Получаем новость для просмотра
	router.GET("api/event/:ticker/:slug", controllers.GetNew)

	//Получаем элементы для создания новой новости
	router.GET("api/event/create/:ticker", controllers.CreateNew)

	// получаем новость для изменения, схранение в зависимости от публикации, новая или старая
	router.GET("api/event/change/:hash", controllers.GetNewsHash)

	// записываем новость, новую или предлагаем для редактирования
	router.POST("api/event/save", controllers.SaveNews)

	//router.Use(cors.New(config))
	//	router.Use(cors.Default())

	router.Run("localhost:8083")
	//router.Run("moexbox.ru:8080")
}

/**

https://internet-lab.ru/port-forwarding


10.10.30.15 TCP 8080 на адрес 195.44.22.33 TCP 8081:
netsh interface portproxy add v4tov4 listenaddress=87.242.105.150 listenport=8080 connectaddress=1270.0.0.1 connectport=8083
netsh interface portproxy add v4tov4 listenaddress=moexbox.ru listenport=8080 connectaddress=1270.0.0.1 connectport=8083

Удалить переадресацию
netsh interface portproxy delete v4tov4 listenaddress=10.10.30.15 listenport=8080

Очистить все правила
netsh interface portproxy reset

netsh interface portproxy show all
*/
