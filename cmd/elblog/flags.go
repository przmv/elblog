package main

import (
	"github.com/codegangsta/cli"
)

var Flags = []cli.Flag{
	flagCSV,
	flagTemplate,
}

var flagCSV = cli.BoolFlag{
	Name:  "csv",
	Usage: "output in CSV format",
}

var flagTemplate = cli.StringFlag{
	Name:  "template",
	Usage: "output the result of template execute",
}
