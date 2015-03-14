package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Version = Version
	app.Usage = ""
	app.Author = "Yoshiki Aoki"
	app.Email = "yoshiki_aoki@dwango.co.jp"
	app.Commands = Commands

	app.Run(os.Args)
}
