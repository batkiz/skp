package main

import (
	"github.com/batkiz/skp/commands"
	. "github.com/batkiz/skp/utils"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if args[1] == "info" {
		dir, err := filepath.Abs(os.Args[2])
		HandleErr(err)

		commands.Info(dir)
	}
}
