package passport

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// DataFile defines where to read input data from
var DataFile = "data/passport.txt"

type passport map[string]string

// Answer provides the day's answers
func Answer() (int, int, error) {
	f, err := os.Open(DataFile)
	if err != nil {
		return 0, 0, nil
	}
	defer f.Close()

	raw := ""
	numValid1 := 0
	numValid2 := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		line := strings.TrimRight(scanner.Text(), "\n")
		if line == "" {
			numValid1 += passportValid(parsePassport(raw))
			numValid2 += passportExtraValid(parsePassport(raw))
			raw = ""
			continue
		}
		raw += " " + line

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Process last passport
	numValid1 += passportValid(parsePassport(raw))
	numValid2 += passportExtraValid(parsePassport(raw))

	return numValid1, numValid2, nil
}

func parsePassport(raw string) passport {
	p := make(passport)
	for _, item := range strings.Split(strings.TrimLeft(raw, " "), " ") {
		i := strings.Index(item, ":")
		p[item[:i]] = item[i+1:]
	}
	return p
}

func passportValid(p passport) int {
	var keys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	missing := 0
	for _, key := range keys {
		if _, ok := p[key]; !ok {
			if key != "cid" {
				missing++
			}
		}
	}
	if missing > 0 {
		return 0
	}
	return 1
}

var reHcl = regexp.MustCompile(`#[0-9a-f]{6}`)
var reEcl = regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
var reHgt = regexp.MustCompile(`(\d+)(cm|in)`)

func pad(x string) string {
	return fmt.Sprintf("%4s", x)
}

func passportExtraValid(p passport) int {
	if passportValid(p) == 0 {
		// If not basic valid, then
		// not worth checking further.
		return 0
	}

	// If we're already basic valid we can
	// assume that all the keys are present
	// (apart from cid as it's optional)

	// Strings of the same length can be alphabetically-numerically compared
	byr := p["byr"]
	if len(byr) != 4 || byr < "1920" || byr > "2002" {
		return 0
	}

	iyr := p["iyr"]
	if len(iyr) != 4 || iyr < "2010" || iyr > "2020" {
		return 0
	}

	eyr := p["eyr"]
	if len(byr) != 4 || eyr < "2020" || eyr > "2030" {
		return 0
	}

	if !reHcl.MatchString(p["hcl"]) {
		return 0
	}

	if !reEcl.MatchString(p["ecl"]) {
		return 0
	}

	if len(p["pid"]) != 9 {
		return 0
	}

	matches := reHgt.FindStringSubmatch(p["hgt"])
	if len(matches) != 3 {
		return 0
	}
	h := matches[1]

	switch matches[2] {
	case "cm":
		if pad(h) < pad("150") || pad(h) > pad("193") {
			return 0
		}
	case "in":
		if pad(h) < pad("59") || pad(h) > pad("76") {
			return 0
		}
	}

	return 1
}
