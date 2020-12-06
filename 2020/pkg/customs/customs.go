package customs

import (
	"bytes"
	"io/ioutil"
)

// DataFile defines where to read input data from
var DataFile = "data/customs.txt"

// Answer provides the day's answers
func Answer() (int, int, error) {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		return 0, 0, err
	}
	data = bytes.TrimRight(data, "\n")

	answerCountPart1 := 0
	answerCountPart2 := 0

	for _, group := range bytes.Split(data, []byte("\n\n")) {
		answerCountPart1 += CountAnyAnswered(group)
		answerCountPart2 += CountAllAnswered(group)
	}

	return answerCountPart1, answerCountPart2, nil
}

// CountAnyAnswered returns the number of any
// question answered for the given group.
func CountAnyAnswered(group []byte) int {
	answers := make(map[byte]int)
	for _, person := range bytes.Split(group, []byte("\n")) {
		for _, answer := range person {
			answers[answer]++
		}
	}
	return len(answers)
}

// CountAllAnswered returns the number of questions
// answered by ALL in the group.
func CountAllAnswered(group []byte) (all int) {
	answers := make(map[byte]int)
	people := bytes.Split(group, []byte("\n"))
	for _, person := range people {
		for _, answer := range person {
			if answers[answer]++; answers[answer] == len(people) {
				all++
			}
		}
	}
	return all
}
