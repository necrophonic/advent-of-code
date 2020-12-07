package main

import (
	"fmt"
	"log"

	"github.com/necrophonic/advent-of-code/2019/pkg/computer"
	"github.com/necrophonic/advent-of-code/2019/pkg/fuel"
)

func main() {
	fmt.Println("Answers:")
	answer(1, fuel.Answer)
	answer(2, computer.Answer)
	noAnswer(3)
	noAnswer(4)
	noAnswer(5)
	noAnswer(6)
	noAnswer(7)
	noAnswer(8)
	noAnswer(9)
	noAnswer(10)
	noAnswer(11)
	noAnswer(12)
	noAnswer(13)
	noAnswer(14)
	noAnswer(15)
	noAnswer(16)
	noAnswer(17)
	noAnswer(18)
	noAnswer(19)
	noAnswer(20)
	noAnswer(21)
	noAnswer(22)
	noAnswer(23)
	noAnswer(24)
}

func answer(day int, f func() (int, int, error)) {
	p1, p2, err := f()
	if err != nil {
		log.Fatalf("error when answering day %d: %v", day, err)
	}
	fmt.Printf("- [day %2d] part 1: %-12d part 2: %-12d\n", day, p1, p2)
}

func noAnswer(day int) {
	fmt.Printf("- [day %2d] no answer yet\n", day)
}
