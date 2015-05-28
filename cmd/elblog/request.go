package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog"
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

var requestFlagParam = cli.StringFlag{
	Name:  "param",
	Usage: "Specify HTTP request parameter",
}

func doRequest(c *cli.Context) {
	param := c.String("param")
	// Show help if param flag is empty
	if param == "" {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}
	reducer := elblog.NewRequestParamCount(param)
	reader := NewReader(c, reducer)
	entry, err := reader.Read()
	assert(err)
	for k, v := range entry.Fields() {
		fmt.Println(k, v)
	}
}
