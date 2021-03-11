package argos

import (
	"log"
	"src/src/helpers"

	"github.com/gocolly/colly"
)

const url = "https://www.argos.co.uk/product/8349000"

func Argos(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", "Argos...")
	})

	c.OnHTML("body", process)

	c.Visit(url)
}

func process(e *colly.HTMLElement) {
	childText := e.ChildText("#subCopy")
	text := "We're working hard to make this available as soon as possible."

	if childText != text {
		helpers.Notify("Argos", url)
	}
}
