package day

import (
	"fmt"
	"strings"

	"github.com/necrophonic/advent-of-code/pkg/advent"
	"github.com/necrophonic/advent-of-code/pkg/debug"
)

type RockPaperScissors struct{}

const (
	rock = iota + 1
	paper
	scissors
)

const (
	rpsWin  = 6
	rpsDraw = 3
	rpsLose = 0
)

func (c RockPaperScissors) Name() string {
	return "Rock Paper Scissors"
}

func (rps RockPaperScissors) Answer() (int, int, error) {
	data, err := advent.LoadData("rockpaper.txt")
	if err != nil {
		return -1, -1, err
	}

	return rps.Part1(data), rps.Part2(data), nil
}

func (rps RockPaperScissors) Part1(games []string) int {
	sum := 0
	for _, game := range games {
		score := rps.ScoreGame(game)
		sum += score
	}
	return sum
}

func (rps RockPaperScissors) Part2(games []string) int {
	sum := 0
	for _, game := range games {
		// Get the desired hand to play
		handInt := rps.ChooseHand(game)

		elf, _ := rps.SplitGame(game)
		// Reconstruct and score the game
		hand := ""
		switch handInt {
		case 1:
			hand = "X"
		case 2:
			hand = "Y"
		case 3:
			hand = "Z"
		}
		debug.Print("Chosen hand: %s", hand)
		score := rps.ScoreGame(rps.MakeGame(elf, hand))
		debug.Print("Game %s -> %d\n", game, score)
		sum += score
	}
	return sum
}

func (RockPaperScissors) SplitGame(game string) (string, string) {
	parts := strings.Split(game, " ")
	return parts[0], parts[1]
}

func (RockPaperScissors) MakeGame(elf, player string) string {
	return fmt.Sprintf("%s %s", elf, player)
}

func (rps RockPaperScissors) ChooseHand(game string) int {
	debug.Print("Choose hand: %s", game)
	elf, target := rps.SplitGame(game)
	elfZeroIndexed := rps.MapCode(elf) - 1

	switch target {
	case "X": // lose
		return (elfZeroIndexed+2)%3 + 1
	case "Y": // draw
		return elfZeroIndexed + 1
	case "Z": // win
		return (elfZeroIndexed+1)%3 + 1
	}
	return -1
}

func (rps RockPaperScissors) GetWin(elf string) string {
	switch elf {
	case "A":
		return "Y" // paper
	case "B":
		return "Z" // scissors
	default:
		return "X" // rock
	}
}

func (rps RockPaperScissors) ScoreGame(game string) int {
	debug.Print("Scoring game: %s", game)
	players := strings.Split(game, " ")

	elfDecode := rps.MapCode(players[0])
	youDecode := rps.MapCode(players[1])

	result := youDecode - elfDecode

	switch result {
	case 2, -1:
		return youDecode + rpsLose
	case 1, -2:
		return youDecode + rpsWin
	default:
		return youDecode + rpsDraw
	}
}

func (RockPaperScissors) MapCode(code string) int {
	switch strings.ToUpper(code) {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	}
	return -1
}
