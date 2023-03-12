package controllers

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"moex/models"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// получаем одну новость
func CreateNew(c *gin.Context) {
	ticker := c.Param("ticker")
	fmt.Println(ticker)
	status, instrument := GetInstrumentNameInstanse(ticker)
	if !status {

		fmt.Println(instrument)
		c.JSON(404, gin.H{"error": "instrument not found"})
		return
	}

	status, datePrices := InstrumentOnlyDateFromPrices(ticker)
	if !status {

		fmt.Println(instrument)
		c.JSON(404, gin.H{"error": "date not found from prices"})
		return
	}

	fmt.Println("есть новость")
	EventCreate := models.New{}
	EventInstrument := models.EventInstrument{
		Data: EventCreate,
		Instrument: models.Instrument{
			InstrumentID:   instrument.InstrumentID,
			InstrumentName: instrument.InstrumentName,
			Site:           instrument.Site,
			Logo:           instrument.Logo,
			Ticker:         instrument.Ticker,
		},
		PriceDate: datePrices,
	}

	fmt.Println(EventInstrument)
	c.JSON(200, EventInstrument)
}

// получаем одну новость
func GetNew(c *gin.Context) {
	ticker := c.Param("ticker")
	slug := c.Param("slug")
	var chartData models.Events

	fmt.Println(ticker, slug)
	status, instrument := GetInstrumentNameInstanse(ticker)

	if !status {

		fmt.Println(instrument)
		c.JSON(404, gin.H{"error": "instrument not found"})
		return
	}

	status, datePrices := InstrumentOnlyDateFromPrices(ticker)
	if !status {

		fmt.Println(instrument)
		c.JSON(404, gin.H{"error": "date not found from prices"})
		return
	}

	if err := db.Table("events").Where("instrument_id = ? and slug = ?", instrument.InstrumentID, slug).Limit(1).Find(&chartData).Error; err != nil {
		fmt.Println(err)

		EventInstrument := models.EventInstrument{
			// временно для записи
			Data: models.New{},
			Instrument: models.Instrument{
				InstrumentID:   instrument.InstrumentID,
				InstrumentName: instrument.InstrumentName,
				Site:           instrument.Site,
				Logo:           instrument.Logo,
				Ticker:         instrument.Ticker,
			},
			PriceDate: datePrices,
		}

		fmt.Println("нет новостей")
		fmt.Println(EventInstrument)
		c.JSON(200, EventInstrument)
		//	c.JSON(404, gin.H{"error": "news not found"})
	} else {

		fmt.Println("есть новость")
		EventCreate := models.New{
			TypeID:    chartData.TypeID,
			Title:     chartData.Title,
			Date:      revertDateFromBase(chartData.Date),
			Slug:      chartData.Slug,
			Hash:      chartData.Hash,
			EventID:   chartData.EventID,
			Source:    chartData.Source,
			Shorttext: chartData.Shorttext,
			Fulltext:  chartData.Fulltext,
		}

		EventInstrument := models.EventInstrument{
			Data: EventCreate,
			Instrument: models.Instrument{
				InstrumentID:   instrument.InstrumentID,
				InstrumentName: instrument.InstrumentName,
				Site:           instrument.Site,
				Type:           instrument.Type,
				Logo:           instrument.Logo,
				Ticker:         instrument.Ticker,
			},
		}

		//fmt.Println(EventInstrument)
		c.JSON(200, EventInstrument)
	}
}

// получаем одну новость по хэшу для редактирования
func GetNewsHash(c *gin.Context) {
	hash := c.Param("hash")

	var chartData models.Events

	if err := db.Table("events").Where("hash = ? ", hash).Limit(1).Find(&chartData).Error; err != nil {
		fmt.Println(err)

		fmt.Println("нет новостей")
		//	fmt.Println(EventInstrument)
		//	c.JSON(200, EventInstrument)
		c.JSON(404, gin.H{"error": "нет новостей"})
	} else {

		error, instrument := GetInstrumentIdInstanse(chartData.InstrumentID)
		if !error {

			fmt.Println(instrument)
			c.JSON(404, gin.H{"error": "instrument not found"})
			return
		}
		fmt.Println("есть новость")
		EventCreate := models.New{
			TypeID:    chartData.TypeID,
			Title:     chartData.Title,
			Date:      revertDateFromBase(chartData.Date),
			Slug:      chartData.Slug,
			Hash:      chartData.Hash,
			EventID:   chartData.EventID,
			Source:    chartData.Source,
			Shorttext: chartData.Shorttext,
			Fulltext:  chartData.Fulltext,
		}

		EventInstrument := models.EventInstrument{
			Data: EventCreate,
			Instrument: models.Instrument{
				InstrumentID:   instrument.InstrumentID,
				InstrumentName: instrument.InstrumentName,
				Site:           instrument.Site,
				Logo:           instrument.Logo,
				Type:           instrument.Type,
				Ticker:         instrument.Ticker,
			},
		}

		fmt.Println(EventInstrument)
		c.JSON(200, EventInstrument)
	}
}

// получаем новости по тикеру
func GetNews(c *gin.Context) {
	ticker := c.Param("ticker")

	status, instrument := GetInstrumentNameInstanse(ticker)

	if !status {
		c.JSON(404, gin.H{"error": "instrument not found"})
		return
	}

	var news []models.News
	if err := db.Table("events").Where("published = 1 and instrument_id = ? ", instrument.InstrumentID).Find(&news).Error; err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "news not found"})
	} else {

		eventInstrument := models.EventsInstrument{
			Data: news,
			Instrument: models.Instrument{
				InstrumentID:   instrument.InstrumentID,
				InstrumentName: instrument.InstrumentName,
				Site:           instrument.Site,
				Logo:           instrument.Logo,
				Type:           instrument.Type,
				Ticker:         instrument.Ticker,
			},
		}

		fmt.Println(eventInstrument)
		c.JSON(200, eventInstrument)
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Should be min than " + fe.Param()
	case "max":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error1"
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// сохраняем или обновляем новость
func SaveNews(context *gin.Context) {

	rand.Seed(time.Now().UnixNano())

	eventInput := models.EventInput{}

	if err := context.ShouldBindJSON(&eventInput); err != nil {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	if eventInput.Date == "" {
		context.JSON(400, gin.H{"error": "Fields date are empty"})
		return
	}

	//c.JSON(201, "asdasd")
	//return

	reversed := revertDate(eventInput.Date)

	if eventInput.Hash != "" {
		userId := 101
		slug := slug.Make(eventInput.Title)
		updateData := models.Events{
			Title:        eventInput.Title,
			TypeID:       eventInput.TypeID,
			Shorttext:    eventInput.Shorttext,
			Fulltext:     eventInput.Fulltext,
			InstrumentID: eventInput.InstrumentID,
			Date:         reversed,
			Source:       eventInput.Source,
			Slug:         slug,
			UserId:       101,
		}

		errorUpdate := db.Model(&updateData).Where("hash = ? and user_id = ? and published = ?", eventInput.Hash, userId, "0").Updates(updateData)
		fmt.Println(errorUpdate)

		fmt.Println("update new news")
		if errorUpdate.RowsAffected == 0 {

			fmt.Println("create")
		} else {
			content := &models.EventResult{
				InstrumentID: eventInput.InstrumentID,
				Slug:         slug,
				Status:       "Updates",
				Hash:         eventInput.Hash,
			}

			context.JSON(200, content)
			fmt.Println("finish")
			return
		}
	}

	fmt.Println("eventInput.InstrumentID" + eventInput.InstrumentID)
	error, instrument := GetInstrumentInstanse(eventInput.InstrumentID)

	if !error {

		context.JSON(400, gin.H{"error": "Fields are empty 2"})

	} else {

		if eventInput.Date != "" && eventInput.TypeID != 0 && eventInput.Source != "" && eventInput.InstrumentID != "" &&
			eventInput.Title != "" && eventInput.Shorttext != "" && eventInput.Fulltext != "" {
			// Получаем уникальный урл
			slug := slug.Make(eventInput.Title)
			// получаем уникальный хэш
			hash := fmt.Sprintf("%x", md5.Sum([]byte(eventInput.Title+eventInput.Date+string(rand.Intn(50)))))
			// Преобразуем дату  ПРиходит 01/03/2022 записываем 2019-12-29
			// re := regexp.MustCompile(`([0-9]{2})/([0-9]{2})/([0-9]{4})`)
			// rD := (re.FindAllStringSubmatch(eventInput.Date, -1))
			// reversed := fmt.Sprintf("%s-%s-%s", rD[0][3], rD[0][2], rD[0][1])

			EventCreate := models.EventsResult{
				Date:         reversed,
				Slug:         slug,
				Hash:         hash,
				UserId:       100,
				TypeID:       eventInput.TypeID,
				Title:        eventInput.Title,
				InstrumentID: instrument.InstrumentID,
				Source:       eventInput.Source,
				Shorttext:    eventInput.Shorttext,
				Fulltext:     eventInput.Fulltext,
				Published:    "0",
			}

			err := db.Table("events").Create(EventCreate).Error

			if err != nil {

				//	return nil, err
			}
			fmt.Println(err)
			fmt.Println("========err")

			content := &models.EventResult{
				InstrumentID: eventInput.InstrumentID,
				Slug:         slug,
				Status:       "Create",
				Hash:         hash,
			}

			context.JSON(200, content)

		} else {
			fmt.Println("||| 400 400")
			context.JSON(400, gin.H{"error": "Fields are empty 1"})
		}
		//	return chartData
	}
}

// конвертируем время
func getHash(eventInput models.EventInput) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(eventInput.Title+eventInput.Date+string(rand.Intn(50)))))
}

// конвертируем время
func revertDate(date string) string {

	fmt.Println(" 1 date " + date)
	fmt.Println(" 1 date " + date)
	fmt.Println(" 1 date " + date)
	// re := regexp.MustCompile(`([0-9]{4})-([0-9]{2})-([0-9]{2})`)
	// rD := (re.FindAllStringSubmatch(date, -1))
	// return fmt.Sprintf("%s-%s-%s", rD[0][3], rD[0][2], rD[0][1])

	re := regexp.MustCompile(`([0-9]{2})/([0-9]{2})/([0-9]{4})`)
	rD := (re.FindAllStringSubmatch(date, -1))
	fmt.Println(rD)
	fmt.Println(rD[0][3], rD[0][2], rD[0][1])
	dateResult := fmt.Sprintf("%s-%s-%s", rD[0][3], rD[0][2], rD[0][1])

	fmt.Println("dateResult " + dateResult)
	return dateResult
}

// конвертируем время
func revertDateFromBase(date string) string {
	if date == "" {
		return ""
	}
	fmt.Println("date:: " + date)
	fmt.Println("date:: " + date)
	fmt.Println("date:: " + date)
	fmt.Println("date:: " + date)
	// re := regexp.MustCompile(`([0-9]{4})-([0-9]{2})-([0-9]{2})`)
	// rD := (re.FindAllStringSubmatch(date, -1))
	// return fmt.Sprintf("%s-%s-%s", rD[0][3], rD[0][2], rD[0][1])

	re := regexp.MustCompile(`([0-9]{4})-([0-9]{2})-([0-9]{2})`)
	rD := (re.FindAllStringSubmatch(date, -1))
	fmt.Println("rD")
	fmt.Println(rD)
	fmt.Println(rD[0][3], rD[0][2], rD[0][1])
	dateResult := fmt.Sprintf("%s/%s/%s", rD[0][3], rD[0][2], rD[0][1])

	fmt.Println("dateResult:: " + dateResult)
	return dateResult
}
