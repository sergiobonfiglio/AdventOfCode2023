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

func FilterNil[T any](x []*T) []*T {
	var next []*T
	for _, el := range x {
		if el != nil {
			next = append(next, el)
		}
	}
	return next
}
