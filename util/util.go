package util

import (
	"strconv"
	"unicode"
)

type Number interface {
	~int | ~int64 | ~float64
}

func NextNum[T Number](b string, i int) (T, int, bool) {
	for ; i < len(b) && !unicode.IsDigit(rune(b[i])) && b[i] != '-' && b[i] != '.'; i++ {
	}
	start := i
	for ; i < len(b) && (unicode.IsDigit(rune(b[i])) || b[i] == '.' || b[i] == '-'); i++ {
	}
	if start == i {
		return 0, i, false
	}
	num, err := strconv.ParseFloat(b[start:i], 64)
	if err != nil {
		return 0, start, false
	}

	return T(num), i, true
}

func GetNumsFromString[T Number](str string) []T {
	var nums []T
	i := 0
	for i < len(str) {
		num, nextI, ok := NextNum[T](str, i)
		if nextI == i {
			break
		}
		if ok {
			nums = append(nums, num)
		}
		i = nextI
	}
	return nums
}

// GetIntsFromStrings extracts all integers from a slice of strings
func GetNumsFromStringSlice[T Number](strs []string) []T {
	var allNums []T
	for _, str := range strs {
		nums := GetNumsFromString[T](str)
		allNums = append(allNums, nums...)
	}
	return allNums
}
