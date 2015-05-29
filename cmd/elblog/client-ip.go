package main

import (
	"io"

	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog"
	"github.com/pshevtsov/elblog/output"
	"github.com/pshevtsov/gonx"
)

var commandClientIP = cli.Command{
	Name:  "client-ip",
	Usage: "Count different client IP addresses",
	Description: `The command 'client-ip' outputs the list of different client
   IP addresses along with the total amount of requests.

   If global flag '--csv' was set, the first column of the output denotes the IP address
   and the second one is the count.

TEMPLATE DATA:
   If using global flag '--template', the following data type is sent to the template
   to execute:

   []map[string]string

   The possible map keys are 'ClientIp' and 'Count'.

   The example template is the following:

   {{range $i, $r := .}}
       {{$i}}. {{$r.ClientIp}} ({{$r.Count}})
   {{end}}

   See https://golang.org/pkg/text/template/ for the reference.`,
	Action: doClientIP,
}

func doClientIP(c *cli.Context) {
	reducer := elblog.NewGroupByClientIP(new(gonx.Count))
	reader := NewReader(c, reducer)
	o := &output.Output{}
	o.SetNames("client ip", "count")
	for {
		entry, err := reader.Read()
		if err == io.EOF {
			break
		}
		assert(err)
		k, err := entry.Field(elblog.FieldClientIP)
		assert(err)
		v, err := entry.Field("count")
		assert(err)
		o.Add(k, v)
	}
	err := Output(c, o)
	assert(err)
}
