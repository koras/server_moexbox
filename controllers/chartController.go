package controllers

type Charts struct {
	Label string
	Price string
}

var chartData []Charts

//var db *gorm.DB = config.ConnectDB()

// func GetChart(c *gin.Context) {

// 	ticker := c.Query("ticker")
// 	charts := Chart(ticker)
// 	fmt.Println(ticker)

// 	//	const YYYYMMDD = "2016-11-02"
// 	//	loc, _ := time.LoadLocation("Europe/Berlin")
// 	//var chartsInfo chartData
// 	for i := 0; i < len(charts); i++ {
// 		var tmp Charts

// 		//fmt.Println(time.Date(charts[i].Date))

// 		tmp.Label = charts[i].Date
// 		tmp.Price = charts[i].Price
// 		//("Label":charts[i].Date, charts[i].Price)
// 		// устанавливаем данные
// 		chartData = append(chartData, tmp)

// 	}
// 	c.JSON(http.StatusOK, chartData)
// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"name": "gggg",
// 	// 	"age":  6666,
// 	// })

// 	//c.IndentedJSON(http.StatusOK, Chart)
// }

// func Chart(ticker string) []models.ChartData {
// 	// SELECT  * FROM `stacks` JOIN `users`  ON `users`.id = `stacks`.users_id ORDER BY `lang`, `nickname`
// 	var chartData []models.ChartData
// 	if err := db.Select("date", "`prices`.`price`").Table("prices").Where("prices.instrument_id = ?", ticker).Joins("INNER JOIN instruments ON instruments.instrument_id = prices.instrument_id").Order("date asc").Find(&chartData).Error; err != nil {
// 		fmt.Println(err)
// 	}
// 	return chartData
// }
