package main

import (
	"github.com/codegangsta/cli"
)

var Flags = []cli.Flag{
	flagOutput,
}

var flagOutput = cli.StringFlag{
	Name:  "output",
	Value: "csv",
	Usage: "specify the output format: csv or text",
}
