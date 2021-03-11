package very

import (
	"log"
	"src/src/helpers"

	"github.com/gocolly/colly"
)

const url = "https://www.very.co.uk/playstation-5-sign-up.page"

func Very(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", "Very...")
	})

	c.OnHTML("body", process)

	c.Visit(url)
}

func process(e *colly.HTMLElement) {
	childAttr := e.ChildAttr("[data-src=\"https://content.very.co.uk/assets/static/2020/09/events/ps5-slices/ps5-stock-coming-soon/registerinterestpage-desktop.jpg\"]", "data-src")

	if childAttr == "" {
		helpers.Notify("Very", url)
	}
}
