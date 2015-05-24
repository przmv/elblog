package main

import (
	"errors"

	"github.com/codegangsta/cli"
)

var (
	ErrMutuallyExclusiveFlags = errors.New("mutually exclusive flags")
)

func Before(c *cli.Context) error {
	if err := validateInterval(c); err != nil {
		return err
	}
	return nil
}

func validateInterval(c *cli.Context) error {
	hourly := c.GlobalBool("hourly")
	daily := c.GlobalBool("daily")
	interval := c.GlobalDuration("interval")
	if hourly && daily {
		return ErrMutuallyExclusiveFlags
	}
	if interval != flagInterval.Value && (hourly || daily) {
		return ErrMutuallyExclusiveFlags
	}
	return nil
}
