package utils

import (
	"strconv"
	"strings"
)

func ToIntegerArray[T int | int64](str string) []T {
	parts := strings.Split(strings.TrimSpace(str), " ")

	var res []T
	for _, part := range parts {
		if part != "" {
			n, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				panic("error")
			}
			res = append(res, T(n))
		}
	}
	return res
}

func ToIntArray(str string) []int {
	return ToIntegerArray[int](str)
}

func ToInt64Array(str string) []int64 {
	return ToIntegerArray[int64](str)
}

func GCD[T int | int64](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM[T int | int64](a, b T, integers ...T) T {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
