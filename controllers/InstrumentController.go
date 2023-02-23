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

// получаем список инструментов
func InstrumentsList(c *gin.Context) {
	var instrumentsList []models.Instrument
	if err := db.Table("instruments").Where("instruments.published = ?", "1").Find(&instrumentsList).Error; err != nil {
		fmt.Println(err)
		//c.JSON(404, gin.H{"error": "Instrument not found"})
	}
	c.JSON(200, instrumentsList)
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

//type New struct {
// EventID   int    `json:"eventId"`
// TypeID    int    `json:"typeId"`
// Hash      string `json:"hash"`
// Source    string `json:"source"`
// Slug      string `json:"slug"`
// Date      string `json:"date"`
// Title     string `json:"title"`
// Shorttext string `json:"shorttext"`
// Fulltext  string `json:"fulltext"`

// получаем список инструментов
func InstrumentTickerPrice(c *gin.Context) {
	ticker := c.Param("ticker")

	var instrument models.Instrument
	if err := db.Table("instruments").Where("instruments.ticker = ?", ticker).Find(&instrument).Error; err != nil {
		fmt.Println(err)
		c.JSON(400, "Instrument not found")
		return
	}
	// SELECT price, prices.date, events.title, events.slug, events.hash, events.type_id, events.event_id,source,instrument_id,shorttext

	// FROM `prices`
	// left join events on prices.date = events.date   AND events.instrument_id = '1' and published = 1
	//  WHERE prices.name = 'RASP'
	//  ORDER BY prices.date

	var prices []models.Prices
	// .Joins("left join emails on emails.user_id = users.id").Scan(&result{})
	if err := db.Table("prices").Select("price, prices.date, events.title, events.slug, events.hash, events.type_id, events.event_id,source,instrument_id,shorttext").Joins("left join events on events.date = prices.date and events.instrument_id = ? and published = ? ", instrument.InstrumentID, 1).Where("prices.name = ? ", ticker).Order("prices.date asc").Find(&prices).Error; err != nil {
		fmt.Println(err)
		//c.JSON(404, gin.H{"error": "Instrument not found"})
	}
	//fmt.Println(prices)
	c.JSON(200, prices)
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

		//	filename := header.Filename
		//		fmt.Println(handler.Filename)

		// fmt.Printf("Uploaded file name: %+v\n", handler.Filename)
		// fmt.Printf("Uploaded file size %+v\n", handler.Size)
		// fmt.Printf("File mime type %+v\n", handler.Header)

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
