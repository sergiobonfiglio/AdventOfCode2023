package day4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput() []string {

	readFile, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var data []string

	for fileScanner.Scan() {

		for fileScanner.Text() != "" {
			data = append(data, fileScanner.Text())
			fileScanner.Scan()
		}

	}

	err = readFile.Close()
	if err != nil {
		panic(err)
	}

	return data
}

func SolveDay() {

	var input = readInput()

	sol1, sol2 := solveParts(input)

	fmt.Println("Part 1:", sol1)

	fmt.Println("Part 2:", sol2)
}

func solveParts(input []string) (int, int) {

	totPoints := 0
	totCards := 0

	cardsCopies := map[int]int{}

	for card, row := range input {

		rowParts := strings.Split(row, ":")
		_, nums := rowParts[0], rowParts[1]

		numsSplit := strings.Split(nums, "|")

		winning := toIntArray(numsSplit[0])
		myNums := toIntArray(numsSplit[1])

		winningNums := 0
		for _, num := range myNums {
			if slices.Contains(winning, num) {
				winningNums++
			}
		}
		if winningNums > 0 {
			totPoints += int(math.Pow(2, float64(winningNums-1)))
		}

		mult := cardsCopies[card] + 1

		totCards += winningNums * mult

		for c := 0; c < mult; c++ {
			for i := 0; i < winningNums; i++ {
				cardsCopies[card+i+1]++
			}
		}

	}

	return totPoints, totCards + len(input)
}

func toIntArray(str string) []int {
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
