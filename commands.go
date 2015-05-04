package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/pshevtsov/gonx"
)

var Commands = []cli.Command{
	commandToken,
	commandIp,
	commandError,
	commandLatency,
}

var commandToken = cli.Command{
	Name:  "token",
	Usage: "",
	Description: `
`,
	Action: doToken,
}

var commandIp = cli.Command{
	Name:  "ip",
	Usage: "",
	Description: `
`,
	Action: doIp,
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
		log.Fatal(err)
	}
}

func doToken(c *cli.Context) {
	count, ok := <-gonx.MapReduce(Reader, Parser, new(TokenCount))
	if !ok {
		log.Fatal("Error occured")
	}
	format := c.GlobalString("output")
	output(format, count)
}

func doIp(c *cli.Context) {
}

func doError(c *cli.Context) {
}

func doLatency(c *cli.Context) {
}
