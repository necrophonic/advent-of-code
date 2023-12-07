package trebuchet_test

import (
	"testing"

	trebuchet "github.com/necrohonic/advent-of-code/2023/go/01-trebuchet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	treb := &trebuchet.Trebuchet{
		D: []string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet",
		},
	}
	ans, err := treb.Part1()
	require.NoError(t, err, "Should not error")
	require.IsType(t, ans, int(0))
	assert.Equal(t, 142, ans.(int), "Expected answer")
}

func TestPart2(t *testing.T) {
	t.Run("Standard test", func(t *testing.T) {
		treb := &trebuchet.Trebuchet{
			D: []string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
		}
		ans, err := treb.Part2()
		require.NoError(t, err, "Should not error")
		require.IsType(t, ans, int(0))
		assert.Equal(t, 281, ans.(int), "Expected answer")
	})

	t.Run("Overlapping numbers", func(t *testing.T) {
		treb := &trebuchet.Trebuchet{
			D: []string{
				"eightfive37688eighteightwof",
			},
		}
		ans, err := treb.Part2()
		require.NoError(t, err, "Should not error")
		require.IsType(t, ans, int(0))
		assert.Equal(t, 82, ans.(int), "Expected answer")
	})
}
