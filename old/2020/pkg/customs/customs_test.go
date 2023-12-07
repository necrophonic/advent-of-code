package customs_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/customs"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	customs.DataFile = "../../data/customs.txt"
	p1, p2, err := customs.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 6885, p1, "Part 1")
	assert.Equal(t, 3550, p2, "Part 2")
}

func TestCountAnyAnswered(t *testing.T) {
	type tc struct {
		group    []byte
		expected int
	}

	tcs := []tc{
		{[]byte("abc"), 3},
		{[]byte("a\nb\nc"), 3},
		{[]byte("ab\nac"), 3},
		{[]byte("a\na\na\na"), 1},
		{[]byte("b"), 1},
	}

	all := 0
	for _, test := range tcs {
		count := customs.CountAnyAnswered(test.group)
		all += count
		assert.Equalf(t, test.expected, count, "Answer for %s", test.group)
	}
	assert.Equal(t, 11, all, "Total count correct")
}

func TestCountAllAnswered(t *testing.T) {
	type tc struct {
		group    []byte
		expected int
	}

	tcs := []tc{
		{[]byte("abc"), 3},
		{[]byte("a\nb\nc"), 0},
		{[]byte("ab\nac"), 1},
		{[]byte("a\na\na\na"), 1},
		{[]byte("b"), 1},
		{[]byte("rgvby\nfbgdyirx\nbpmgtnujrzyoq\nbarygw"), 4},
		{[]byte("wmfpvn\nwhbzmvjplc\nvwpsmk\nsovwpm\nmsvrpwdf"), 4},
	}

	all := 0
	for _, test := range tcs {
		count := customs.CountAllAnswered(test.group)
		all += count
		assert.Equalf(t, test.expected, count, "Answer for %s", test.group)
	}
	assert.Equal(t, 14, all, "Total count correct")
}
