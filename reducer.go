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

type GroupByClientIP struct {
	reducers []gonx.Reducer
}

func NewGroupByClientIP(reducers ...gonx.Reducer) *GroupByClientIP {
	return &GroupByClientIP{reducers: reducers}
}

var (
	FieldClientIP = "client_ip"
)

// Apply related reducers and group data by client IP.
func (r *GroupByClientIP) Reduce(input chan *gonx.Entry, output chan *gonx.Entry) {
	subInput := make(map[string]chan *gonx.Entry)
	subOutput := make(map[string]chan *gonx.Entry)

	// Read reducer master input channel and create discinct input chanel
	// for each entry key we group by
	for entry := range input {
		clientIPEntry := r.clientIPEntry(entry)
		key := clientIPEntry.FieldsHash([]string{FieldClientIP})
		if _, ok := subInput[key]; !ok {
			subInput[key] = make(chan *gonx.Entry, cap(input))
			subOutput[key] = make(chan *gonx.Entry, cap(output)+1)
			subOutput[key] <- clientIPEntry
			go gonx.NewChain(r.reducers...).Reduce(subInput[key], subOutput[key])
		}
		subInput[key] <- entry
	}
	for _, ch := range subInput {
		close(ch)
	}
	for _, ch := range subOutput {
		entry := <-ch
		entry.Merge(<-ch)
		output <- entry
	}
	close(output)
}

func (r *GroupByClientIP) clientIPEntry(entry *gonx.Entry) *gonx.Entry {
	client, err := entry.Field(FieldClient)
	if err != nil {
		return gonx.NewEmptyEntry()
	}
	clientIP := strings.Split(client, ":")[0]
	return gonx.NewEntry(gonx.Fields{
		FieldClientIP: clientIP,
	})
}
