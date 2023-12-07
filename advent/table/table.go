package table

import (
	"fmt"
	"sort"
	"strings"
)

type Row struct {
	Day    int
	V1, V2 any
	Err    error
}

type Table struct {
	rows                             map[int]Row
	longest1, longest2, longestError int
	hadError                         bool
}

func New() *Table {
	return &Table{
		longest1:     len("Part 1"),
		longest2:     len("Part 2"),
		longestError: len("ERROR"),
		rows:         make(map[int]Row),
	}
}

func (t *Table) Add(i int, v1, v2 any, err error) {
	if err != nil {
		t.hadError = true
		t.rows[i] = Row{
			Day: i,
			V1:  "n/a",
			V2:  "n/a",
			Err: err,
		}
		if l := len(err.Error()); l > t.longestError {
			t.longestError = l
		}
		return
	}
	t.rows[i] = Row{
		Day: i,
		V1:  v1,
		V2:  v2,
		Err: err,
	}

	if l := len(fmt.Sprint(v1)); l > t.longest1 {
		t.longest1 = l
	}
	if l := len(fmt.Sprint(v2)); l > t.longest2 {
		t.longest2 = l
	}
}

func (t *Table) String() string {
	sb := &strings.Builder{}

	sb.WriteString("| Day |")

	p1fmt := " %" + fmt.Sprint(t.longest1) + "v |"
	p2fmt := " %" + fmt.Sprint(t.longest2) + "v |"
	// errfmt := " %" + fmt.Sprint(t.longestError) + "v |"

	sb.WriteString(fmt.Sprintf(p1fmt, "Part 1"))
	sb.WriteString(fmt.Sprintf(p2fmt, "Part 2"))
	if t.hadError {
		sb.WriteString(" ERROR")
	}
	sb.WriteString("\n")
	sb.WriteString("+-----+")
	sb.WriteString(strings.Repeat("-", t.longest1+2) + "+")
	sb.WriteString(strings.Repeat("-", t.longest2+2) + "+")
	if t.hadError {
		sb.WriteString(strings.Repeat("-", t.longestError+2) + "+")
	}
	sb.WriteString("\n")

	keys := make([]int, 0, len(t.rows))
	for k := range t.rows {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		r := t.rows[k]
		sb.WriteString(fmt.Sprintf("|  %2d |", r.Day))
		sb.WriteString(fmt.Sprintf(p1fmt, r.V1))
		sb.WriteString(fmt.Sprintf(p2fmt, r.V2))
		if r.Err != nil {
			sb.WriteString(" " + r.Err.Error())
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
