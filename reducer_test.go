package elblog_test

import (
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/pshevtsov/elblog"
	"github.com/pshevtsov/gonx"
)

func TestCountRequestParam(t *testing.T) {
	f, err := os.Open("testdata/count_request_param.log")
	if err != nil {
		t.Fatal(err)
	}
	param := "test_param"
	reducer := elblog.NewRequestParamCount(param)
	reader := elblog.NewReader(f, reducer)
	entry, err := reader.Read()
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
	want := gonx.NewEntry(gonx.Fields{
		"foo":   "2",
		"bar":   "1",
		"baz":   "1",
		"quuix": "1",
	})
	if !reflect.DeepEqual(entry, want) {
		t.Fatalf("want %v got %v", want, entry)
	}
}
