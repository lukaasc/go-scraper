package main

import (
	"src/src/crawler"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	for {
		c := colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36"),
		)

		cGame := c.Clone()
		crawler.Game(cGame)

		time.Sleep(30 * time.Second)
	}
}
