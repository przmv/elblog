package elblog

import (
	"fmt"

	"github.com/pshevtsov/gonx"
)

var (
	FieldTimestamp              = "timestamp"
	FieldELB                    = "elb"
	FieldClient                 = "client"
	FieldBackend                = "backend"
	FieldRequestProcessingTime  = "request_processing_time"
	FieldBackendProcessingTime  = "backend_processing_time"
	FieldResponseProcessingTime = "response_processing_time"
	FieldELBStatus              = "elb_status_code"
	FieldBackendStatus          = "backend_status_code"
	FieldReceivedBytes          = "received_bytes"
	FieldSentBytes              = "sent_bytes"
	FieldRequest                = "request"
)

type field string

func (f field) String() string {
	s := string(f)
	if s == FieldRequest {
		return fmt.Sprintf(`"$%s"`, s)
	}
	return fmt.Sprint("$", s)
}

func NewParser() *gonx.Parser {
	slice := []field{
		field(FieldTimestamp),
		field(FieldELB),
		field(FieldClient),
		field(FieldBackend),
		field(FieldRequestProcessingTime),
		field(FieldBackendProcessingTime),
		field(FieldResponseProcessingTime),
		field(FieldELBStatus),
		field(FieldBackendStatus),
		field(FieldReceivedBytes),
		field(FieldSentBytes),
		field(FieldRequest),
	}
	format := fmt.Sprint(slice)
	format = format[1 : len(format)-1]
	return gonx.NewParser(format)
}
