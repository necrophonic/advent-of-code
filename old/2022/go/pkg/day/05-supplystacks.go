package day

import (
	"regexp"
	"strings"

	"github.com/necrophonic/advent-of-code/pkg/advent"
	"github.com/necrophonic/advent-of-code/pkg/debug"
)

var reStackBlock = regexp.MustCompile(`(.{3}) ?`)

type SupplyStacks struct {
	Stacks       [][]string
	backup       [][]string
	Instructions []string
}

func (SupplyStacks) Name() string {
	return "Supply Stacks"
}

func (s SupplyStacks) Answer() (any, any, error) {
	data, err := advent.LoadData("supplystacks.txt")
	if err != nil {
		return -1, -1, err
	}

	// FIXME!!!
	s.ConstructInstructions(data)

	return s.Part1(), -1, nil
}

func (s *SupplyStacks) Visualise() string {
	rows := []string{}

	outputWritten := true

	row := ""

	for ri := 0; outputWritten; ri++ {
		outputWritten = false
		for _, stack := range s.Stacks {
			if len(stack) > ri {
				row += stack[ri] + " "
				outputWritten = true
				continue
			}
			row += "    "
		}
		rows = append(rows, row)
		row = ""
	}

	out := ""
	for ri := len(rows) - 1; ri >= 0; ri-- {
		out += strings.TrimRight(rows[ri], " ") + "\n"
	}
	return out
}

func (s *SupplyStacks) Part1() string {
	s.Run()
	return s.TopStack()
}

func (s SupplyStacks) TopStack() string {
	top := ""
	for _, stack := range s.Stacks {
		top += s.UnpackBox(stack[0])
	}
	return top
}

// Run will run the instructions on the stacks
func (s *SupplyStacks) Run() []string {

	// Back up before we start so we can run again
	s.backup = make([][]string, len(s.Stacks))
	copy(s.backup, s.Stacks)

	for _, ins := range s.Instructions {
		parsed := s.ParseInstruction(ins)
		amount := parsed[0]
		from := parsed[1]
		to := parsed[2]

		debug.Print(s.Visualise())
		debug.Print("Move %d from %d to %d", amount, from+1, to+1)

		// Pull from the source stack
		move := s.Stacks[from][amount:]
		// Attach to the target stack
		s.Stacks[to] = append(s.Stacks[to], move...)
		// Prune the old stack
		s.Stacks[from] = s.Stacks[from][:amount]
	}

	return nil
}

// Reset resets the stacks to starting states
func (s *SupplyStacks) Reset() {
	copy(s.Stacks, s.backup)
}

func (SupplyStacks) UnpackBox(box string) string {
	return strings.ReplaceAll(strings.ReplaceAll(box, "]", ""), "[", "")
}

func (s *SupplyStacks) ConstructInstructions(data []string) {
	for _, ins := range data {

		if strings.HasPrefix(ins, "move") {
			s.Instructions = append(s.Instructions, ins)
			continue
		}

		stacks := reStackBlock.FindAllStringSubmatch(ins, -1)

		// Instantiate the stacks if not already
		if s.Stacks == nil {
			s.Stacks = make([][]string, len(stacks))
		}

		// Build the stacks.
		for i, stack := range stacks {
			if !strings.HasPrefix(stack[1], " ") {
				s.Stacks[i] = append([]string{stack[1]}, s.Stacks[i]...)
			}
		}
	}
}

func (SupplyStacks) ParseInstruction(ins string) []int {
	parts := strings.Split(ins, " ")
	stack := advent.StringToInt(parts[1])
	// minus one as we zero index the stacks
	from := advent.StringToInt(parts[3]) - 1
	to := advent.StringToInt(parts[5]) - 1
	return []int{stack, from, to}
}
