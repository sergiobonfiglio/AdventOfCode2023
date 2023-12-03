package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readInput() []string {

	readFile, err := os.Open("day1/input.txt")
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

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func SolveDay() {

	var input = readInput()

	sum := getSumFirstLastDigit(input)

	fmt.Println("Part 1:", sum)

	sum2 := getSumFirstLastDigitWords(input)
	fmt.Println("Part 2:", sum2)
}

func getSumFirstLastDigit(input []string) int {
	var sum int
	for _, row := range input {

		var first rune
		for i := 0; i < len(row); i++ {
			r := rune(row[i])
			if unicode.IsDigit(r) {
				first = r
				break
			}
		}

		var last rune
		for i := len(row) - 1; i >= 0; i-- {
			r := rune(row[i])
			if unicode.IsDigit(r) {
				last = r
				break
			}
		}

		val, err := strconv.Atoi(string(first) + string(last))
		if err != nil {
			panic("error")
		}

		sum += val
	}
	return sum
}

func getSumFirstLastDigitWords(input []string) int {
	var sum int
	for _, row := range input {

		for digit, value := range digitMap {
			if strings.Contains(row, digit) {
				row = strings.Replace(
					row,
					digit,
					string(digit[0])+strconv.Itoa(value)+string(digit[len(digit)-1]),
					-1)
			}
		}

		var first rune
		for i := 0; i < len(row); i++ {
			r := rune(row[i])
			if unicode.IsDigit(r) {
				first = r
				break
			}
		}

		var last rune
		for i := len(row) - 1; i >= 0; i-- {
			r := rune(row[i])
			if unicode.IsDigit(r) {
				last = r
				break
			}
		}

		val, err := strconv.Atoi(string(first) + string(last))
		if err != nil {
			panic("error")
		}

		sum += val
	}
	return sum
}
