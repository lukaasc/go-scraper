package argos

import (
	"log"

	"github.com/gen2brain/beeep"
	"github.com/gocolly/colly"
)

func Argos(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", "Argos...")
	})

	c.OnHTML("body", process)

	c.Visit("https://www.argos.co.uk/product/8349000")
}

func process(e *colly.HTMLElement) {
	childText := e.ChildText("#subCopy")
	text := "We're working hard to make this available as soon as possible."

	if childText != text {
		beeep.Alert("Argos", "", "")
		log.Println("Available at Argos:", "https://www.argos.co.uk/product/8349000")
	}
}
