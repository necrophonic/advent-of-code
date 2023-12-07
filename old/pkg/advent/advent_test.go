package advent_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/pkg/advent"
	"github.com/stretchr/testify/assert"
)

func TestByteToInt(t *testing.T) {
	assert.Equal(t, 1, advent.ByteToInt('1'))
	assert.Equal(t, 9, advent.ByteToInt('9'))
}
