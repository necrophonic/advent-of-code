package adaptor_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/adaptor"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	adaptor.DataFile = "../../data/adaptor.txt"
	p1, p2, err := adaptor.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 1917, p1, "Part 1")
	assert.Equal(t, 0, p2, "Part 2")
}

func TestJoltDifferences(t *testing.T) {
	tcs := []struct {
		adaptors []int
		expected int
	}{
		{[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, 35},
		{[]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}, 220},
	}

	for _, test := range tcs {
		assert.Equal(t, test.expected, adaptor.JoltDifferences(test.adaptors))
	}
}

// func TestCombinations(t *testing.T) {
// 	tcs := []struct {
// 		adaptors []int
// 		expected int
// 	}{
// 		// {[]int{5, 1, 4, 7, 2}, 8},
// 		{[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, 8},
// 		// 31	// {[]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}, 19208},
// 	}

// 	for _, test := range tcs {
// 		assert.Equal(t, test.expected, adaptor.Combinations(test.adaptors))
// 	}
// }
