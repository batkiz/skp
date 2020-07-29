package main

import (
	"github.com/batkiz/skp/commands"
	"github.com/batkiz/skp/lib"
	"github.com/batkiz/skp/utils"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if len(args) >= 2 {
		switch args[1] {
		case "info":
			dir, err := filepath.Abs(os.Args[2])
			utils.HandleErr(err)

			commands.Info(dir)
		case "home":
			dir, err := filepath.Abs(os.Args[2])
			utils.HandleErr(err)

			app := lib.ReadManifest(dir)

			commands.Home(app)
		case "bucket":
			bargs := args[2:]
			commands.Bucket(bargs)
		}
	}
}
