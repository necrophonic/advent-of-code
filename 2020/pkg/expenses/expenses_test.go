package expenses_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/expenses"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	expenses.DataFile = "../../data/expenses.txt"
	p1, p2, err := expenses.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 870331, p1, "Part 1")
	assert.Equal(t, 283025088, p2, "Part 2")
}
