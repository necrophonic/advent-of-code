package main

import (
	"fmt"
	"log"

	"github.com/necrophonic/advent-of-code/2020/pkg/adaptor"
	"github.com/necrophonic/advent-of-code/2020/pkg/boarding"
	"github.com/necrophonic/advent-of-code/2020/pkg/customs"
	"github.com/necrophonic/advent-of-code/2020/pkg/encoding"
	"github.com/necrophonic/advent-of-code/2020/pkg/expenses"
	"github.com/necrophonic/advent-of-code/2020/pkg/game"
	"github.com/necrophonic/advent-of-code/2020/pkg/luggage"
	"github.com/necrophonic/advent-of-code/2020/pkg/passport"
	"github.com/necrophonic/advent-of-code/2020/pkg/password"
	"github.com/necrophonic/advent-of-code/2020/pkg/seating"
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
	answer(8, game.Answer)
	answer(9, encoding.Answer)
	answer(10, adaptor.Answer)
	answer(11, seating.Answer)
	answer(12, encoding.Answer)
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
