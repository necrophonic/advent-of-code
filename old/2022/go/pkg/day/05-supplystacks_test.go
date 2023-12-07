package day_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/necrophonic/advent-of-code/2022/go/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestSupplyStacks_Answer(t *testing.T) {
	part1, part2, err := day.SupplyStacks{}.Answer()
	assert.NoError(t, err)
	assert.Equal(t, -999, part1)
	assert.Equal(t, -999, part2)
}

func TestSupplyStacks_Part1(t *testing.T) {
	ss := day.SupplyStacks{}
	ss.ConstructInstructions([]string{
		"[D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	})
	ss.Run()
	fmt.Println(ss.Visualise())
	assert.True(t, false)
}

func TestSupplyStacks_ParseInstruction(t *testing.T) {
	ss := day.SupplyStacks{}
	assert.Equal(t, []int{1, 1, 0}, ss.ParseInstruction("move 1 from 2 to 1"))
	assert.Equal(t, []int{4, 2, 1}, ss.ParseInstruction("move 4 from 3 to 2"))
}

func TestSupplyStacks_UnpackBox(t *testing.T) {
	ss := day.SupplyStacks{}
	assert.Equal(t, "M", ss.UnpackBox("[M]"))
	assert.Equal(t, "Z", ss.UnpackBox("[Z]"))
}

func TestSupplyStacks_Visualise(t *testing.T) {
	data := []string{

		"                        [R] [J] [W]",
		"            [R] [N]     [T] [T] [C]",
		"[R]         [P] [G]     [J] [P] [T]",
		"[Q]     [C] [M] [V]     [F] [F] [H]",
		"[G] [P] [M] [S] [Z]     [Z] [C] [Q]",
		"[P] [C] [P] [Q] [J] [J] [P] [H] [Z]",
		"[C] [T] [H] [T] [H] [P] [G] [L] [V]",
		"[F] [W] [B] [L] [P] [D] [L] [N] [G]",
	}
	s := day.SupplyStacks{}
	s.ConstructInstructions(data)

	assert.Equal(t, "\n"+strings.Join(data, "\n")+"\n", s.Visualise())

	// fmt.Println(s.Visualise())
}
