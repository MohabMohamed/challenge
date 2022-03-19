package a

import "fmt"

func buildLongestSufPre(needle string) []int {
	needleSize := len(needle)
	longestSufPre := make([]int, needleSize)
	longestSufPre[0] = 0

	MatchLength := 0
	idx := 1

	for idx < needleSize {

		if needle[idx] == needle[MatchLength] {
			MatchLength++
			longestSufPre[idx] = MatchLength
			idx++
		} else if MatchLength == 0 {
			longestSufPre[idx] = 0
			idx++
		} else {
			MatchLength = longestSufPre[MatchLength-1]
		}

	}

	return longestSufPre
}

func FindNeedleIndex(haystack, needle string) int {
	needleSize := len(needle)
	haystackSize := len(haystack)

	if needleSize == 0 {
		return -1
	}

	longestSufPre := buildLongestSufPre(needle)

	hayIdx, needleIdx := 0, 0

	for hayIdx < haystackSize {

		if needle[needleIdx] == haystack[hayIdx] {
			needleIdx++
			hayIdx++
		} else {
			if needleIdx == 0 {
				hayIdx++
			} else {
				needleIdx = longestSufPre[needleIdx-1]
			}
		}

		if needleIdx == needleSize {
			return hayIdx - needleIdx
		}

	}

	return -1
}

func GetInput() (haystack, needle string) {

	fmt.Println("Please enter the haystack:")
	fmt.Scanln(&haystack)

	fmt.Println("Please enter the needle:")
	fmt.Scanln(&needle)
	return
}
