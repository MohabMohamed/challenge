package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MohabMohamed/challenge/a"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("you should provide a command-line argument for which problem to run.")
	}
	challengeNum := strings.ToLower(args[1])
	switch challengeNum {
	case "a":
		haystack, needle := a.GetInput()
		fmt.Printf("output = %d\n", a.FindNeedleIndex(haystack, needle))

	default:
		panic("invalid challenge number")
	}
}
