package main

import (
	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog"
	"github.com/satyrius/gonx"
)

var commandRequest = cli.Command{
	Name:  "request",
	Usage: "Various HTTP request line analyses",
	Description: `Request Description
`,
	Flags:  requestFlags,
	Action: doRequest,
}

var requestFlags = []cli.Flag{
	requestFlagParam,
}

var requestFlagParam = cli.StringSliceFlag{
	Name:  "param",
	Usage: "Specify HTTP request parameter",
	Value: &cli.StringSlice{},
}

func doRequest(c *cli.Context) {
	params := c.StringSlice("param")
	// Show help if param flag is empty
	if len(params) == 0 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}
	reader, err := elblog.NewReader(c.Args())
	assert(err)
	parser := elblog.NewParser()
	for _, param := range params {
		reducer := elblog.NewRequestParamCount(param)
		count, ok := <-gonx.MapReduce(
			reader,
			parser,
			reducer,
		)
		debug(count, ok)
	}
}
