package controllers

import (
	"fmt"
	"moex/models"

	"github.com/gin-gonic/gin"
)

// получаем одну новость
func GetInstrument(c *gin.Context) {
	ticker := c.Param("ticker")

	//var instrument models.Instrument
	status, instrument := GetInstrumentInstanse(ticker)
	//.Select("event_id", "source", "type_id", "hash")
	if status {
		c.JSON(200, instrument)
	} else {
		c.JSON(404, gin.H{"error": "Instrument not found"})
	}
}

// получаем одну новость
func GetInstrumentInstanse(instrumentId string) (bool, models.Instrument) {

	var instrument models.Instrument
	//.Select("event_id", "source", "type_id", "hash")
	if err := db.Table("instruments").Where("instruments.instrument_id = ?", instrumentId).Find(&instrument).Error; err != nil {
		fmt.Println(err)
		return false, instrument
	} else {
		return true, instrument
	}
}

// получаем инструмент по тикеру
func GetInstrumentNameInstanse(ticker string) (bool, models.Instrument) {

	var instrument models.Instrument
	if err := db.Table("instruments").Where("instruments.ticker = ?", ticker).Find(&instrument).Error; err != nil {
		fmt.Println(err)
		return false, instrument
	} else {
		return true, instrument
	}
}

// получаем инструмент по  ID
func GetInstrumentIdInstanse(id string) (bool, models.Instrument) {

	var instrument models.Instrument
	if err := db.Table("instruments").Where("instruments.instrument_id = ?", id).Find(&instrument).Error; err != nil {
		fmt.Println(err)
		return false, instrument
	} else {
		return true, instrument
	}
}
