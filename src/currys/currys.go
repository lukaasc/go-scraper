package currys

import (
	"log"
	"src/src/helpers"

	"github.com/gocolly/colly"
)

const url = "https://www.currys.co.uk/gbuk/playstation-5-sony-1714-commercial.html"

func Currys(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", "Currys...")
	})

	c.OnHTML("body", process)

	c.Visit(url)
}

func process(e *colly.HTMLElement) {
	soldOutImageAltText := e.ChildAttr(".sold-out-banner img[alt=\"Playstation 5 sold out\"]", "alt")

	if soldOutImageAltText == "" {
		helpers.Notify("Currys:", url)
	}
}
