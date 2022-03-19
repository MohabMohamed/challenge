package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MohabMohamed/challenge/a"
	"github.com/MohabMohamed/challenge/b"
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
	case "b":
		str1, str2 := b.GetInput()
		fmt.Printf("output = %v\n", b.IsAnagram(str1, str2))
	default:
		panic("invalid challenge number")
	}
}
