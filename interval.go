package elblog

import (
	"time"

	"github.com/pshevtsov/gonx"
)

const (
	IntervalHourly = -1 * time.Hour
	IntervalDaily  = -24 * time.Hour
)

func NewInterval(d time.Duration, t time.Time) *gonx.Interval {
	return &gonx.Interval{
		Field:  FieldTimestamp,
		Format: time.RFC3339Nano,
		Start:  t.Add(d),
		End:    t,
	}
}
