package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "elb-log-analyzer"
	app.Version = Version
	app.Usage = ""
	app.Author = "Petr Shevtsov"
	app.Email = "petr.shevtsov@gmail.com"
	app.Flags = Flags
	app.Commands = Commands
	app.Before = Before

	app.Run(os.Args)
}
