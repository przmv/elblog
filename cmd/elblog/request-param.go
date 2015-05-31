package main

import (
	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog"
	"github.com/pshevtsov/elblog/output"
)

var commandRequestParam = cli.Command{
	Name:  "request-param",
	Usage: "Count request parameters",
	Description: `The command 'request-param' outputs the list of different
   request parameter's values along with the total amount of these values in the
   request path field across the log entries provided.

   The required '--param' flag is used to specify the HTTP request parameter.

   If global flag '--csv' was set, the first column of the output denotes the request
   parameter value and the second one is the count.

   If no arguments were provided, the command reads standard input.

TEMPLATE DATA:
   If using global flag '--template', the following data type is sent to the template
   to execute:

   []map[string]string

   The possible map keys are 'RequestParameterValue' and 'Count'.

   The example template is the following:

   {{range $i, $r := .}}
       {{$i}}. {{$r.RequestParameterValue}} ({{$r.Count}})
   {{end}}

   See https://golang.org/pkg/text/template/ for the reference.`,
	Flags:  requestParamFlags,
	Action: doRequestParam,
}

var requestParamFlags = []cli.Flag{
	requestParamFlagParam,
}

var requestParamFlagParam = cli.StringFlag{
	Name:  "param",
	Usage: "HTTP request parameter",
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
	o := &output.Output{}
	o.SetNames("request parameter value", "count")
	for k, v := range entry.Fields() {
		o.Add(k, v)
	}
	err = Output(c, o)
	assert(err)
}
