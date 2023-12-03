package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func readInput() []string {

	readFile, err := os.Open("day3/input.txt")
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

	sum := 0
	gearRatioSum := 0
	for i, row := range input {

		for c, s := range row {
			if !unicode.IsDigit(s) && s != '.' {
				// symbol
				nums := getAdjacentNums(input, i, c)

				for _, num := range nums {
					sum += num
				}

				if rune(row[c]) == '*' && len(nums) == 2 {
					gearRatioSum += nums[0] * nums[1]
				}

			}
		}

	}

	return sum, gearRatioSum
}

type Coord struct {
	row int
	col int
}

func getAdjacentNums(input []string, row int, col int) []int {

	var res []int

	width := len(input[0])
	height := len(input)
	var visited []Coord

	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r >= 0 && r < height && c >= 0 && c < width {
				cc := Coord{
					row: r,
					col: c,
				}

				isVisited := false
				for _, v := range visited {
					if v.row == r && v.col == c {
						isVisited = true
					}
				}

				if isVisited {
					continue
				}

				visited = append(visited, cc)

				s := rune(input[r][c])
				isDigit := unicode.IsDigit(s)

				if isDigit {
					digits := []rune{s}

					dc := c - 1
					//try left
					for dc >= 0 && unicode.IsDigit(rune(input[r][dc])) {
						visited = append(visited, Coord{
							row: r,
							col: dc,
						})
						digits = append([]rune{rune(input[r][dc])}, digits...)
						dc--
					}

					dc = c + 1
					//try right
					for dc < len(input[r]) && unicode.IsDigit(rune(input[r][dc])) {
						visited = append(visited, Coord{
							row: r,
							col: dc,
						})
						digits = append(digits, rune(input[r][dc]))
						dc++
					}

					partNum, err := strconv.Atoi(string(digits))
					if err != nil {
						panic("error")
					}

					res = append(res, partNum)
				}

			}
		}
	}

	return res
}
