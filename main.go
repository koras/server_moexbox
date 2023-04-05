package main

import (
	"moex/config"
	"moex/parser"

	"moex/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var token string

var (
	db *gorm.DB = config.ConnectDB()
)

func init() {

}

func main() {
	defer config.DisconnectDB(db)

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
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
	router.GET("apidata/events/:ticker", controllers.GetNews)

	// события для модерации
	router.GET("apidata/inspected", controllers.GetEvents)

	//Получаем новость для просмотра
	router.GET("apidata/event/get/:ticker/:slug", controllers.GetNew)

	//Получаем элементы для создания новой новости
	router.GET("apidata/event/create/:ticker", controllers.CreateNew)

	// получаем новость для изменения, схранение в зависимости от публикации, новая или старая
	router.GET("apidata/event/change/:hash", controllers.GetNewsHash)
	// получаем новость для изменения, схранение в зависимости от публикации, новая или старая
	router.GET("apidata/event/inspect/:hash", controllers.GetNewsInspectHash)

	// записываем новость, новую или предлагаем для редактирования
	router.POST("apidata/event/save", controllers.SaveNews)

	// записываем новость, новую или предлагаем для редактирования
	router.GET("apidata/instruments/list", controllers.InstrumentsList)
	// получаем информацию по инструменту
	router.GET("apidata/instrument/get/:InstrumentId", controllers.InstrumentGet)
	// получение инструмента
	router.GET("apidata/data/:ticker", controllers.InstrumentTickerPrice)

	// обновление инструмента
	router.POST("apidata/instrument/update", controllers.InstrumentUpdate)

	router.GET("apidata/parser/history/list", parser.GetPriceMoexHistory)
	router.GET("apidata/parser/store/list", parser.GetPriceMoexOnline)

	// sitemap
	router.GET("apidata/sitemap/create", controllers.CreateSitemaps)

	router.Run("localhost:8093")
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
