package very

import (
	"log"

	"github.com/gen2brain/beeep"
	"github.com/gocolly/colly"
)

func Very(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", "Very...")
	})

	c.OnHTML("body", process)

	c.Visit("https://www.very.co.uk/playstation-5-sign-up.page")
}

func process(e *colly.HTMLElement) {
	childAttr := e.ChildAttr("[data-src=\"https://content.very.co.uk/assets/static/2020/09/events/ps5-slices/ps5-stock-coming-soon/registerinterestpage-desktop.jpg\"]", "data-src")

	if childAttr == "" {
		beeep.Alert("Very", "", "")
		log.Println("Available at Very:", "https://www.very.co.uk/playstation-5-sign-up.page")
	}
}
