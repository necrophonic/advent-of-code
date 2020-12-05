package passport_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/passport"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	passport.DataFile = "../../data/passport.txt"
	p1, p2, err := passport.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 226, p1, "Part 1")
	assert.Equal(t, 160, p2, "Part 2")
}
