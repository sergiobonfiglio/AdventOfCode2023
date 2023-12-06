package utils

import (
	"strconv"
	"strings"
)

func ToIntArray(str string) []int {
	parts := strings.Split(strings.TrimSpace(str), " ")

	var res []int
	for _, part := range parts {
		if part != "" {
			n, err := strconv.Atoi(part)
			if err != nil {
				panic("error")
			}
			res = append(res, n)
		}
	}
	return res
}

func ToInt64Array(str string) []int64 {
	parts := strings.Split(strings.TrimSpace(str), " ")

	var res []int64
	for _, part := range parts {
		if part != "" {
			n, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				panic("error")
			}
			res = append(res, n)
		}
	}
	return res
}
