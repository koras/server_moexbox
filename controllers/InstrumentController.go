package controllers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"moex/models"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	}

	if instrument.InstrumentID == "" {
		return false, instrument
	}

	return true, instrument
}

type DashBordResult struct {
	Instrument []models.Instrument     `json:"instrument"`
	Prices     []models.PricesDashBord `json:"prices"`
}

// получаем список инструментов для дожборда
func InstrumentsList(c *gin.Context) {
	//typeId=1&level=2
	typeId := c.Query("typeId")
	level := c.Query("level")

	var instrumentsList []models.Instrument

	request := db.Table("instruments").Where("instruments.published = ?", "1")
	fmt.Println("typeId" + typeId)

	if typeId != "all" && typeId != "0" {
		request = request.Where("instruments.type = ?", typeId)

		fmt.Println("level |" + level)
		if typeId == "shares" && (level == "1" || level == "2" || level == "3") {
			fmt.Println("level = " + level)
			request = request.Where("instruments.level = ?", level)
		}
	}

	if err := request.Limit(50).Offset(0).Find(&instrumentsList).Error; err != nil {
		fmt.Println(err)
		c.JSON(200, err)
		return
		//c.JSON(404, gin.H{"error": "Instrument not found"})
	}

	var instrumentsArray []string
	for i := 0; i < len(instrumentsList); i++ {
		instrumentsArray = append(instrumentsArray, instrumentsList[i].Ticker)
	}
	//fmt.Println(instrumentsArray)
	var pricesDashBord []models.PricesDashBord

	now := time.Now()
	fmt.Println("Today:", now)

	after := now.AddDate(-1, 0, 0)

	date := after.Format("2006-01-02")
	fmt.Println("Subtract 1 Year:", date)

	//	fmt.Println(instrumentsArray)

	if err := db.Table("prices").Where("name  IN  (?) and `prices`.`date` > ?", instrumentsArray, date).Find(&pricesDashBord).Error; err != nil {
		fmt.Println("get pricess info")
		fmt.Println(err)
		c.JSON(200, err)
		return
	}

	dashBordResult := DashBordResult{
		Instrument: instrumentsList,
		Prices:     pricesDashBord,
	}

	c.JSON(200, dashBordResult)
}

// получаем список инструментов
func InstrumentGet(c *gin.Context) {
	InstrumentId := c.Param("InstrumentId")

	var instrument models.Instrument

	if err := db.Table("instruments").Where("instruments.published = ? and instruments.instrument_id = ?", "1", InstrumentId).Find(&instrument).Error; err != nil {
		fmt.Println(err)
		//c.JSON(404, gin.H{"error": "Instrument not found"})
	}

	c.JSON(200, instrument)
}

// получаем список инструментов
func InstrumentTickerPrice(c *gin.Context) {
	ticker := c.Param("ticker")

	var instrument models.Instrument
	if err := db.Table("instruments").
		Where("instruments.ticker = ?", ticker).
		Find(&instrument).Error; err != nil {

		fmt.Println("get instrument ")
		fmt.Println(err)
		c.JSON(400, err)
		return
	}

	var prices []models.Prices

	if err := db.Table("prices").
		Select("price, `prices`.`date`, events.title, events.slug, events.hash, events.type_id, events.event_id, source, instrument_id, shorttext").
		Joins("left join events on events.date = `prices`.`date` and events.instrument_id = ? and published = ? ", instrument.InstrumentID, 1).
		Where("prices.name = ? ", ticker).Order(" `prices`.`date` asc ").Find(&prices).Error; err != nil {

		fmt.Println("InstrumentTickerPrice")
		fmt.Println(err)
		c.JSON(400, "Instrument not found  InstrumentTickerPrice 2")
		return
	}

	PricesInstrument := models.PricesInstrument{

		Instruments: instrument,
		Prices:      prices,
	}
	c.JSON(200, PricesInstrument)
}

// получаем список дат торгов
func InstrumentOnlyDateFromPricesWeeks(ticker string, instrumentID string, date string) (bool, []models.Prices) {

	var priceDate []models.Prices

	//	SELECT * FROM table_name
	//WHERE date_column BETWEEN DATE_SUB('2023-04-12', INTERVAL 1 WEEK) AND DATE_ADD('2023-04-12', INTERVAL 1 WEEK);

	if err := db.Table("prices").
		Select("price, `prices`.`date`, events.title, events.slug, events.hash, events.type_id, events.event_id,source,instrument_id,shorttext").
		Joins("left join events on events.date = `prices`.`date` and events.instrument_id = ? and published = ? ", instrumentID, 1).
		Where("prices.name = ? and `prices`.`date` BETWEEN DATE_SUB(?, INTERVAL 1 WEEK) AND DATE_ADD(?, INTERVAL 1 WEEK)", ticker, date, date).
		Order("`prices`.`date` ASC").
		Find(&priceDate).Error; err != nil {

		fmt.Println("InstrumentOnlyDateFromPricesWeeks ")
		fmt.Println(err)
		return false, priceDate
	}
	return true, priceDate
	//fmt.Println(prices)

}

// получаем список дат торгов
func InstrumentOnlyDateFromPrices(ticker string) (bool, []models.PriceDate) {

	var priceDate []models.PriceDate
	if err := db.Table("prices").Select("date").Where("name = ?", ticker).Find(&priceDate).Error; err != nil {
		fmt.Println(err)
		return false, priceDate
	}
	return true, priceDate
	//fmt.Println(prices)

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
func InstrumentUpdate(context *gin.Context) {
	// InstrumentInput
	logo := ""
	InstrumentInput := models.InstrumentInput{}

	fmt.Println("InstrumentInput")

	fmt.Println("parser")
	if err := context.ShouldBind(&InstrumentInput); err != nil {

		fmt.Println("valid")
		fmt.Println(err)
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}

				//	fmt.Println(getErrorMsg(fe))
				//	context.JSON(200, fe)
			}
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})

		}
		fmt.Println("valid 111")
		return
	}

	file, handler, err := context.Request.FormFile("upload")
	if err != nil {
		fmt.Println("go---------------")
		fmt.Println(err)
	} else {

		fmt.Println(" handler.Filename " + handler.Filename)

		logo = InstrumentInput.InstrumentID + path.Ext(handler.Filename)
		// Create a temporary file with a dir folder
		// tempFile, err: = ioutil.TempFile("temp-files", fileName)
		out, err := os.Create("tmp/" + InstrumentInput.InstrumentID + path.Ext(handler.Filename))

		if err != nil {
			fmt.Println(err)
		}
		defer out.Close()

		//fileBytes, err: = ioutil.ReadAll(file)
		//if err != nil {
		//	  fmt.Println(err)
		//	}

		//tempFile.Write(fileBytes)
		fmt.Println("Successfully uploaded file")

		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}

	if logo != "" {
		InstrumentInput.Logo = logo
	}

	errorUpdate := db.Table("instruments").Where("instrument_id = ? ", InstrumentInput.InstrumentID).Updates(InstrumentInput)
	fmt.Println(errorUpdate)

	context.JSON(200, "ok")
}
