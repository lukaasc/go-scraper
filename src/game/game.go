package game

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/gocolly/colly"
)

type item struct {
	Image,
	Title,
	Strapline,
	Copy string
	Button button
}

type button struct {
	Copy, Link string
}

func Game(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", "Game...")
	})

	c.OnResponse(process)

	c.Visit("https://assets.game.net/_master/hardwarePages/playstation5/merch/LaunchCopyAvail/ps5.js?_=1614893095264")
}

func process(r *colly.Response) {
	regex, _ := regexp.Compile("(?:contentPanelsThree = )(\\[([^\\]])+\\])")

	var response []item

	js := strings.ReplaceAll(string(regex.FindSubmatch(r.Body)[1]), "'", "\"")

	json.Unmarshal([]byte(js), &response)

	if response[0].Button.Copy != "Out of stock" {
		beeep.Alert("Game", "", "")
		log.Println("Available at Game:", "https://www.game.co.uk/"+response[0].Button.Link)
	}
}
