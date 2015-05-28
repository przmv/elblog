package elblog

import (
	"net/url"
	"strings"

	"github.com/pshevtsov/gonx"
)

type RequestParamCount struct {
	param string
}

func NewRequestParamCount(param string) *RequestParamCount {
	return &RequestParamCount{param}
}

func (r *RequestParamCount) Reduce(input chan *gonx.Entry, output chan *gonx.Entry) {
	sum := make(map[string]uint64)
	for entry := range input {
		req, err := entry.Field(FieldRequest)
		if err != nil {
			continue
		}
		param := r.getParamValue(req)
		if param != "" {
			sum[param]++
		}
	}
	entry := gonx.NewEmptyEntry()
	for name, val := range sum {
		entry.SetUintField(name, val)
	}
	output <- entry
	close(output)
}

// Get query parameter from the request field
func (r RequestParamCount) getParamValue(s string) string {
	parts := strings.Split(s, " ")
	u, _ := url.Parse(parts[1])
	return u.Query().Get(r.param)
}
