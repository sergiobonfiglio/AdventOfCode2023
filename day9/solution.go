package main

import (
	"AdventOfCode2023/utils"
	"strings"
)

func getDiffs(values []int) []int {
	var diffs []int
	for i := 0; i < len(values)-1; i++ {
		diffs = append(diffs, values[i+1]-values[i])
	}
	return diffs
}

func allZero(values []int) bool {
	for _, val := range values {
		if val != 0 {
			return false
		}
	}
	return true
}

func part1(input string) any {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		values := utils.ToIntArray(line)
		lastVal := []int{values[len(values)-1]}

		currVal := values
		allZeroes := false

		for !allZeroes {
			diffs := getDiffs(currVal)

			//fmt.Printf("diffs: %v\n", diffs)

			lastVal = append(lastVal, diffs[len(diffs)-1])

			allZeroes = allZero(diffs)
			currVal = diffs
		}

		pred := 0
		for _, v := range lastVal {
			pred += v
		}
		//fmt.Printf("last values: %v => pred: %d\n", lastVal, pred)

		sum += pred
	}

	return sum
}

func part2(input string) any {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		values := utils.ToIntArray(line)

		//fmt.Printf("-------------\nvalues: %v\n", values)

		firstVal := []int{values[0]}

		currVal := values
		allZeroes := false

		for !allZeroes {
			diffs := getDiffs(currVal)

			//fmt.Printf("diffs: %v\n", diffs)

			firstVal = append(firstVal, diffs[0])

			allZeroes = allZero(diffs)
			currVal = diffs
		}

		pred := 0
		for i := 0; i < len(firstVal); i++ {
			if i%2 == 0 {
				pred += firstVal[i]
			} else {
				pred -= firstVal[i]
			}
		}
		//fmt.Printf("first values: %v => pred: %d\n", firstVal, pred)

		sum += pred
	}
	return sum

}
