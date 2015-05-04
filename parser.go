package main

import (
	"github.com/pshevtsov/gonx"
)

const (
	parserFormat = `$timestamp $elb $client $backend $request_processing_time $backend_processing_time $response_processing_time $elb_status_code $backend_status_code $received_bytes $sent_bytes "$request"`
)

var Parser = gonx.NewParser(parserFormat)
