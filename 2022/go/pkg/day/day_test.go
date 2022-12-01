package day_test

import (
	"os"
	"testing"

	"github.com/necrophonic/advent-of-code/pkg/advent"
)

func TestMain(m *testing.M) {
	advent.DataFolder = "../../data"
	os.Exit(m.Run())
}
