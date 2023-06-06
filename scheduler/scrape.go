package scheduler

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/st4rgaze/otaqu/app/models"
	"github.com/st4rgaze/otaqu/utils"
)

func Run() {
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c.OnHTML(".col-avail.hotel", func(e *colly.HTMLElement) {
		hotel := models.Hotel{
			Name:     e.ChildText("h3"),
			Address:  e.ChildText(".loct"),
			ImageUrl: e.ChildAttr("img.img-hotel", "src"),
		}

		// star rating count
		e.ForEach("i.fas.fa-star.star-hotel", func(_ int, star *colly.HTMLElement) {
			hotel.StarRating++
		})

		// convert price to int
		priceStr := e.ChildText(".price-hotel > h6")
		hotel.Price = utils.ConvertPriceToUint(priceStr)

		// create/update hotel
		err := hotel.Create()
		if err != nil {
			log.Println(err)
		}
	})

	// run every SCRAPE_SECONDS in .env
	scrapeSecondsStr := os.Getenv("SCRAPE_SECONDS")
	scrapeSeconds, err := strconv.Atoi(scrapeSecondsStr)
	if err != nil {
		log.Fatalf("Invalid SCRAPE_SECONDS value: %s", scrapeSecondsStr)
	}

	scrapeInterval := time.Duration(scrapeSeconds) * time.Second
	ticker := time.NewTicker(scrapeInterval)

	for range ticker.C {
		err := c.Visit(os.Getenv("SCRAPE_URL"))
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Scraping page")
		}
	}
}
