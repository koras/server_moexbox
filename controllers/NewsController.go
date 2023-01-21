package controllers

import (
	"crypto/md5"
	"fmt"
	"log"
	"math/rand"
	"moex/models"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// получаем одну новость
func GetNew(c *gin.Context) {
	ticker := c.Param("ticker")
	url := c.Param("url")
	var chartData models.New
	if err := db.Select("event_id", "source", "type_id", "hash").Table("events").Where("instruments.instrument_id = ? and events.url = ?", ticker, url).Joins("INNER JOIN instruments ON instruments.instrument_id = events.instrument_id").Find(&chartData).Error; err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "user not found"})
	} else {
		c.JSON(200, chartData)
	}

	//	return chartData
}

// сохраняем новость
func SaveNews(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())

	var eventInput models.EventInput
	c.Bind(&eventInput)
	//c.Bind(&event)

	if eventInput.Date != "" && eventInput.TypeID != 0 && eventInput.Source != "" && eventInput.InstrumentID != 0 &&
		eventInput.Title != "" && eventInput.Shorttext != "" && eventInput.Fulltext != "" {
		// Получаем уникальный урл
		slug := slug.Make(eventInput.Title)
		// получаем уникальный хэш
		hash := fmt.Sprintf("%x", md5.Sum([]byte(eventInput.Title+eventInput.Date+string(rand.Intn(50)))))
		// Преобразуем дату  ПРиходит 01/03/2022 записываем 2019-12-29
		re := regexp.MustCompile(`([0-9]{2})/([0-9]{2})/([0-9]{4})`)
		rD := (re.FindAllStringSubmatch(eventInput.Date, -1))
		reversed := fmt.Sprintf("%s-%s-%s", rD[0][3], rD[0][2], rD[0][1])

		EventCreate := models.Events{
			UserId:       100,
			TypeID:       eventInput.TypeID,
			Title:        eventInput.Title,
			Date:         reversed,
			InstrumentID: eventInput.InstrumentID,
			Slug:         slug,
			Hash:         hash,
			Source:       eventInput.Source,
			Shorttext:    eventInput.Shorttext,
			Fulltext:     eventInput.Fulltext,
		}

		db.Create(EventCreate)

		content := &models.EventResult{
			InstrumentID: eventInput.InstrumentID,
			Slug:         slug,
			Hash:         hash,
		}
		c.JSON(201, content)

		//} else {
		//		checkErr(err, "Insert failed")
		//}
		//	}

	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
	//	return chartData
}
