package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MohabMohamed/challenge/a"
	"github.com/MohabMohamed/challenge/b"
	"github.com/MohabMohamed/challenge/c"
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
	case "c":
		digitCount, err := c.GetInput()
		if err != nil {
			panic(err)
		}
		fmt.Printf("output = %s\n", c.GetPalindromeNumber(digitCount))
	default:
		panic("invalid challenge number")
	}
}
