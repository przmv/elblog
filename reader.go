package elblog

import (
	"io"

	"github.com/pshevtsov/gonx"
)

type Reader struct {
	file    io.Reader
	parser  *gonx.Parser
	reducer gonx.Reducer
	entries chan *gonx.Entry
}

func NewReader(file io.Reader, reducer gonx.Reducer) *Reader {
	return &Reader{
		file:    file,
		parser:  NewParser(),
		reducer: reducer,
	}
}

func (r *Reader) Read() (entry *gonx.Entry, err error) {
	if r.reducer == nil {
		r.reducer = new(gonx.ReadAll)
	}
	if r.entries == nil {
		r.entries = gonx.MapReduce(r.file, r.parser, r.reducer)
	}
	entry, ok := <-r.entries
	if !ok {
		err = io.EOF
	}
	return
}
