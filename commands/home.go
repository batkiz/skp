package commands

import (
	"github.com/batkiz/skp/lib"
	. "github.com/batkiz/skp/utils"
	"os/exec"
)

// Home is the base func for `scoop home <app>` command
func Home(app lib.Manifest) {
	url := *app.Homepage
	
	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	HandleErr(err)
}
