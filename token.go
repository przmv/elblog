package main

import (
	"net/url"
	"strings"

	"github.com/pshevtsov/gonx"
)

// Implements Reducer interface for summarize Token URL values for the request field
type TokenCount struct{}

func (r *TokenCount) Reduce(input chan *gonx.Entry, output chan *gonx.Entry) {
	sum := make(map[string]uint64)
	for entry := range input {
		req, err := entry.Field("request")
		if err != nil {
			panic(err)
		}
		token := getToken(req)
		if token != "" {
			sum[token]++
		}
	}
	entry := gonx.NewEmptyEntry()
	for name, val := range sum {
		entry.SetUintField(name, val)
	}
	output <- entry
	close(output)
}

// Get token query parameter from the request field
func getToken(s string) string {
	parts := strings.Split(s, " ")
	u, err := url.Parse(parts[1])
	if err != nil {
		return ""
	}
	return u.Query().Get("token")
}
