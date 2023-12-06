package day6

import (
	"AdventOfCode2023/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {

	readFile, err := os.Open("day6/input.txt")
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

func solveParts(input []string) (int, int64) {

	var times, distances []int
	for i, row := range input {
		parts := strings.Split(row, ":")
		if i == 0 {
			times = utils.ToIntArray(parts[1])
		} else if i == 1 {
			distances = utils.ToIntArray(parts[1])
		}
	}

	totPossSolutions := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		record := distances[i]

		discriminant := time*time - 4*record
		hasSolutions := discriminant > 0
		if hasSolutions {
			squaredDisc := math.Sqrt(float64(discriminant))
			sol1 := (float64(time) + squaredDisc) / 2
			sol2 := (float64(time) - squaredDisc) / 2

			intSols := int(math.Ceil(sol1) - math.Floor(sol2) - 1)
			totPossSolutions *= intSols

		}
	}

	totTimeStr, totRecordStr := "", ""
	for i := 0; i < len(times); i++ {
		totTimeStr += strconv.Itoa(times[i])
		totRecordStr += strconv.Itoa(distances[i])
	}

	totTime, err := strconv.ParseInt(totTimeStr, 10, 64)
	if err != nil {
		panic("err")
	}
	totRecord, err := strconv.ParseInt(totRecordStr, 10, 64)
	if err != nil {
		panic("err")
	}

	discriminant := totTime*totTime - 4*totRecord
	hasSolutions := discriminant > 0
	part2Sol := int64(0)
	if hasSolutions {
		squaredDisc := math.Sqrt(float64(discriminant))
		sol1 := (float64(totTime) + squaredDisc) / 2
		sol2 := (float64(totTime) - squaredDisc) / 2

		intSols := int64(math.Ceil(sol1) - math.Floor(sol2) - 1)
		part2Sol = intSols
	}

	return totPossSolutions, part2Sol
}
