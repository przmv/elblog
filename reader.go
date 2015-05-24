package elblog

import (
	"io"
	"os"
)

func NewReader(args []string) (io.Reader, error) {
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
