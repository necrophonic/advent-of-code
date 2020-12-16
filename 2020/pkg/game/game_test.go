package game_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/game"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	game.DataFile = "../../data/game.txt"
	p1, p2, err := game.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 2003, p1, "Part 1")
	assert.Equal(t, 1984, p2, "Part 2")
}

func TestProcess(t *testing.T) {
	program := [][]byte{
		[]byte("nop +0"),
		[]byte("acc +1"),
		[]byte("jmp +4"),
		[]byte("acc +3"),
		[]byte("jmp -3"),
		[]byte("acc -99"),
		[]byte("acc +1"),
		[]byte("jmp -4"),
		[]byte("acc +6"),
	}
	result, err := game.Process(program)
	assert.EqualError(t, err, "Already seen")
	assert.Equal(t, 5, result)
}

func TestProcessFix(t *testing.T) {
	program := [][]byte{
		[]byte("nop +0"),
		[]byte("acc +1"),
		[]byte("jmp +4"),
		[]byte("acc +3"),
		[]byte("jmp -3"),
		[]byte("acc -99"),
		[]byte("acc +1"),
		[]byte("jmp -4"),
		[]byte("acc +6"),
	}
	result, err := game.ProcessFix(program)
	assert.NoError(t, err)
	assert.Equal(t, 8, result)
}
