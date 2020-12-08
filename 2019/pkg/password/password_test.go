package password_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2019/pkg/password"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	p1, p2, err := password.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 1873, p1, "Part 1")
	assert.Equal(t, 1264, p2, "Part 2")
}

func TestMatchCriteria(t *testing.T) {
	type tc struct {
		password string
		expected int
	}

	tcs := []tc{
		{"111111", 1},
		{"223450", 0},
		{"123789", 0},
		{"123779", 1},
	}

	for _, test := range tcs {
		assert.Equal(t, test.expected, password.MatchCriteria(test.password))
	}
}

func TestExtendedMatchCriteria(t *testing.T) {
	type tc struct {
		password string
		expected int
	}

	tcs := []tc{
		{"112233", 1},
		{"111234", 0},
		{"112345", 1},
		{"123444", 0},
		{"111122", 1},
		{"136788", 1},
		{"136777", 0},
		{"589999", 0},
	}

	for _, test := range tcs {
		match := password.MatchExtendedCriteria(test.password)
		assert.Equal(t, test.expected, match, test.password)
		t.Logf("%s = %d", test.password, match)
	}
}
