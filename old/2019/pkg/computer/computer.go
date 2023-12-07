package computer

import (
	"bytes"
	"io/ioutil"
	"strconv"

	"github.com/pkg/errors"
)

// DataFile defines where to read input data from
var DataFile = "data/computer.txt"

// Answer provides the day's answers
func Answer() (int, int, error) {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		return 0, 0, err
	}
	data = bytes.TrimRight(data, "\n")

	opcodes := bytes.Split(data, []byte(","))
	program := make([]int, len(opcodes))
	for i, opcode := range opcodes {
		code, err := strconv.Atoi(string(opcode))
		if err != nil {
			return 0, 0, errors.WithMessage(err, "failed to parse opcode")
		}
		program[i] = code
	}

	output, err := Run(SetInput(program, 12, 2))
	if err != nil {
		return 0, 0, errors.WithMessage(err, "failed to run program")
	}

	noun, verb, err := FindInputValues(program, 19690720)
	if err != nil {
		return 0, 0, errors.WithMessage(err, "failed to find noun and verb")
	}

	return output, 100*noun + verb, nil
}

// FindInputValues attempts to determine the noun
// and verb to calculate a given target.
func FindInputValues(program []int, target int) (int, int, error) {
	for verb := 0; verb <= 99; verb++ {
		for noun := 0; noun <= 99; noun++ {
			output, err := Run(SetInput(program, noun, verb))
			if err != nil {
				return 0, 0, err
			}
			if output == target {
				return noun, verb, nil
			}
		}
	}
	return 0, 0, errors.New("could not calculate target")
}

// EOP is an End Of Program error
type EOP struct{}

// Error satisfies the error interface
func (e EOP) Error() string {
	return "EndOfProgram"
}

// SetInput sets the input positions
// of a program
func SetInput(program []int, noun, verb int) []int {
	p := make([]int, len(program))
	copy(p, program)
	p[1] = noun
	p[2] = verb
	return p
}

// Run runs a program
func Run(program []int) (int, error) {
	position := 0
	var err error
	for {
		position, err = ProcessInstruction(program, position)
		if err != nil {
			if !errors.Is(err, EOP{}) {
				return 0, err
			}
			break
		}
	}
	return program[0], nil
}

// ProcessInstruction processes an opcode instruction
func ProcessInstruction(pr []int, p int) (int, error) {
	switch pr[p] {
	case 1:
		pr[pr[p+3]] = pr[pr[p+1]] + pr[pr[p+2]]
	case 2:
		pr[pr[p+3]] = pr[pr[p+1]] * pr[pr[p+2]]
	case 99:
		return -1, EOP{}
	}
	return p + 4, nil
}
