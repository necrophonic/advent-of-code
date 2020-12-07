package luggage_test

import (
	"testing"

	"github.com/necrophonic/advent-of-code/2020/pkg/luggage"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	luggage.DataFile = "../../data/luggage.txt"
	p1, p2, err := luggage.Answer()
	assert.NoError(t, err)
	assert.Equal(t, 296, p1, "Part 1")
	assert.Equal(t, 9339, p2, "Part 2")
}

func testBagCanContain(t *testing.T) {
	b := luggage.Bag{
		Name: "purple",
		Contains: map[string]int{
			"red":  1,
			"blue": 1,
		},
	}
	assert.True(t, b.CanContain("red"))
	assert.False(t, b.CanContain("pink"))
}

func TestParseRules(t *testing.T) {
	rules := [][]byte{
		[]byte("dotted silver bags contain 2 posh blue bags, 3 posh maroon bags."),
		[]byte("plaid salmon bags contain 4 light yellow bags."),
		[]byte("posh black bags contain 3 dark lavender bags, 3 mirrored coral bags, 1 dotted chartreuse bag."),
		[]byte("pale chartreuse bags contain no other bags."),
		[]byte("wavy magenta bags contain 2 vibrant crimson bags, 3 mirrored teal bags, 1 shiny lime bag."),
	}

	expected := luggage.BagMap{
		"dotted silver":   {"dotted silver", map[string]int{"posh blue": 2, "posh maroon": 3}},
		"plaid salmon":    {"plaid salmon", map[string]int{"light yellow": 4}},
		"posh black":      {"posh black", map[string]int{"dark lavender": 3, "mirrored coral": 3, "dotted chartreuse": 1}},
		"pale chartreuse": {"pale chartreuse", map[string]int{}},
		"wavy magenta":    {"wavy magenta", map[string]int{"vibrant crimson": 2, "mirrored teal": 3, "shiny lime": 1}},
	}

	bags := luggage.ParseRules(rules)
	assert.Equal(t, expected, bags)
}

func TestNumberBagsContained(t *testing.T) {
	rules := [][]byte{
		[]byte("shiny gold bags contain 2 dark red bags."),
		[]byte("dark red bags contain 2 dark orange bags."),
		[]byte("dark orange bags contain 2 dark yellow bags."),
		[]byte("dark yellow bags contain 2 dark green bags."),
		[]byte("dark green bags contain 2 dark blue bags."),
		[]byte("dark blue bags contain 2 dark violet bags."),
		[]byte("dark violet bags contain no other bags."),
	}

	bags := luggage.ParseRules(rules)
	count := luggage.NumberBagsContained("shiny gold", bags) - 1
	assert.Equal(t, 126, count)
}
