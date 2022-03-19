package c

import (
	"errors"
	"fmt"
	"strconv"
)

func GetPalindromeNumber(digitCount int) (answer string) {
	maxNumber := int((digitCount + 1) / 2)

	for currentNumber := 1; currentNumber <= maxNumber; currentNumber++ {
		answer += strconv.Itoa(currentNumber)
	}

	if digitCount%2 == 0 {
		answer += strconv.Itoa(maxNumber)
	}

	for currentNumber := maxNumber - 1; currentNumber > 0; currentNumber-- {
		answer += strconv.Itoa(currentNumber)
	}
	return
}

func GetInput() (digitCount int, err error) {

	fmt.Println("Please enter the digits count:")
	_, err = fmt.Scanf("%d", &digitCount)
	if digitCount > 18 || digitCount < 1 {
		err = errors.New("can't use any integer greater than 18 or less than 1")
	}

	return
}
