package game

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

// DataFile defines where to read input data from
var DataFile = "data/game.txt"

// Answer provides the day's answers
func Answer() (int, int, error) {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		return 0, 0, err
	}

	instructions := bytes.Split(data, []byte("\n"))
	result, err := Process(instructions)
	if err != nil {
		if !errors.Is(err, ErrSeen) {
			return 0, 0, err
		}
	}

	return result, -1, nil
}

var ErrSeen = errors.New("Already seen")

func Process(instructions [][]byte) (int, error) {

	accumulator := 0
	seen := make(map[int]bool)

	for i := 0; i < len(instructions); i++ {
		op, sign, value, err := parseInstruction(instructions[i])
		if err != nil {
			return 0, errors.WithMessage(err, "failed to parse instructions")
		}
		// fmt.Printf(" - %s (%s%d)\n", op, sign, value)

		if seen[i] {
			return accumulator, ErrSeen
		}
		seen[i] = true

		switch op {
		case "nop":
			continue
		case "acc":
			if sign == "+" {
				accumulator += value
				continue
			}
			accumulator -= value
		case "jmp":
			if sign == "+" {
				i += value - 1
				continue
			}
			i -= value + 1
		}
	}
	return accumulator, nil
}

var reInstruction = regexp.MustCompile(`(nop|acc|jmp) ([+-])(\d+)`)

func parseInstruction(instruction []byte) (op string, sign string, value int, err error) {
	m := reInstruction.FindStringSubmatch(string(instruction))
	op = m[1]
	sign = m[2]
	value, err = strconv.Atoi(m[3])
	if err != nil {
		return "", "", 0, err
	}
	return
}
