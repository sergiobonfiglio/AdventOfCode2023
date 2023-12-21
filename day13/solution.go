package main

import (
	"strings"
)

func part1(input string) any {

	cur := 0
	var patterns [][]string
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			cur++
		} else {
			if len(patterns) == 0 {
				patterns = append(patterns, []string{})
			}
			if cur > len(patterns)-1 {
				patterns = append(patterns, []string{})
			}

			patterns[cur] = append(patterns[cur], line)
		}

	}
	sumHor := 0
	sumVert := 0
	for _, pattern := range patterns {
		hRefl := getHorizontalReflection(pattern)
		vRefl := getVerticalReflection(pattern)

		//fmt.Printf("pattern\n")
		//fmt.Printf("horiz: %d, vert: %d\n", hRefl, vRefl)
		if hRefl != -1 {
			sumHor += hRefl
		}
		if vRefl != -1 {
			sumVert += vRefl
		}

	}

	return sumVert + 100*sumHor
}

func part2(input string) any {

	cur := 0
	var patterns [][]string
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			cur++
		} else {
			if len(patterns) == 0 {
				patterns = append(patterns, []string{})
			}
			if cur > len(patterns)-1 {
				patterns = append(patterns, []string{})
			}

			patterns[cur] = append(patterns[cur], line)
		}

	}
	sumHor := 0
	sumVert := 0
	for _, pattern := range patterns {
		hRefl := getHorizontalReflectionSmudge(pattern)
		vRefl := getVerticalReflectionSmudge(pattern)

		//fmt.Printf("pattern\n")
		//fmt.Printf("horiz: %d, vert: %d\n", hRefl, vRefl)
		if hRefl != -1 {
			sumHor += hRefl
		}
		if vRefl != -1 {
			sumVert += vRefl
		}

	}

	return sumVert + 100*sumHor
}

func getHorizontalReflectionSmudge(pattern []string) int {
	var candidates []int
	for i := 1; i < len(pattern); i++ {
		if getSingleDiffChar(pattern[i], pattern[i-1]) >= -1 {
			candidates = append(candidates, i-1)
		}
	}

	for _, cand := range candidates {

		isRefl := true
		smudgeCount := 0
		for i := 0; cand-i >= 0 && cand+1+i < len(pattern) && cand+1+i >= 0 && isRefl && smudgeCount < 2; i++ {
			diff := getSingleDiffChar(pattern[cand-i], pattern[cand+1+i])
			if diff < -1 {
				isRefl = false
			}
			if diff >= 0 {
				smudgeCount++
			}
		}

		if isRefl && smudgeCount == 1 {
			return cand + 1
		}
	}

	return -1
}

func getVerticalReflectionSmudge(pattern []string) int {
	var candidates []int

	colPrev := getCol(pattern, 0)
	for i := 1; i < len(pattern[0]); i++ {
		colNext := getCol(pattern, i)

		if getSingleDiffChar(colPrev, colNext) >= -1 {
			candidates = append(candidates, i-1)
		}
		colPrev = colNext
	}

	for _, cand := range candidates {

		isRefl := true
		smudgeCount := 0
		for i := 0; cand-i >= 0 && cand+1+i >= 0 && cand+1+i < len(pattern[0]) && isRefl && smudgeCount < 2; i++ {
			colNext := getCol(pattern, cand-i)
			colPrev := getCol(pattern, cand+1+i)
			diff := getSingleDiffChar(colNext, colPrev)
			if diff < -1 {
				isRefl = false
			}
			if diff >= 0 {
				smudgeCount++
			}
		}

		if isRefl && smudgeCount == 1 {
			return cand + 1
		}
	}

	return -1
}

func getHorizontalReflection(pattern []string) int {
	var candidates []int
	for i := 1; i < len(pattern); i++ {
		if pattern[i] == pattern[i-1] {
			candidates = append(candidates, i-1)
		}
	}

	for _, cand := range candidates {

		isRefl := true
		for i := 1; cand-i >= 0 && cand+1+i < len(pattern) && cand+1+i >= 0 && isRefl; i++ {
			if pattern[cand-i] != pattern[cand+1+i] {
				isRefl = false
			}
		}

		if isRefl {
			return cand + 1
		}
	}

	return -1
}

func getCol(pattern []string, c int) string {
	col := ""

	for i := 0; i < len(pattern); i++ {
		col += string(pattern[i][c])
	}

	return col
}
func getVerticalReflection(pattern []string) int {
	var candidates []int

	colPrev := getCol(pattern, 0)
	for i := 1; i < len(pattern[0]); i++ {
		colNext := getCol(pattern, i)

		if colPrev == colNext {
			candidates = append(candidates, i-1)
		}
		colPrev = colNext
	}

	for _, cand := range candidates {

		isRefl := true

		for i := 1; cand-i >= 0 && cand+1+i >= 0 && cand+1+i < len(pattern[0]) && isRefl; i++ {
			colNext := getCol(pattern, cand-i)
			colPrev := getCol(pattern, cand+1+i)
			if colNext != colPrev {
				isRefl = false
			}
		}

		if isRefl {
			return cand + 1
		}
	}

	return -1
}

func getSingleDiffChar(a, b string) int {

	diff := -1

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if diff != -1 {
				return -2 // too many
			}
			diff = i
		}
	}

	return diff

}
