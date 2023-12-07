package password

import (
	"strconv"
)

const min = 136760
const max = 595730

// Answer provides the day's answers
func Answer() (int, int, error) {

	match := 0
	matchExtended := 0
	for i := min; i < max; i++ {
		password := strconv.Itoa(i)
		match += MatchCriteria(password)
		matchExtended += MatchExtendedCriteria(password)
	}

	return match, matchExtended, nil
}

// MatchCriteria returns 1 if the
// password matches the criteria.
func MatchCriteria(password string) int {
	return basicMatch(password)
}

// MatchExtendedCriteria returns 1 if the password
// matches the extended matching criteria.
func MatchExtendedCriteria(password string) int {
	// fmt.Println(password)
	if basicMatch(password) == 1 {
		double := false
		lp := len(password)
		for d := 0; d < lp-1; d++ {
			if password[d] == password[d+1] {
				double = true
				d++
				if d+1 != lp && password[d] == password[d+1] {
					double = false
					// Read until end of this digit
					for ; d < lp; d++ {
						if password[d] != password[d-1] {
							d--
							break
						}
					}
				}
			}
			if double {
				return 1
			}
		}
	}
	return 0
}

func basicMatch(password string) int {
	if len(password) != 6 {
		return 0
	}
	// fmt.Println(password)
	double := false
	for d := 0; d < len(password)-1; d++ {
		if password[d] > password[d+1] {
			return 0
		}
		if password[d] == password[d+1] {
			double = true
		}
	}
	if double {
		return 1
	}
	return 0
}
