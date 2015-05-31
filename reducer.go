package elblog

import (
	"math"
	"net/url"
	"sort"
	"strconv"
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

var (
	FieldCount             = "count"
	FieldMinimum           = "min"
	FieldMaximum           = "max"
	FieldMean              = "mean"
	FieldStandardDeviation = "standard deviation"
)

func FieldPercentile(p float64) string {
	return "p" + strconv.Itoa(int(p*100))
}

type Latency struct {
	Percentiles []float64
}

func (r *Latency) Reduce(input chan *gonx.Entry, output chan *gonx.Entry) {
	var (
		min      float64 = math.MaxFloat64
		max      float64
		count    float64
		total    float64
		mean     float64
		vSum     float64
		variance float64
		values   []float64
	)
	for entry := range input {
		sum := entry.SumFields([]string{
			FieldRequestProcessingTime,
			FieldBackendProcessingTime,
			FieldResponseProcessingTime,
		})
		sum = sum * 1000 // ms
		values = append(values, sum)

		min = math.Min(min, sum)
		max = math.Max(max, sum)

		total += sum
		count++

		mean = total / count

		d := sum - mean
		vSum += d * d
		variance = vSum / count
	}
	entry := gonx.NewEmptyEntry()
	entry.SetUintField("count", uint64(count))
	entry.SetFloatField(FieldMinimum, min)
	entry.SetFloatField(FieldMaximum, max)
	entry.SetFloatField(FieldMean, mean)
	entry.SetFloatField(FieldStandardDeviation, math.Sqrt(variance))
	percentiles := r.calcPercentiles(values)
	for i, p := range r.Percentiles {
		k := FieldPercentile(p)
		v := percentiles[i]
		entry.SetFloatField(k, v)
	}
	output <- entry
	close(output)
}

func (r *Latency) calcPercentiles(values []float64) []float64 {
	scores := make([]float64, len(r.Percentiles))
	size := len(values)
	if size > 0 {
		sort.Float64s(values)
		for i, p := range r.Percentiles {
			pos := p * float64(size+1)
			if pos < 1.0 {
				scores[i] = values[0]
			} else if pos >= float64(size) {
				scores[i] = values[size-1]
			} else {
				lower := values[int(pos)-1]
				upper := values[int(pos)]
				scores[i] = lower + (pos-math.Floor(pos))*(upper-lower)
			}
		}
	}
	return scores
}
