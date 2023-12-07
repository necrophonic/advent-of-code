package day_test

import (
	"os"
	"testing"

	"github.com/necrophonic/advent-of-code/2021/go/pkg/advent"
)

func TestMain(m *testing.M) {
	advent.DataFolder = "../../data"
	os.Exit(m.Run())
}
