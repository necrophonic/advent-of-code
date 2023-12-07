package day

import (
	"bytes"
	"sort"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/advent"
)

type SyntaxScoring struct {
	Data []string
}

func (*SyntaxScoring) Name() string {
	return "Syntax Scoring"
}

func (s *SyntaxScoring) Answer() (int, int, error) {
	data, err := advent.LoadData("syntax-scoring.txt")
	if err != nil {
		return -1, -1, err
	}
	s.Data = data

	return s.Part1(), s.Part2(), nil
}

func (s *SyntaxScoring) Part1() int {
	score := 0
	for i := 0; i < len(s.Data); i++ {
		add := s.CheckLine(s.Data[i])
		if add != 0 {
			score += add
			s.Data[i] = ""
		}
	}
	return score
}

func (s *SyntaxScoring) Part2() int {
	scores := []int{}
	for _, line := range s.Data {
		if line == "" {
			continue // Line has been removed by part 1
		}
		scores = append(scores, s.FixLine(line))
	}
	sort.Ints(scores)
	return scores[(len(scores) / 2)]
}

func (*SyntaxScoring) FixLine(line string) int {
	score, stackCount := 0, 0

	for i := len(line) - 1; i >= 0; i-- {
		char := line[i]

		// If closing brace, increment stack count by 1
		if bytes.Contains([]byte(`))]}>`), []byte{char}) {
			stackCount++
			continue
		}

		// If opening brace, if stack >0, then dec stack and move on,
		// otherwise add a "fix" brace to the output
		if stackCount > 0 {
			stackCount--
			continue
		}

		switch char {
		case '(':
			score = (score * 5) + 1
		case '[':
			score = (score * 5) + 2
		case '{':
			score = (score * 5) + 3
		case '<':
			score = (score * 5) + 4
		}
	}

	return score
}

// CheckLine looks for a syntax error in the given line. If found it will return
// the score for the bad character, or 0 otherwise.
func (*SyntaxScoring) CheckLine(line string) int {
	stack := make([]byte, 0, len(line))
	stack = append(stack, line[0])

	for i := 1; i < len(line); i++ {
		char := line[i]
		if bytes.Contains([]byte(`)]}>`), []byte{char}) {
			// Found a closing brace - pop off the stack to compare
			var compare byte
			compare, stack = stack[len(stack)-1], stack[:len(stack)-1]
			switch char {
			case ')':
				if compare != '(' {
					return 3
				}
			case ']':
				if compare != '[' {
					return 57
				}
			case '}':
				if compare != '{' {
					return 1197
				}
			case '>':
				if compare != '<' {
					return 25137
				}
			}
			continue
		}
		stack = append(stack, char)
	}
	return 0
}
