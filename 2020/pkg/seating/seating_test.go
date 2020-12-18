package seating_test

import (
	"fmt"
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/seating"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	seating.DataFile = "../../data/seating.txt"
	p1, p2, err := seating.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 2166, p1, "Part 1")
	assert.Equal(t, 1955, p2, "Part 2")
}

func TestGetXY(t *testing.T) {
	g := seating.NewGrid([]byte("ABC\nDEF\nGHI"))

	assert.Equal(t, byte('B'), g.GetXY(1, 0))
	assert.Equal(t, byte('F'), g.GetXY(2, 1))
	assert.Equal(t, byte('H'), g.GetXY(1, 2))
}

func TestRound(t *testing.T) {
	seats := []byte("L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL")
	rounds := []string{
		"#.##.##.## #######.## #.#.#..#.. ####.##.## #.##.##.## #.#####.## ..#.#..... ########## #.######.# #.#####.## ",
		"#.LL.L#.## #LLLLLL.L# L.L.L..L.. #LLL.LL.L# #.LL.LL.LL #.LLLL#.## ..L.L..... #LLLLLLLL# #.LLLLLL.L #.#LLLL.## ",
		"#.##.L#.## #L###LL.L# L.#.#..#.. #L##.##.L# #.##.LL.LL #.###L#.## ..#.#..... #L######L# #.LL###L.L #.#L###.## ",
		"#.#L.L#.## #LLL#LL.L# L.L.L..#.. #LLL.##.L# #.LL.LL.LL #.LL#L#.## ..L.L..... #L#LLLL#L# #.LLLLLL.L #.#L#L#.## ",
		"#.#L.L#.## #LLL#LL.L# L.#.L..#.. #L##.##.L# #.#L.LL.LL #.#L#L#.## ..L.L..... #L#L##L#L# #.LLLLLL.L #.#L#L#.## ",
	}

	g := seating.NewGrid(seats)
	c := 0

	for _, plan := range rounds {
		c = g.Round()
		assert.Equal(t, plan, g.GetString())
	}
	c = g.Round()
	assert.Equal(t, 0, c)
	assert.Equal(t, 37, g.NumberOccupied, "Should have 37 seats occupied")
}

func TestNeighbours(t *testing.T) {
	tcs := []struct {
		x, y, expected int
		grid           *seating.Grid
	}{
		{0, 0, 0, seating.NewGrid([]byte("L"))},
		{1, 0, 2, seating.NewGrid([]byte("#L#"))},
		{1, 1, 3, seating.NewGrid([]byte("....#\n#L.#.\n.##.."))},
		{4, 2, 1, seating.NewGrid([]byte("....#\n#L.#.\n.##.."))},
	}

	for _, test := range tcs {
		assert.Equal(t, test.expected, test.grid.Neighbours(test.x, test.y))
	}
}

func TestPickyRound(t *testing.T) {
	seats := []byte("L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL")
	rounds := []string{
		"#.##.##.## #######.## #.#.#..#.. ####.##.## #.##.##.## #.#####.## ..#.#..... ########## #.######.# #.#####.## ",
		"#.LL.LL.L# #LLLLLL.LL L.L.L..L.. LLLL.LL.LL L.LL.LL.LL L.LLLLL.LL ..L.L..... LLLLLLLLL# #.LLLLLL.L #.LLLLL.L# ",
		"#.L#.##.L# #L#####.LL L.#.#..#.. ##L#.##.## #.##.#L.## #.#####.#L ..#.#..... LLL####LL# #.L#####.L #.L####.L# ",
		"#.L#.L#.L# #LLLLLL.LL L.L.L..#.. ##LL.LL.L# L.LL.LL.L# #.LLLLL.LL ..L.L..... LLLLLLLLL# #.LLLLL#.L #.L#LL#.L# ",
		"#.L#.L#.L# #LLLLLL.LL L.L.L..#.. ##L#.#L.L# L.L#.#L.L# #.L####.LL ..#.#..... LLL###LLL# #.LLLLL#.L #.L#LL#.L# ",
		"#.L#.L#.L# #LLLLLL.LL L.L.L..#.. ##L#.#L.L# L.L#.LL.L# #.LLLL#.LL ..#.L..... LLL###LLL# #.LLLLL#.L #.L#LL#.L# ",
	}

	g := seating.NewGrid(seats)
	c := 0

	for _, plan := range rounds {
		c = g.PickyRound()
		assert.Equal(t, plan, g.GetString())
	}
	c = g.Round()
	assert.Equal(t, 0, c)
	assert.Equal(t, 26, g.NumberOccupied, "Should have 26 seats occupied")
}

func TestPickyNeighbours(t *testing.T) {
	tcs := []struct {
		x, y, expected int
		grid           *seating.Grid
	}{
		// First case should be 8, but we cap the check at 5
		{3, 4, 5, seating.NewGrid([]byte(".......#.\n...#.....\n.#.......\n.........\n..#L....#\n....#....\n.........\n#........\n...#....."))},
		{1, 1, 0, seating.NewGrid([]byte(".............\n.L.L.#.#.#.#.\n............."))},
		{3, 3, 0, seating.NewGrid([]byte(".##.##.\n#.#.#.#\n##...##\n...L...\n##...##\n#.#.#.#\n.##.##."))},
	}

	for _, test := range tcs {
		fmt.Println(test.grid.GetXY(1, 0))
		assert.Equal(t, test.expected, test.grid.PickyNeighbours(test.x, test.y))
	}
}
