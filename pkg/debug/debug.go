package debug

import "fmt"

var On = false

func Print(f string, args ...any) {
	if On {
		fmt.Printf(f+"\n", args...)
	}
}
