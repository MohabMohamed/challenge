package d

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ExtremeType int8

const (
	minimum ExtremeType = iota
	maximum
)

type extremeTwoNumbers struct {
	First       int
	Second      int
	FirstIdx    int
	SecondIdx   int
	ExtremeType ExtremeType
}

func newExtremeTwoNumbers(ExtremeType ExtremeType) extremeTwoNumbers {

	if ExtremeType == minimum {
		return extremeTwoNumbers{
			First:       math.MaxInt,
			Second:      math.MaxInt,
			FirstIdx:    0,
			SecondIdx:   0,
			ExtremeType: minimum,
		}
	} else {
		return extremeTwoNumbers{
			First:       math.MinInt,
			Second:      math.MinInt,
			FirstIdx:    0,
			SecondIdx:   0,
			ExtremeType: maximum,
		}
	}
}

func GetMaxMinSumIndices(numbers []int) (minIndices, maxIndices []int) {
	minExtreme := newExtremeTwoNumbers(minimum)
	maxExtreme := newExtremeTwoNumbers(maximum)

	for i, number := range numbers {
		minExtreme.proccess(number, i)
		maxExtreme.proccess(number, i)

	}

	minIndices = append(minIndices, minExtreme.FirstIdx)
	minIndices = append(minIndices, minExtreme.SecondIdx)

	maxIndices = append(maxIndices, maxExtreme.FirstIdx)
	maxIndices = append(maxIndices, maxExtreme.SecondIdx)

	return
}

func (twoNumbers *extremeTwoNumbers) proccess(number, idx int) {
	if twoNumbers.ExtremeType == minimum {
		twoNumbers.proccessMin(number, idx)
	} else {
		twoNumbers.proccessMax(number, idx)
	}
}

func (twoNumbers *extremeTwoNumbers) proccessMin(number, idx int) {

	if number < twoNumbers.First {
		twoNumbers.Second = twoNumbers.First
		twoNumbers.SecondIdx = twoNumbers.FirstIdx

		twoNumbers.First = number
		twoNumbers.FirstIdx = idx

	} else if number < twoNumbers.Second {
		twoNumbers.Second = number
		twoNumbers.SecondIdx = idx
	}

}

func (twoNumbers *extremeTwoNumbers) proccessMax(number, idx int) {
	if number > twoNumbers.Second {
		twoNumbers.First = twoNumbers.Second
		twoNumbers.FirstIdx = twoNumbers.SecondIdx

		twoNumbers.Second = number
		twoNumbers.SecondIdx = idx

	} else if number > twoNumbers.First {
		twoNumbers.First = number
		twoNumbers.FirstIdx = idx
	}
}

func GetInput() ([]int, error) {

	fmt.Println("Please enter the numbers separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	integersStr := scanner.Text()

	integerArray, err := convertIntegersStrToIntegerArray(integersStr)

	if err != nil {
		return nil, err
	}

	if len(integerArray) < 3 {
		return nil, errors.New("the numbers count should be more than 2")
	}

	return integerArray, nil

}

func convertIntegersStrToIntegerArray(integersStr string) ([]int, error) {
	integersStr = strings.TrimSpace(integersStr)
	strArray := strings.Fields(integersStr)

	integerArray := make([]int, len(strArray))

	var err error

	for idx, numberString := range strArray {
		integerArray[idx], err = strconv.Atoi(numberString)

		if err != nil {
			return nil, err
		}

	}

	return integerArray, nil
}
