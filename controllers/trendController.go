package controllers

import (
	"github.com/gin-gonic/gin"
)

type Instrument struct {
	Ticker         string  `json:"ticker" gorm:"column:ticker"`
	InstrumentName string  `json:"instrument_name" gorm:"column:instrument_name"`
	Type           string  `json:"type" gorm:"column:type"`
	Mark           string  `json:"mark" gorm:"column:mark"`
	Logo           string  `json:"logo" gorm:"column:logo"`
	Level          string  `gorm:"column:level"`
	Price          float64 `json:"price" gorm:"column:price"`
	Date           string  `json:"max_date" gorm:"column:max_date"`
	PricesPrice    float64 `json:"years_price" gorm:"column:prices_price"`
}

type AllTrend struct {
	PriceYear  []Instrument `json:"year"`
	PriceYear5 []Instrument `json:"year5"`
	PriceMonth []Instrument `json:"month"`
}

// получаем список инструментов для дожборда
func TrendList(c *gin.Context) {

	var instrumentsYears5 []Instrument
	var instrumentsYears []Instrument
	var instrumentsMonth []Instrument
	//var prices []Price

	// Выполнить запрос

	resultYears := db.Table("instruments").
		Select("MAX(prices.date) AS max_date, instruments.type, instruments.ticker,  instruments.instrument_name, instruments.mark, instruments.logo, instruments.level, prices.price AS prices_price, instruments.price").
		Joins("LEFT JOIN (SELECT p1.name, MAX(p1.date) AS max_date FROM prices p1 WHERE p1.date BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 YEAR) AND DATE_SUB(NOW(), INTERVAL 360 DAY) GROUP BY p1.name) AS latest_prices ON latest_prices.name = instruments.ticker").
		Joins("LEFT JOIN prices ON prices.name = latest_prices.name AND prices.date = latest_prices.max_date").
		Where("instruments.published = ?", true).
		Where("instruments.logo != ?", "").
		Group("instruments.ticker, instruments.instrument_name, instruments.ticker, instruments.type,instruments.mark, instruments.logo, instruments.level, prices.price, instruments.price").
		Scan(&instrumentsYears)
	if resultYears.Error != nil {
		panic(resultYears.Error)
	}

	resultYears5 := db.Table("instruments").
		Select("MAX(prices.date) AS max_date, instruments.type, instruments.ticker,  instruments.instrument_name, instruments.mark, instruments.logo, instruments.level, prices.price AS prices_price, instruments.price").
		Joins("LEFT JOIN (SELECT p1.name, MAX(p1.date) AS max_date FROM prices p1 WHERE p1.date BETWEEN DATE_SUB(CURDATE(), INTERVAL 5 YEAR) AND DATE_SUB(NOW(), INTERVAL 1820 DAY) GROUP BY p1.name) AS latest_prices ON latest_prices.name = instruments.ticker").
		Joins("LEFT JOIN prices ON prices.name = latest_prices.name AND prices.date = latest_prices.max_date").
		Where("instruments.published = ?", true).
		Where("instruments.logo != ?", "").
		Group("instruments.ticker, instruments.instrument_name, instruments.ticker, instruments.type,instruments.mark, instruments.logo, instruments.level, prices.price, instruments.price").
		Scan(&instrumentsYears5)
	if resultYears5.Error != nil {
		panic(resultYears5.Error)
	}

	resultMonth := db.Table("instruments").
		Select("MAX(prices.date) AS max_date, instruments.type, instruments.ticker,  instruments.instrument_name, instruments.mark, instruments.logo, instruments.level, prices.price AS prices_price, instruments.price").
		Joins("LEFT JOIN (SELECT p1.name, MAX(p1.date) AS max_date FROM prices p1 WHERE p1.date BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH) AND DATE_SUB(NOW(), INTERVAL 25 DAY) GROUP BY p1.name) AS latest_prices ON latest_prices.name = instruments.ticker").
		Joins("LEFT JOIN prices ON prices.name = latest_prices.name AND prices.date = latest_prices.max_date").
		Where("instruments.published = ?", true).
		Where("instruments.logo != ?", "").
		Group("instruments.ticker, instruments.instrument_name, instruments.ticker, instruments.type,instruments.mark, instruments.logo, instruments.level, prices.price, instruments.price").
		Scan(&instrumentsMonth)
	if resultMonth.Error != nil {
		panic(resultMonth.Error)
	}

	allTrend := AllTrend{
		PriceYear:  instrumentsYears,
		PriceYear5: instrumentsYears5,
		PriceMonth: instrumentsMonth,
	}
	// Обработать результаты
	//for _, instrument := range instruments {
	//	instrument.LatestPriceDate = time.Time{} // обнулить поле, если дата пустая
	//	fmt.Println(instrument.LatestPriceDate, instrument.InstrumentName, instrument.Mark, instrument.Logo, instrument.Level, instrument.PricesPrice, instrument.Price)
	//	}
	//typeId := c.Query("typeId")
	//	level := c.Query("level")

	c.JSON(200, allTrend)
}
