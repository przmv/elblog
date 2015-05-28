package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	//"github.com/pshevtsov/gonx"
)

var Commands = []cli.Command{
	commandInterval,
	commandRequestParam,
	commandClientIP,

	commandError,
	commandLatency,
}

var commandError = cli.Command{
	Name:  "error",
	Usage: "",
	Description: `
`,
	Action: doError,
}

var commandLatency = cli.Command{
	Name:  "latency",
	Usage: "",
	Description: `
`,
	Action: doLatency,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func doError(c *cli.Context) {
}

func doLatency(c *cli.Context) {
}
