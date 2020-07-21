package commands

import (
	. "github.com/batkiz/skp/utils"
	"os/exec"
)

func Home(app string) {
	url := ""

	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	HandleErr(err)
}
