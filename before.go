package main

import (
	"io"
	"os"

	"github.com/codegangsta/cli"
)

var Reader io.Reader

func Before(c *cli.Context) error {
	var files []io.Reader
	for _, arg := range c.Args().Tail() {
		file, err := os.Open(arg)
		assert(err)
		files = append(files, file)
	}
	Reader = io.MultiReader(files...)
	// TODO apply global interval
	return nil
}
