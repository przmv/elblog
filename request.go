package elblog

import (
	"net/url"
	"strings"

	"github.com/satyrius/gonx"
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
			panic(err)
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
	u, err := url.Parse(parts[1])
	if err != nil {
		return ""
	}
	return u.Query().Get(r.param)
}
