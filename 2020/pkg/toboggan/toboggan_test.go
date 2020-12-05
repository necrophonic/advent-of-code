package toboggan_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/toboggan"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	toboggan.DataFile = "../../data/toboggan.txt"
	p1, p2, err := toboggan.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 292, p1, "Part 1")
	assert.Equal(t, 9354744432, p2, "Part 2")
}
