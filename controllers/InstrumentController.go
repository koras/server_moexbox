package controllers

// func GetInstrument(c *gin.Context) models.New {
// 	ticker := c.Param("ticker")
// 	url := c.Param("url")
// 	// SELECT  * FROM `stacks` JOIN `users`  ON `users`.id = `stacks`.users_id ORDER BY `lang`, `nickname`
// 	var chartData models.New
// 	if err := db.Select("date", "`prices`.`price`").Table("prices").Where("prices.instrument_id = ? and", ticker, url).Joins("INNER JOIN instruments ON instruments.instrument_id = prices.instrument_id").Order("date asc").Find(&chartData).Error; err != nil {
// 		fmt.Println(err)
// 	}
// 	return chartData
// }
