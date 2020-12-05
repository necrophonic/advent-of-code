package boarding

import (
	"bufio"
	"log"
	"os"
)

var (
	seats []int
	rows  [128]int
	cols  [8]int
)

// DataFile defines where to read input data from
var DataFile = "data/boarding.txt"

// Answer provides the day's answers
func Answer() (int, int, error) {
	f, err := os.Open(DataFile)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	// Populate seats, rows and cols
	seats = make([]int, 1024)
	for i := 0; i < 1024; i++ {
		seats[i] = i
		if i < 128 {
			rows[i] = i
		}
		if i < 8 {
			cols[i] = i
		}
	}

	highest := 0
	lowest := 1024

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		directions := scanner.Text()
		_, _, seat := FindSeat(directions)

		// Determine if highest
		if seat > highest {
			highest = seat
		}
		if seat < lowest {
			lowest = seat
		}

		// Set found seats to -1
		seats[seat] = -1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Find missing seat
	// Cut out the front and the back that can't be the seat we want
	missing := 0
	for _, missing = range seats[lowest:highest] {
		if missing != -1 {
			break
		}
	}

	return highest, missing, nil
}

// FindSeat will locate a seat
func FindSeat(directions string) (row, col, seat int) {
	// Slice of all possible rows. We can then reslice this
	// as we search through the plane to find our seat.
	sr := rows[:]
	sc := cols[:]

	for _, move := range directions[:7] {
		if move == 'F' {
			sr = sr[:len(sr)/2]
			continue
		}
		sr = sr[len(sr)/2:]
	}
	row = sr[0]

	for _, move := range directions[7:] {
		if move == 'L' {
			sc = sc[:len(sc)/2]
			continue
		}
		sc = sc[len(sc)/2:]
	}
	col = sc[0]
	return row, col, (row * 8) + col
}
