package output

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/template"
	"unicode/utf8"
)

type Output struct {
	Writer io.Writer

	rows     [][]string
	lengths  map[int]int
	hasNames bool
}

func (o *Output) Add(row ...string) {
	o.rows = append(o.rows, row)
	o.updateLengths(row)
}

func (o *Output) SetNames(names ...string) {
	o.updateLengths(names)
	if o.hasNames {
		o.rows[0] = names
	} else {
		o.rows = append([][]string{names}, o.rows...)
		o.hasNames = true
	}
}

func (o *Output) updateLengths(row []string) {
	if len(o.lengths) == 0 {
		o.lengths = make(map[int]int)
	}
	for i, col := range row {
		length := utf8.RuneCountInString(col)
		if o.lengths[i] < length {
			o.lengths[i] = length
		}
	}
}

func (o *Output) Table(spaces int) {
	if o.Writer == nil {
		o.Writer = os.Stdout
	}
	for i, row := range o.rows {
		for j, col := range row {
			n := strconv.Itoa(o.lengths[j] + spaces)
			if o.hasNames && i == 0 {
				col = strings.ToUpper(col)
			}
			fmt.Fprintf(o.Writer, "%-"+n+"s", col)
		}
		fmt.Fprintln(o.Writer, "")
	}
}

func (o *Output) CSV() error {
	if o.Writer == nil {
		o.Writer = os.Stdout
	}
	w := csv.NewWriter(o.Writer)
	if o.hasNames {
		return w.WriteAll(o.rows[1:])
	}
	return w.WriteAll(o.rows)
}

func (o *Output) Template(s string) error {
	if o.Writer == nil {
		o.Writer = os.Stdout
	}
	t := template.Must(template.New("output").Parse(s))
	if o.hasNames {
		return o.namedDataTemplate(t)
	}
	return t.Execute(o.Writer, o.rows)
}

func (o *Output) namedDataTemplate(t *template.Template) error {
	names := o.rows[0]
	for i, name := range names {
		name := strings.ToLower(name)
		name = strings.Title(name)
		names[i] = strings.Replace(name, " ", "", -1)
	}
	data := []map[string]string{}
	for _, row := range o.rows[1:] {
		m := map[string]string{}
		for i, col := range row {
			m[names[i]] = col
		}
		data = append(data, m)
	}
	return t.Execute(o.Writer, data)
}
