package main

import (
	"fmt"
	"log"

	"github.com/necrophonic/advent-of-code/2020/pkg/boarding"
	"github.com/necrophonic/advent-of-code/2020/pkg/customs"
	"github.com/necrophonic/advent-of-code/2020/pkg/expenses"
	"github.com/necrophonic/advent-of-code/2020/pkg/luggage"
	"github.com/necrophonic/advent-of-code/2020/pkg/passport"
	"github.com/necrophonic/advent-of-code/2020/pkg/password"
	"github.com/necrophonic/advent-of-code/2020/pkg/toboggan"
)

func main() {
	fmt.Println("Answers:")
	answer(1, expenses.Answer)
	answer(2, password.Answer)
	answer(3, toboggan.Answer)
	answer(4, passport.Answer)
	answer(5, boarding.Answer)
	answer(6, customs.Answer)
	answer(7, luggage.Answer)
}

func answer(day int, f func() (int, int, error)) {
	p1, p2, err := f()
	if err != nil {
		log.Fatalf("error when answering day %d: %v", day, err)
	}
	print(day, p1, p2)
}

func print(day, part1, part2 int) {
	fmt.Printf("- [day %2d] part 1: %-12d part 2: %-12d\n", day, part1, part2)
}
