package controllers

import (
	"fmt"
	"log"
	"moex/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sabloger/sitemap-generator/smg"
)

func CreateSitemaps(c *gin.Context) {
	GetSitemapEvents(c)
	GetSitemapInstruments(c)
}

// получаем новости и строим сайтмап
func GetSitemapEvents(c *gin.Context) {
	//<loc>https://www.example.com/some/uri.html</loc>

	var news []models.SitemapNews
	//now := time.Now().UTC()
	server_host := os.Getenv("SERVER_HOST")
	sitemap_path := os.Getenv("SITEMAP_PATH")

	sm := smg.NewSitemap(true)
	sm.SetName("sitemap_event")
	sm.SetHostname(server_host)
	sm.SetOutputPath(sitemap_path)
	//	sm.SetLastMod(&now)
	sm.SetCompress(false) // Default is true

	//.Joins("JOIN department on department.id = employee.department_id")
	request := db.Table("events").Select("events.type_id, events.create_at, events.slug, events.date, events.title, instruments.ticker")

	request = request.Where("events.published = 1")

	request = request.Joins("LEFT JOIN instruments on instruments.instrument_id = events.instrument_id")

	if err := request.Find(&news).Error; err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "news not found"})
	} else {

		for i := 0; i < len(news); i++ {
			// CreateAt
			///events/RASP/pavel-basinskii-v-iskusstve-tsel-opravdyvaet-sredstva
			err := sm.Add(&smg.SitemapLoc{
				Loc: "events/" + news[i].Ticker + "/" + news[i].Slug,
				//		LastMod:    &now,
				ChangeFreq: smg.Always,
				Priority:   0.6,
			})
			if err != nil {
				log.Fatal("Unable to add SitemapLoc:", err)
			}
		}
		filenames, err := sm.Save()
		if err != nil {
			log.Fatal("Unable to Save Sitemap:", err)
		}
		for i, filename := range filenames {
			fmt.Println("file no.", i+1, filename)
		}
	}
}

// получаем новости и строим сайтмап
func GetSitemapInstruments(c *gin.Context) {
	//<loc>https://www.example.com/some/uri.html</loc>

	var instruments []models.SitemapInstrument
	//now := time.Now().UTC()
	server_host := os.Getenv("SERVER_HOST")
	sitemap_path := os.Getenv("SITEMAP_PATH")

	sm := smg.NewSitemap(true)
	sm.SetName("sitemap_instruments")
	sm.SetHostname(server_host)
	sm.SetOutputPath(sitemap_path)
	//sm.SetLastMod(&now)
	sm.SetCompress(false) // Default is true

	//.Joins("JOIN department on department.id = employee.department_id")
	request := db.Table("instruments").Select("instruments.type,   instruments.ticker")
	request = request.Where("instruments.published = 1")
	if err := request.Find(&instruments).Error; err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "news not found"})
	} else {

		for i := 0; i < len(instruments); i++ {
			// CreateAt
			///events/RASP/pavel-basinskii-v-iskusstve-tsel-opravdyvaet-sredstva
			err := sm.Add(&smg.SitemapLoc{
				Loc: instruments[i].Type + "/" + instruments[i].Ticker,
				//	LastMod:    &now,
				ChangeFreq: smg.Always,
				//		Priority:   0.4,
			})
			if err != nil {
				log.Fatal("Unable to add SitemapLoc:", err)
			}
		}
		filenames, err := sm.Save()
		if err != nil {
			log.Fatal("Unable to Save Sitemap:", err)
		}
		for i, filename := range filenames {
			fmt.Println("file no.", i+1, filename)
		}
	}
}
