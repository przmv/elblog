package main

import (
	"io"
	"os"

	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog"
	"github.com/pshevtsov/gonx"
)

func NewReader(c *cli.Context, reducer gonx.Reducer) *elblog.Reader {
	f, err := file(c.Args())
	assert(err)
	if reducer == nil {
		reducer = new(gonx.ReadAll)
	}
	return elblog.NewReader(f, reducer)
}

func file(args []string) (io.Reader, error) {
	var files []io.Reader
	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	if len(files) > 0 {
		return io.MultiReader(files...), nil
	}
	return os.Stdin, nil
}
