package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"moex/controllers"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// https://www.moex.com/ru/marketdata/
// получаем стоимость всех бумаг
// https://iss.moex.com/iss/history/engines/stock/markets/shares/boardgroups/57/securities.jsonp?iss.meta=off&iss.json=extended&date=2023-02-08&start=400&limit=100&sort_column=VALUE&sort_order=des

//https://iss.moex.com/iss/engines/stock/markets/shares/boardgroups/57/securities.jsonp?iss.meta=off&iss.json=extended&callback=JSON_CALLBACK&lang=ru&security_collection=3&sort_column=VALTODAY&sort_order=desc

func GetPriceMoexHistory(c *gin.Context) {
	t := time.Date(2023, time.February, 9, 23, 0, 0, 0, time.UTC)
	for day := 0; day <= 2; day++ {
		t2 := t.AddDate(0, 0, day)
		date := t2.Format("2006-01-02")
		fmt.Println(date)
		for pageStart := 0; pageStart <= 600; pageStart += 100 {
			url := fmt.Sprintf("https://iss.moex.com/iss/history/engines/stock/markets/shares/boardgroups/57/securities.jsonp?iss.meta=off&iss.json=extended&date=%v&start=%v&limit=100&sort_order=desc", date, pageStart)
			fmt.Println(url)
			res, err := http.Get(url)
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Printf("fail %s ", url)
				panic(err)
			}
			content := string(body)
			contentReplace := ClearText(content)
			if contentReplace != "" {
				fmt.Print("\r")
				var moex LASTVOLUME
				if err := json.Unmarshal([]byte(contentReplace), &moex); err != nil { // Parse []byte to go struct pointer

					fmt.Printf("fail 2 %s ", url)
					panic(err)
				}
				for _, rec := range moex.History {
					if rec.Close != 0 {
						prices := controllers.Prices{
							Name:  rec.Secid,
							Date:  rec.Tradedate,
							Price: rec.Close,
							Value: rec.Value,
						}
						controllers.SaveParice(prices)
					}
				}
			}
		}

	}
	c.JSON(200, "ok")
}

func GetPriceMoexOnline(c *gin.Context) {

	url := fmt.Sprintf("https://iss.moex.com/iss/engines/stock/markets/shares/boardgroups/57/securities.jsonp?iss.meta=off&iss.json=extended&lang=ru&iss.json=extended&security_collection=3&sort_column=VALTODAY&sort_order=desc")
	fmt.Println(url)
	res, err := http.Get(url)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("fail %s ", url)
		panic(err)
	}
	content := string(body)
	contentReplace := ClearText(content)

	//fmt.Println(contentReplace)
	if contentReplace != "" {
		fmt.Print("\r")
		var moex ONLINE
		if err := json.Unmarshal([]byte(contentReplace), &moex); err != nil { // Parse []byte to go struct pointer
			fmt.Printf("fail 2 %s ", url)
			panic(err)
		}
		for _, rec := range moex.Securities {
			//		c.JSON(200, rec)
			if rec.Prevprice != 0 {

				controllers.UpdatePrice(rec.Secid, rec.Prevprice)
				//SECID
				//		prices := controllers.Prices{
				//			Name:  rec.Secid,
				//			Date:  rec.Tradedate,
				//			Price: rec.Close,
				//			Value: rec.Value,
				//		}
				//		controllers.SaveParice(prices)
			}
		}
	}
	c.JSON(200, "ok")
}

func ClearText(content string) string {
	//	re := regexp.MustCompile("JSON_CALLBACK\\(")
	//replaced := re.ReplaceAllString(content, "")

	re3 := regexp.MustCompile("\\{\"charsetinfo\"\\: \\{\"name\": \"utf-8\"\\}\\},")
	replaced3 := re3.ReplaceAllString(content, "")
	re2 := regexp.MustCompile("\\)$")
	replaced2 := re2.ReplaceAllString(replaced3, "")
	replaced4 := strings.TrimSpace(replaced2)
	re4 := regexp.MustCompile("^\\[")
	replaced5 := re4.ReplaceAllString(replaced4, "")
	re5 := regexp.MustCompile("\\]$")
	replaced6 := re5.ReplaceAllString(replaced5, "")

	re6 := regexp.MustCompile("\n$")
	contentReplace := re6.ReplaceAllString(replaced6, "")

	//	fmt.Print(contentReplace)

	return contentReplace
}

type LASTVOLUME struct {
	History []struct {
		Boardid                 string      `json:"BOARDID"`
		Tradedate               string      `json:"TRADEDATE"`
		Shortname               string      `json:"SHORTNAME"`
		Secid                   string      `json:"SECID"`
		Numtrades               int64       `json:"NUMTRADES"`
		Value                   float64     `json:"VALUE"`
		Open                    interface{} `json:"OPEN"`
		Low                     interface{} `json:"LOW"`
		High                    interface{} `json:"HIGH"`
		Legalcloseprice         float64     `json:"LEGALCLOSEPRICE"`
		Waprice                 interface{} `json:"WAPRICE"`
		Close                   float64     `json:"CLOSE"`
		Volume                  int64       `json:"VOLUME"`
		Marketprice2            interface{} `json:"MARKETPRICE2"`
		Marketprice3            interface{} `json:"MARKETPRICE3"`
		Admittedquote           float64     `json:"ADMITTEDQUOTE"`
		Mp2Valtrd               float64     `json:"MP2VALTRD"`
		Marketprice3Tradesvalue float64     `json:"MARKETPRICE3TRADESVALUE"`
		Admittedvalue           float64     `json:"ADMITTEDVALUE"`
		Waval                   float64     `json:"WAVAL"`
		Tradingsession          int         `json:"TRADINGSESSION"`
	} `json:"history,omitempty"`
}

type ONLINE struct {
	Securities []struct {
		Secid               string      `json:"SECID"`
		Boardid             string      `json:"BOARDID"`
		Shortname           string      `json:"SHORTNAME"`
		Prevprice           float64     `json:"PREVPRICE"`
		Lotsize             int         `json:"LOTSIZE"`
		Facevalue           float64     `json:"FACEVALUE"`
		Status              string      `json:"STATUS"`
		Boardname           string      `json:"BOARDNAME"`
		Decimals            int         `json:"DECIMALS"`
		Secname             string      `json:"SECNAME"`
		Remarks             interface{} `json:"REMARKS"`
		Marketcode          string      `json:"MARKETCODE"`
		Instrid             string      `json:"INSTRID"`
		Sectorid            interface{} `json:"SECTORID"`
		Minstep             float64     `json:"MINSTEP"`
		Prevwaprice         float64     `json:"PREVWAPRICE"`
		Faceunit            string      `json:"FACEUNIT"`
		Prevdate            string      `json:"PREVDATE"`
		Issuesize           int64       `json:"ISSUESIZE"`
		Isin                string      `json:"ISIN"`
		Latname             string      `json:"LATNAME"`
		Regnumber           string      `json:"REGNUMBER"`
		Prevlegalcloseprice float64     `json:"PREVLEGALCLOSEPRICE"`
		Prevadmittedquote   interface{} `json:"PREVADMITTEDQUOTE"`
		Currencyid          string      `json:"CURRENCYID"`
		Sectype             string      `json:"SECTYPE"`
		Listlevel           int         `json:"LISTLEVEL"`
		Settledate          string      `json:"SETTLEDATE"`
	} `json:"securities"`
	// Marketdata []struct {
	// 	Secid                         string      `json:"SECID"`
	// 	Boardid                       string      `json:"BOARDID"`
	// 	Bid                           float64     `json:"BID"`
	// 	Biddepth                      interface{} `json:"BIDDEPTH"`
	// 	Offer                         interface{} `json:"OFFER"`
	// 	Offerdepth                    interface{} `json:"OFFERDEPTH"`
	// 	Spread                        int         `json:"SPREAD"`
	// 	Biddeptht                     int         `json:"BIDDEPTHT"`
	// 	Offerdeptht                   int         `json:"OFFERDEPTHT"`
	// 	Open                          int         `json:"OPEN"`
	// 	Low                           float64     `json:"LOW"`
	// 	High                          float64     `json:"HIGH"`
	// 	Last                          float64     `json:"LAST"`
	// 	Lastchange                    float64     `json:"LASTCHANGE"`
	// 	Lastchangeprcnt               float64     `json:"LASTCHANGEPRCNT"`
	// 	Qty                           int         `json:"QTY"`
	// 	Value                         float64     `json:"VALUE"`
	// 	ValueUsd                      float64     `json:"VALUE_USD"`
	// 	Waprice                       float64     `json:"WAPRICE"`
	// 	Lastcngtolastwaprice          float64     `json:"LASTCNGTOLASTWAPRICE"`
	// 	Waptoprevwapriceprcnt         float64     `json:"WAPTOPREVWAPRICEPRCNT"`
	// 	Waptoprevwaprice              float64     `json:"WAPTOPREVWAPRICE"`
	// 	Closeprice                    interface{} `json:"CLOSEPRICE"`
	// 	Marketpricetoday              interface{} `json:"MARKETPRICETODAY"`
	// 	Marketprice                   float64     `json:"MARKETPRICE"`
	// 	Lasttoprevprice               float64     `json:"LASTTOPREVPRICE"`
	// 	Numtrades                     int         `json:"NUMTRADES"`
	// 	Voltoday                      int         `json:"VOLTODAY"`
	// 	Valtoday                      int64       `json:"VALTODAY"`
	// 	ValtodayUsd                   int         `json:"VALTODAY_USD"`
	// 	Etfsettleprice                interface{} `json:"ETFSETTLEPRICE"`
	// 	Tradingstatus                 string      `json:"TRADINGSTATUS"`
	// 	Updatetime                    string      `json:"UPDATETIME"`
	// 	Admittedquote                 interface{} `json:"ADMITTEDQUOTE"`
	// 	Lastbid                       interface{} `json:"LASTBID"`
	// 	Lastoffer                     interface{} `json:"LASTOFFER"`
	// 	Lcloseprice                   interface{} `json:"LCLOSEPRICE"`
	// 	Lcurrentprice                 float64     `json:"LCURRENTPRICE"`
	// 	Marketprice2                  interface{} `json:"MARKETPRICE2"`
	// 	Numbids                       interface{} `json:"NUMBIDS"`
	// 	Numoffers                     interface{} `json:"NUMOFFERS"`
	// 	Change                        float64     `json:"CHANGE"`
	// 	Time                          string      `json:"TIME"`
	// 	Highbid                       interface{} `json:"HIGHBID"`
	// 	Lowoffer                      interface{} `json:"LOWOFFER"`
	// 	Priceminusprevwaprice         float64     `json:"PRICEMINUSPREVWAPRICE"`
	// 	Openperiodprice               int         `json:"OPENPERIODPRICE"`
	// 	Seqnum                        int64       `json:"SEQNUM"`
	// 	Systime                       string      `json:"SYSTIME"`
	// 	Closingauctionprice           float64     `json:"CLOSINGAUCTIONPRICE"`
	// 	Closingauctionvolume          int         `json:"CLOSINGAUCTIONVOLUME"`
	// 	Issuecapitalization           int64       `json:"ISSUECAPITALIZATION"`
	// 	IssuecapitalizationUpdatetime string      `json:"ISSUECAPITALIZATION_UPDATETIME"`
	// 	Etfsettlecurrency             interface{} `json:"ETFSETTLECURRENCY"`
	// 	ValtodayRur                   int64       `json:"VALTODAY_RUR"`
	// 	Tradingsession                string      `json:"TRADINGSESSION"`
	// } `json:"marketdata"`
	// Dataversion []struct {
	// 	DataVersion int   `json:"data_version"`
	// 	Seqnum      int64 `json:"seqnum"`
	// } `json:"dataversion"`
	// MarketdataYields []interface{} `json:"marketdata_yields"`
}
