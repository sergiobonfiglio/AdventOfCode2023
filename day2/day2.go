package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {

	readFile, err := os.Open("day2/input.txt")
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

// The Elf would first like to know which games would have been possible if the bag contained only
// 12 red cubes, 13 green cubes, and 14 blue cubes?
const confRed = 12
const confGreen = 13
const confBlue = 14

func SolveDay() {

	var input = readInput()

	sol1 := solvePart1(input)

	fmt.Println("Part 1:", sol1)

	fmt.Println("Part 2:", "")
}

type Set struct {
	red   int
	green int
	blue  int
}

func newSet(s string) Set {
	cubes := strings.Split(s, ",")

	res := Set{}
	for _, cc := range cubes {
		cc = strings.TrimSpace(cc)

		parts := strings.Split(cc, " ")
		num, col := parts[0], parts[1]
		numInt, err := strconv.Atoi(num)
		if err != nil {
			panic("error")
		}
		if col == "red" {
			res.red = numInt
		} else if col == "blue" {
			res.blue = numInt
		} else if col == "green" {
			res.green = numInt
		} else {
			panic("unknown color: " + col)
		}

	}

	return res
}

func solvePart1(input []string) int {

	sumPossible := 0

	for game, row := range input {
		rowparts := strings.Split(row, ":")

		setsStr := strings.Split(rowparts[1], ";")

		maxSet := Set{}
		//var sets []Set
		for _, setStr := range setsStr {
			set := newSet(setStr)
			//sets = append(sets, set)

			if set.red > maxSet.red {
				maxSet.red = set.red
			}
			if set.green > maxSet.green {
				maxSet.green = set.green
			}
			if set.blue > maxSet.blue {
				maxSet.blue = set.blue
			}

		}

		//is it possible?
		possible := maxSet.blue <= confBlue && maxSet.red <= confRed && maxSet.green <= confGreen

		if possible {
			sumPossible += game + 1
		}

	}

	return sumPossible
}
