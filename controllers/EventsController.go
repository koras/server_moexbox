package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Events struct {
	EventID          int    `gorm:"column:event_id" json:"eventId"`
	TypeID           int    `gorm:"column:type_id" json:"typeId"`
	Hash             string `gorm:"column:hash" json:"hash"`
	Source           string `gorm:"column:source" json:"source"`
	Slug             string `gorm:"column:slug" json:"slug"`
	Date             string `gorm:"column:date" json:"date"`
	Title            string `gorm:"column:title" json:"title"`
	InstrumentId     string `gorm:"column:instrument_id" json:"instrument_id"`
	InstrumentName   string `gorm:"column:instrument_name" json:"instrument_name"`
	InstrumentLogo   string `gorm:"column:logo" json:"logo"`
	InstrumentTicker string `gorm:"column:ticker" json:"instrument_ticker"`
	Shorttext        string `gorm:"column:shorttext" json:"shorttext"`
	//	Fulltext       string `json:"fulltext"`
}

// получаем новости по тикеру
func GetEvents(c *gin.Context) {

	var events []Events

	db := db.Table("events").Select(
		"event_id",
		"type_id",
		"hash",
		"source",
		"date",
		"slug",
		"title",
		"instruments.instrument_name",
		"instruments.instrument_id",
		"instruments.ticker",
		"instruments.logo",
		"shorttext")
	db = db.Joins("JOIN instruments on events.instrument_id = instruments.instrument_id").Where("events.published = 0 ")

	if err := db.Find(&events).Error; err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "news not found"})
	} else {
		c.JSON(200, events)
	}
}
