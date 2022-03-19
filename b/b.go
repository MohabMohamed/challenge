package b

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsAnagram(str1, str2 string) bool {
	str1, str2 = sanitizeInput(str1, str2)
	occurrences := make(map[rune]int)

	for _, letter := range str1 {
		occurrences[letter]++
	}

	for _, letter := range str2 {
		if count, found := occurrences[letter]; found {
			if 1 == count {
				delete(occurrences, letter)
			} else {
				occurrences[letter]--
			}
		} else {
			return false
		}
	}

	return (len(occurrences) == 0)
}

func GetInput() (str1, str2 string) {

	fmt.Println("Please enter first string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str1 = scanner.Text()

	fmt.Println("Please enter second string:")
	scanner.Scan()
	str2 = scanner.Text()

	return
}

func sanitizeInput(str1, str2 string) (string, string) {
	str1 = strings.ReplaceAll(str1, " ", "")
	str1 = strings.ToLower(str1)

	str2 = strings.ReplaceAll(str2, " ", "")
	str2 = strings.ToLower(str2)
	return str1, str2
}
