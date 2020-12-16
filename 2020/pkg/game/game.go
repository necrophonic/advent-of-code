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

	data = bytes.TrimRight(data, "\n")

	instructions := bytes.Split(data, []byte("\n"))
	result, err := Process(instructions)
	if err != nil {
		if !errors.Is(err, ErrSeen) {
			return 0, 0, err
		}
	}

	resultFixed, err := ProcessFix(instructions)
	if err != nil {
		return 0, 0, err
	}

	return result, resultFixed, nil
}

func flipInstruction(instructions Instructions, index int) {
	if bytes.HasPrefix(instructions[index], []byte("jmp")) {
		instructions[index] = bytes.Replace(instructions[index], []byte("jmp"), []byte("nop"), 1)
		return
	}
	instructions[index] = bytes.Replace(instructions[index], []byte("nop"), []byte("jmp"), 1)
}

// Instructions is an executable program
type Instructions [][]byte

// Custom errors for exit states
var (
	ErrSeen = errors.New("Already seen")
	ErrEOP  = errors.New("End of program")
)

// ProcessFix performs a rolling fix-and-process until the program
// executes until a proper End Of Program.
func ProcessFix(instructions Instructions) (result int, err error) {
	flipped := 0
	flipInstruction(instructions, flipped)
	for {
		result, err = Process(instructions)
		if err != nil {
			if errors.Is(err, ErrEOP) {
				return result, nil
			}
			if !errors.Is(err, ErrSeen) {
				// Genuine error
				return 0, err
			}
			// We saw an instruction twice, so we
			// attempt a fix and then re-run.
		}
		i := 0
		for i = flipped + 1; i < len(instructions); i++ {
			if bytes.HasPrefix(instructions[i], []byte("jmp")) || bytes.HasPrefix(instructions[i], []byte("nop")) {
				flipInstruction(instructions, i)
				break
			}
		}
		// Flip the old one back
		flipInstruction(instructions, flipped)
		flipped = i
	}
}

// Process runs an instruction set
func Process(instructions Instructions) (int, error) {

	accumulator := 0
	seen := make(map[int]bool)

	for i := 0; i < len(instructions); i++ {
		op, sign, value, err := parseInstruction(instructions[i])
		if err != nil {
			return 0, errors.WithMessage(err, "failed to parse instructions")
		}

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
	return accumulator, ErrEOP
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
