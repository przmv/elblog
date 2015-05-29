package main

import (
	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog/output"
)

func Output(c *cli.Context, o *output.Output) error {
	if c.GlobalBool(flagCSV.Name) {
		return o.CSV()
	}
	if t := c.GlobalString(flagTemplate.Name); t != "" {
		return o.Template(t)
	}
	o.Table(4)
	return nil
}
