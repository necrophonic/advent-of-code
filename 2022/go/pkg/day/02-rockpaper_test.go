package day_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2022/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestRockPaperScissors_Answer(t *testing.T) {
	part1, part2, err := day.RockPaperScissors{}.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 11767, part1)
	assert.Equal(t, 13886, part2)
}

func TestRockPaperScissors_Part1(t *testing.T) {
	rps := day.RockPaperScissors{}

	game1 := []string{"A Y", "B X", "C Z"}
	assert.Equal(t, 15, rps.Part1(game1))
}

func TestRockPaperScissors_Part2(t *testing.T) {
	rps := day.RockPaperScissors{}

	game1 := []string{"A Y", "B X", "C Z"}
	assert.Equal(t, 12, rps.Part2(game1))
}

func TestRockPaperScissors_ScoreGame(t *testing.T) {
	rps := day.RockPaperScissors{}
	assert.Equal(t, 8, rps.ScoreGame("A Y"), "A -> Y")

	assert.Equal(t, 3, rps.ScoreGame("A Z"), "A -> Z (lose)")

	assert.Equal(t, 1, rps.ScoreGame("B X"), "B -> X (lose)")
	assert.Equal(t, 9, rps.ScoreGame("B Z"), "B -> X (win)")

	assert.Equal(t, 4, rps.ScoreGame("A X"), "A -> X (draw)")
	assert.Equal(t, 5, rps.ScoreGame("B Y"), "B -> Y (draw)")
	assert.Equal(t, 6, rps.ScoreGame("C Z"), "C -> Z (draw)")
}

func TestRockPaperScissors_MapCode(t *testing.T) {
	rps := day.RockPaperScissors{}
	// Rock
	assert.Equal(t, 1, rps.MapCode("A"), "A")
	assert.Equal(t, 1, rps.MapCode("X"), "X")
	// Paper
	assert.Equal(t, 2, rps.MapCode("B"), "B")
	assert.Equal(t, 2, rps.MapCode("Y"), "Y")
	// Scissors
	assert.Equal(t, 3, rps.MapCode("C"), "C")
	assert.Equal(t, 3, rps.MapCode("Z"), "Y")
}

func TestRockPaperScissors_ChooseHand(t *testing.T) {
	rps := day.RockPaperScissors{}

	assert.Equal(t, 3, rps.ChooseHand("A X"), "A -> lose") // lose
	assert.Equal(t, 1, rps.ChooseHand("A Y"), "A -> draw") // draw
	assert.Equal(t, 2, rps.ChooseHand("A Z"), "A -> win")  // win

	assert.Equal(t, 1, rps.ChooseHand("B X"), "B -> lose") // lose
	assert.Equal(t, 2, rps.ChooseHand("B Y"), "B -> draw") // draw
	assert.Equal(t, 3, rps.ChooseHand("B Z"), "B -> win")  // win

	assert.Equal(t, 2, rps.ChooseHand("C X"), "B -> lose") // lose
	assert.Equal(t, 3, rps.ChooseHand("C Y"), "B -> draw") // draw
	assert.Equal(t, 1, rps.ChooseHand("C Z"), "B -> win")  // win
}
