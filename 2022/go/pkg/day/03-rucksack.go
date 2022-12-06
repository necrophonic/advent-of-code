package day

import (
	"github.com/necrophonic/advent-of-code/pkg/advent"
	"github.com/necrophonic/advent-of-code/pkg/debug"
)

type Rucksack struct{}

func (r Rucksack) Name() string {
	return "Rucksack Reorganization"
}

func (r Rucksack) Answer() (int, int, error) {
	data, err := advent.LoadData("rucksack.txt")
	if err != nil {
		return -1, -1, err
	}

	return r.Part1(data), r.Part2(data), nil
}

func (r Rucksack) Part1(rucksacks []string) int {
	sum := 0
	for _, rucksack := range rucksacks {
		sum += r.ScoreDuplicates(r.FindDuplicates(r.SplitCompartments(rucksack)))
	}
	return sum
}

func (r Rucksack) Part2(rucksacks []string) int {
	sum := 0
	for i := 0; i < len(rucksacks)-2; i += 3 {
		group := rucksacks[i : i+3]
		sum += r.ScoreDuplicates(r.FindDuplicates(group...))
		// 	debug.Print("Group: %v", group)
		// 	groupDupes := make([]string, 3)
		// 	for j := 0; j < 3; j++ {
		// 		groupDupes[j] = r.FindDuplicates(r.SplitCompartments(group[j]))
		// 	}
	}
	return sum
}

func (r Rucksack) SplitCompartments(items string) (string, string) {
	half := len(items) / 2
	return items[:half], items[half:]
}

func (r Rucksack) FindDuplicates(rucksacks ...string) string {

	deduped := ""

	// Loop through each rucksack apart from the last one
	for ri := 0; ri < len(rucksacks)-1; ri++ {

		r1 := r.DedupeCompartment(rucksacks[ri])
		r2 := r.DedupeCompartment(rucksacks[ri+1])

		// Reset
		deduped = ""

		for _, i1 := range r1 {
			for _, i2 := range r2 {
				if i1 == i2 {
					deduped += string(i1)
				}
			}
		}
		debug.Print("Find rucksack dupes (%s, %s) -> %s", r1, r2, deduped)

		rucksacks[ri+1] = deduped
	}
	return deduped
}

// DededupeCompartment removes duplicates within a single compartment.
// It doesn't attempt to sort the output.
func (r Rucksack) DedupeCompartment(compartment string) string {
	dupeMap := map[rune]bool{}
	deduped := ""
	for _, item := range compartment {
		if !dupeMap[item] {
			deduped += string(item)
		}
		dupeMap[item] = true
	}
	return deduped
}

func (r Rucksack) ScoreDuplicates(dupeList string) int {
	score := 0
	for _, dupe := range dupeList {
		score += r.ScoreDupe(dupe)
	}
	return score
}

func (r Rucksack) ScoreDupe(dupe rune) int {
	return (int(dupe) - 38) % 58
}
