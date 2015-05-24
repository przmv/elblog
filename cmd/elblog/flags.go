package main

import (
	"github.com/codegangsta/cli"
)

var Flags = []cli.Flag{
	flagInterval,
	flagHourly,
	flagDaily,
	flagOutput,
}

var flagInterval = cli.DurationFlag{
	Name:  "interval",
	Value: 0,
	Usage: "analyze data in this interval",
}

var flagHourly = cli.BoolFlag{
	Name:  "hourly",
	Usage: "analyze data for the previous hour",
}

var flagDaily = cli.BoolFlag{
	Name:  "daily",
	Usage: "analyze data for the previous day",
}

var flagOutput = cli.StringFlag{
	Name:  "output",
	Value: "csv",
	Usage: "specify the output format: csv or text",
}
