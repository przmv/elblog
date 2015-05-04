package main

import (
	"time"

	"github.com/codegangsta/cli"
)

var Flags = []cli.Flag{
	flagInterval,
	flagOutput,
}

var flagInterval = cli.DurationFlag{
	Name:  "interval",
	Value: time.Hour,
	Usage: "Analyze data in this interval",
}

var flagOutput = cli.StringFlag{
	Name:  "output",
	Value: "csv",
	Usage: "Specify the output format: csv or text",
}
