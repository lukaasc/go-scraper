package helpers

import (
	"log"
	"os/exec"

	"github.com/gen2brain/beeep"
)

func Notify(store, url string) {
	beeep.Alert(store, "", "")
	log.Println("Available at", store, ":", url)
	exec.Command("open", url).Start()
}
