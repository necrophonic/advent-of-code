package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestSyntaxScoringAnswer(t *testing.T) {
	s := &day.SyntaxScoring{}
	part1, part2, err := s.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 358737, part1)
	assert.Equal(t, 4329504793, part2)
}

func TestCheckLine(t *testing.T) {
	type tc struct {
		input    string
		expected int
	}

	tcs := []tc{
		{"{([(<{}[<>[]}>{[]{[(<()>", 1197},
		{"[[<[([]))<([[{}[[()]]]", 3},
		{"[{[{({}]{}}([{[{{{}}([]", 57},
		{"<{([([[(<>()){}]>(<<{{", 25137},
	}

	s := &day.SyntaxScoring{}

	for _, test := range tcs {
		assert.Equal(t, test.expected, s.CheckLine(test.input))
	}
}

func TestFixLine(t *testing.T) {
	type tc struct {
		input    string
		expected int
	}

	tcs := []tc{
		{"[({(<(())[]>[[{[]{<()<>>", 288957},
		{"<{([{{}}[<[[[<>{}]]]>[]]", 294},
	}

	s := &day.SyntaxScoring{}

	for _, test := range tcs {
		assert.Equal(t, test.expected, s.FixLine(test.input))
	}
}
