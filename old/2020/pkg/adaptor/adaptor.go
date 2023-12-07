package adaptor

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// DataFile defines where to read input data from
var DataFile = "data/adaptor.txt"

// Answer provides the day's answers
func Answer() (p1, p2 int, err error) {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		return 0, 0, err
	}

	input := strings.Split(string(data), "\n")

	adaptors := make([]int, 0, len(input))
	for _, adaptor := range input {
		if adaptor == "" {
			continue
		}
		d, err := strconv.Atoi(adaptor)
		if err != nil {
			return 0, 0, err
		}
		adaptors = append(adaptors, d)
	}

	adaptors1 := make([]int, len(adaptors))
	copy(adaptors1, adaptors)

	// adaptors2 := make([]int, len(adaptors))
	// copy(adaptors2, adaptors)

	p1 = JoltDifferences(adaptors1)
	// p2 = Combinations(adaptors2)

	return p1, p2, err
}

func Combinations(adaptors []int) int {
	sort.Ints(adaptors)

	combinations := 1

	// Insert 0 at start and max + 3 at the end
	adaptors = append([]int{0}, adaptors...)
	adaptors = append(adaptors, adaptors[len(adaptors)-1]+3)

	fmt.Println(adaptors)

	for a := 0; a < len(adaptors); a += 4 {
		base := adaptors[a]
		candidates := adaptors[a+1 : a+4]
		fmt.Printf("%d - %v\n", base, candidates)
		in := 0
		// i := 0
		// v := 0
		// c := 0
		for _, v := range candidates {
			// fmt.Printf("  -> %d (%v)\n", v, v-base <= 3)
			if v-base <= 3 {
				in++
				// c = i
			}
		}
		fmt.Printf("  * %d\n", in)
		combinations *= in
		// a += i
		// a += c + 1
	}

	return combinations
}

// func Combinations(adaptors []int) int {
// 	sort.Ints(adaptors)

// 	combinations := 0

// 	// Insert 0 at start and max + 3 at the end
// 	adaptors = append([]int{0}, adaptors...)
// 	adaptors = append(adaptors, adaptors[len(adaptors)-1]+3)

// 	// fmt.Println(adaptors)

// 	for a := 0; a < len(adaptors)-3; a++ {
// 		base := adaptors[a]
// 		candidates := adaptors[a+1 : a+4]
// 		// fmt.Printf("%d - %v\n", base, candidates)
// 		// in := 0

// 		for i, v := range candidates {
// 			// fmt.Println(adaptors[a+i+1:])
// 			if v-base <= 3 {
// 				fmt.Println("Check ", adaptors[a+i+1:])
// 				// combinations += Combinations(adaptors[a+i+1:])
// 			}
// 		}
// 		// fmt.Printf("  * %d\n", in)
// 		// combinations *= in
// 		// a += i
// 	}

// 	return combinations
// 	// return 1
// }

// JoltDifferences returns the differences
// between the set of adaptors.
func JoltDifferences(adaptors []int) int {
	sort.Ints(adaptors)
	jolt1 := 0
	jolt3 := 0

	// Insert 0 at start and max + 3 at the end
	adaptors = append([]int{0}, adaptors...)
	adaptors = append(adaptors, adaptors[len(adaptors)-1]+3)

	for a := 1; a < len(adaptors); a++ {
		diff := adaptors[a] - adaptors[a-1]
		switch diff {
		case 1:
			jolt1++
		case 3:
			jolt3++
		}
	}
	return jolt1 * jolt3
}
