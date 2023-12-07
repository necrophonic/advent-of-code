package password_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/password"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	password.DataFile = "../../data/password.txt"
	p1, p2, err := password.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 410, p1, "Part 1")
	assert.Equal(t, 694, p2, "Part 2")
}
