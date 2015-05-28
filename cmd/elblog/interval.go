package main

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/codegangsta/cli"
	"github.com/pshevtsov/elblog"
	"github.com/pshevtsov/gonx"
)

var (
	ErrMutuallyExclusiveFlags = errors.New("mutually exclusive flags")
)

const (
	IntervalHour = -1 * time.Hour
	IntervalDay  = -24 * time.Hour
)

func NewInterval(d time.Duration, t time.Time) *gonx.Interval {
	return &gonx.Interval{
		Field:  elblog.FieldTimestamp,
		Format: time.RFC3339Nano,
		Start:  t.Add(d),
		End:    t,
	}
}

var commandInterval = cli.Command{
	Name:  "interval",
	Usage: "Filter log entries by the specified interval",
	Description: `The command 'interval' outputs the log entries for
   the specified duration before now.

   If '--hour' flag was specified, the output will contain log entries for
   the previous hour.

   Specifying the '--day' flag will display the log entries since the day before now.

   To output the log entries for the custom duration, use '--duration' flag with
   the custom duration value.

   The value for the '--duration' flug must conform the input for the
   'time.ParseDuration()' Go function (http://golang.org/pkg/time/#ParseDuration).

   A duration string is a possibly signed sequence of decimal numbers, each with optional
   fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units
   are "ns", "us" (or "µs"), "ms", "s", "m", "h".`,
	Flags:  intervalFlags,
	Action: doInterval,
}

var intervalFlags = []cli.Flag{
	intervalFlagDuration,
	intervalFlagHour,
	intervalFlagDay,
}

var intervalFlagDuration = cli.DurationFlag{
	Name:  "duration",
	Value: 0,
	Usage: "output log entries for this previous interval",
}

var intervalFlagHour = cli.BoolFlag{
	Name:  "hour",
	Usage: "output log entries for the previous hour",
}

var intervalFlagDay = cli.BoolFlag{
	Name:  "day",
	Usage: "output log entries for the previous day",
}

func doInterval(c *cli.Context) {
	duration, err := getDuration(c)
	assert(err)
	var reducer gonx.Reducer
	if duration != 0 {
		reducer = NewInterval(duration, time.Now())
	} else {
		reducer = nil
	}
	reader := NewReader(c, reducer)
	for {
		entry, err := reader.Read()
		if err == io.EOF {
			break
		}
		assert(err)
		line, err := entry.Field("raw_line")
		assert(err)
		fmt.Println(line)
	}
}

// getDuration calculates interval value since interval can be set with
// --hour and --day boolean flags
func getDuration(c *cli.Context) (duration time.Duration, err error) {
	if err = validateDuration(c); err != nil {
		return
	}
	if c.Bool("hour") {
		return IntervalHour, nil
	}
	if c.Bool("day") {
		return IntervalDay, nil
	}
	return c.Duration("duration") * -1, nil
}

func validateDuration(c *cli.Context) error {
	hour := c.Bool("hour")
	day := c.Bool("day")
	duration := c.Duration("duration")
	if hour && day {
		return ErrMutuallyExclusiveFlags
	}
	if duration != intervalFlagDuration.Value && (hour || day) {
		return ErrMutuallyExclusiveFlags
	}
	return nil
}
