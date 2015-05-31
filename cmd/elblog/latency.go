package main

import (
	"io"

	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog"
	"github.com/pshevtsov/elblog/output"
)

var commandLatency = cli.Command{
	Name:  "latency",
	Usage: "Calculate latency statistics",
	Description: `The command 'latency' outputs some latency statistics data:
      Total amount of requests,
      Minimal processing time,
      Processing time maximum,
      Mean processing time,
      Latency standard deviation,
      Percentile latency stats (p50, p75 and p99).

   Processing time is the sum of Request processing time, Backend processing time
   and Response processing time values and is represented in milliseconds.

   If global flag '--csv' was set, the first column of the output represents the total
   amount of requests. The second and the third column show the processing time
   minimum and maximum respectively. The column number four consists of mean processing
   time values. Latency standard deviation is stored in the fifth column. The last three
   columns show the percentile latency stats in the following order: p50, p75 and p99.

   If no arguments were provided, the command reads standard input.

TEMPLATE DATA:
   If using global flag '--template', the following data type is sent to the template
   to execute:

   map[string]string

   The possible map keys are: 'Count', 'Min', 'Max', 'Mean', 'StandardDeviation',
   'P50', 'P75' and 'P99'.

   The example template is the following:

   The minimum processing time of {{.Count}} requests is {{.Min}}

   See https://golang.org/pkg/text/template/ for the reference.`,
	Action: doLatency,
}

func doLatency(c *cli.Context) {
	reducer := &elblog.Latency{Percentiles: []float64{0.5, 0.75, 0.99}}
	reader := NewReader(c, reducer)
	o := &output.Output{IsSingle: true}
	fields := []string{
		elblog.FieldCount,
		elblog.FieldMinimum,
		elblog.FieldMaximum,
		elblog.FieldMean,
		elblog.FieldStandardDeviation,
	}
	for _, p := range reducer.Percentiles {
		fields = append(fields, elblog.FieldPercentile(p))
	}
	o.SetNames(fields...)
	for {
		entry, err := reader.Read()
		if err == io.EOF {
			break
		}
		assert(err)
		values := []string{}
		for _, field := range fields {
			val, err := entry.Field(field)
			assert(err)
			values = append(values, val)
		}
		o.Add(values...)
	}
	err := Output(c, o)
	assert(err)
}
