package main

import (
	"fmt"
	"io"

	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog"
	"github.com/pshevtsov/gonx"
)

var commandClientIP = cli.Command{
	Name:  "client-ip",
	Usage: "Count different client IP addresses",
	Description: `The command 'client-ip' outputs the list of different client
   IP addresses along with the total amount of requests.`,
	Action: doClientIP,
}

func doClientIP(c *cli.Context) {
	reducer := elblog.NewGroupByClientIP(new(gonx.Count))
	reader := NewReader(c, reducer)
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
		fmt.Println(k, v)
	}
}
