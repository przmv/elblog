package main

import (
	"fmt"
	"io"

	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog"
	"github.com/pshevtsov/gonx"
)

var commandStatus = cli.Command{
	Name:  "status",
	Usage: "Count status codes",
	Description: `The command 'status' outputs the list of different status codes
   along with the total amount of requests.

   By default it displays the backend status codes.

   If the '--elb' flag was added, status codes for Elastic Load Balancing
   will be outputted.`,
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
		fmt.Println(k, v)
	}
}
