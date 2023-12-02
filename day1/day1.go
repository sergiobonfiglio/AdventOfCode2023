package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readInput() [][]int {

	readFile, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var data [][]int

	for fileScanner.Scan() {

		var chunk []int
		for fileScanner.Text() != "" {
			cal, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				panic(err)
			}
			chunk = append(chunk, cal)
			fileScanner.Scan()
		}
		data = append(data, chunk)

	}

	err = readFile.Close()
	if err != nil {
		panic(err)
	}

	return data
}

func SolveDay() {

	var calByElf = readInput()

	var sums []int
	var maxCal = 0
	for _, cals := range calByElf {

		var elfSum = 0
		for _, cal := range cals {
			elfSum += cal
		}
		sums = append(sums, elfSum)

		if elfSum > maxCal {
			maxCal = elfSum
		}
	}

	fmt.Println("Part 1:", maxCal)

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	var top3 = 0
	for i := 0; i < 3; i++ {
		top3 += sums[i]
	}

	fmt.Println("Part 2:", top3)
}
