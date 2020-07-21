package commands

import (
	"log"
	"os/exec"
)

func home(app string) {
	url := ""

	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	if err != nil {
		log.Println(err)
	}
}
