package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/pshevtsov/gonx"
)

func output(format string, entry *gonx.Entry) {
	format = strings.ToLower(format)
	switch format {
	case "csv":
		outputCSV(entry)
	case "text":
		log.Fatal("Not implemented")
	default:
		log.Fatal("Invalid output format:", format)
	}
}

func outputCSV(entry *gonx.Entry) {
	w := csv.NewWriter(os.Stdout)
	for k, v := range entry.Fields() {
		err := w.Write([]string{k, v})
		assert(err)
	}
	w.Flush()
	assert(w.Error())
}
