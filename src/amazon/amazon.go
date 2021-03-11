package amazon

import (
	"log"
	"src/src/helpers"

	"strings"

	"github.com/gocolly/colly"
)

const url = "https://www.amazon.co.uk/PlayStation-9395003-5-Console/dp/B08H95Y452"

func Amazon(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", "Amazon...")
	})

	c.OnHTML("body", process)

	c.Visit(url)
}

func process(e *colly.HTMLElement) {
	availabilityMessage := e.ChildText("#availability")

	isUnavailable := strings.Contains(availabilityMessage, "Currently unavailable") && strings.Contains(availabilityMessage, "We don't know when or if this item will be back in stock.")

	if !isUnavailable {
		helpers.Notify("Amazon", url)
	}
}
