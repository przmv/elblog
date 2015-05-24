package main

import (
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "elblog"
	app.Version = Version
	app.Usage = "Parse and analyze AWS Elastic Load Balancing access logs"
	app.Author = "Petr Shevtsov"
	app.Email = "petr.shevtsov@gmail.com"
	app.Flags = Flags
	app.Commands = Commands
	app.Before = Before

	app.RunAndExitOnError()
}
