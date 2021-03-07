package main

import (
	"fmt"
	"src/src/argos"
	"src/src/game"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	for {
		c := colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36"),
		)

		cGame := c.Clone()
		cArgos := c.Clone()

		game.Game(cGame)
		argos.Argos(cArgos)

		fmt.Println("")
		time.Sleep(30 * time.Second)
	}
}
