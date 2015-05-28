package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog"
)

var commandRequestParam = cli.Command{
	Name:  "request-param",
	Usage: "Count request parameters",
	Description: `The command 'request-param' outputs the list of different
   request parameter's values along with the total amount of these values in the
   request path field across the log entries provided.

   The required '--param' flag is used to specify the HTTP request parameter.
`,
	Flags:  requestParamFlags,
	Action: doRequestParam,
}

var requestParamFlags = []cli.Flag{
	requestParamFlagParam,
}

var requestParamFlagParam = cli.StringFlag{
	Name:  "param",
	Usage: "Specify HTTP request parameter",
}

func doRequestParam(c *cli.Context) {
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
