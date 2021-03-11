package amazon

import (
	"log"

	"strings"

	"github.com/gen2brain/beeep"
	"github.com/gocolly/colly"
)

func Amazon(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", "Amazon...")
	})

	c.OnHTML("body", process)

	c.Visit("https://www.amazon.co.uk/PlayStation-9395003-5-Console/dp/B08H95Y452")
}

func process(e *colly.HTMLElement) {
	availabilityMessage := e.ChildText("#availability")

	isUnavailable := strings.Contains(availabilityMessage, "Currently unavailable") && strings.Contains(availabilityMessage, "We don't know when or if this item will be back in stock.")

	if !isUnavailable {
		beeep.Alert("Amazon", "", "")
		log.Println("Available at Amazon:", "https://www.amazon.co.uk/PlayStation-9395003-5-Console/dp/B08H95Y452")
	}
}
