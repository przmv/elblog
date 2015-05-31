package main

import (
	"io"

	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog"
	"github.com/pshevtsov/elblog/output"
	"github.com/pshevtsov/gonx"
)

var commandStatus = cli.Command{
	Name:  "status",
	Usage: "Count status codes",
	Description: `The command 'status' outputs the list of different status codes
   along with the total amount of requests.

   By default it displays the backend status codes.

   If the '--elb' flag was added, status codes for Elastic Load Balancing
   will be outputted.

   If global flag '--csv' was set, the first column of the output denotes status code
   and the second one is the count.

   If no arguments were provided, the command reads standard input.

TEMPLATE DATA:
   If using global flag '--template', the following data type is sent to the template
   to execute:

   []map[string]string

   The possible map keys are 'Status' and 'Count'.

   The example template is the following:

   {{range $i, $r := .}}
       {{$i}}. {{$r.Status}} ({{$r.Count}})
   {{end}}

   See https://golang.org/pkg/text/template/ for the reference.`,
	Flags:  statusFlags,
	Action: doStatus,
}

var statusFlags = []cli.Flag{
	statusFlagELB,
}

var statusFlagELB = cli.BoolFlag{
	Name:  "elb",
	Usage: "display status codes report for Elastic Load Balancing",
}

func doStatus(c *cli.Context) {
	fields := []string{elblog.FieldBackendStatus}
	if c.Bool(statusFlagELB.Name) {
		fields = []string{elblog.FieldELBStatus}
	}
	reducer := gonx.NewGroupBy(fields, new(gonx.Count))
	reader := NewReader(c, reducer)
	o := &output.Output{}
	o.SetNames("status", "count")
	for {
		entry, err := reader.Read()
		if err == io.EOF {
			break
		}
		assert(err)
		k, err := entry.Field(fields[0])
		assert(err)
		v, err := entry.Field("count")
		assert(err)
		o.Add(k, v)
	}
	err := Output(c, o)
	assert(err)
}
