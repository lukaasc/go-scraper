package currys

import (
	"log"

	"github.com/gen2brain/beeep"
	"github.com/gocolly/colly"
)

func Currys(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", "Currys...")
	})

	c.OnHTML("body", process)

	c.Visit("https://www.currys.co.uk/gbuk/playstation-5-sony-1714-commercial.html")
}

func process(e *colly.HTMLElement) {
	soldOutImageAltText := e.ChildAttr(".sold-out-banner img[alt=\"Playstation 5 sold out\"]", "alt")

	if soldOutImageAltText == "" {
		beeep.Alert("Currys", "", "")
		log.Println("Available at Currys:", "https://www.currys.co.uk/gbuk/playstation-5-sony-1714-commercial.html")
	}
}
