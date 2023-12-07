package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestSevenSegmentAnswer(t *testing.T) {
	s := &day.SevenSegment{}
	part1, part2, err := s.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 0, part1)
	assert.Equal(t, 0, part2)
}

var sevenSegmentTestData = []string{
	"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
	// "edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
	// "fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
	// "fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
	// "aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
	// "fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
	// "dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
	// "bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
	// "egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
	// "gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
}

func TestSevenSegment_Part1(t *testing.T) {
	s := &day.SevenSegment{}
	answer, err := s.Part1(sevenSegmentTestData)
	if assert.NoError(t, err) {
		assert.Equal(t, 26, answer)
	}
}

func TestSevenSegment_Part2(t *testing.T) {
	s := &day.SevenSegment{}
	answer, err := s.Part2(sevenSegmentTestData)
	if assert.NoError(t, err) {
		assert.Equal(t, 26, answer)
	}
}

func TestSevenSegment_DiffDigits(t *testing.T) {
	s := &day.SevenSegment{}

	assert.Equal(t, "a", s.DiffDigits("abc", "bc"))
	// assert.Equal(t, "ac", s.DiffDigits("abc", "b"))
	// assert.Equal(t, "be", s.DiffDigits("acdeg", "acdfg", "abdfg"))

	assert.Equal(t, "a", s.DiffDigits("abcd", "bc", "cd"))
}
