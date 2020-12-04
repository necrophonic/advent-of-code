package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(\d+)-(\d+) (.): (.+)`)

func main() {

	f, err := os.Open("day-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	numValidPart1 := 0
	numValidPart2 := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mn, mx, l, p, err := parse(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numValidPart1 += passwordValidPart1(p, l, mn, mx)
		numValidPart2 += passwordValidPart2(p, l, mn, mx)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Answer to day 2 part 1 is %d\n", numValidPart1)
	fmt.Printf("Answer to day 2 part 2 is %d\n", numValidPart2)
}

func passwordValidPart1(password string, letter rune, min, max int) int {
	lc := 0
	if min > len(password) {
		return 0
	}
	for _, r := range password {
		if r == letter {
			if lc++; lc > max {
				return 0
			}
		}
	}
	if lc >= min {
		return 1
	}
	return 0
}

func passwordValidPart2(password string, letter rune, min, max int) int {
	rPassword := []rune(password)
	ret := 0
	if rPassword[min-1] == letter {
		ret++
	}
	if rPassword[max-1] == letter {
		ret++
	}
	return ret % 2

}

func parse(line string) (min, max int, letter rune, password string, err error) {
	found := re.FindStringSubmatch(line)
	if len(found) == 0 {
		return
	}

	if min, err = strconv.Atoi(found[1]); err != nil {
		return
	}

	if max, err = strconv.Atoi(found[2]); err != nil {
		return
	}

	letter = []rune(found[3])[0]
	password = found[4]
	return
}
