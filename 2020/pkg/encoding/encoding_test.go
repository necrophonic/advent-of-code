package encoding_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/encoding"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	encoding.DataFile = "../../data/encoding.txt"
	p1, p2, err := encoding.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 50047984, p1, "Part 1")
	assert.Equal(t, 5407707, p2, "Part 2")
}

var testInputNumbers = []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

func TestFindWeakness(t *testing.T) {
	assert.Equal(t, 127, encoding.FindWeakness(testInputNumbers, 5))
}

func TestFindContiguous(t *testing.T) {
	assert.Equal(t, 62, encoding.FindContiguous(testInputNumbers, 127))
}
