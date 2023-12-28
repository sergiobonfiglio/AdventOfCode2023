package main

import (
	"AdventOfCode2023/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const steps1 = 64
const steps2 = 26501365

func part1(input string) any {
	return _part1(input, steps1)
}
func _part1(input string, steps int) any {
	plots, start := parsePlots(input)

	curr := []utils.Cell{start}
	for i := 0; i < steps; i++ {
		next := map[string]utils.Cell{}
		for _, c := range curr {

			up := c.Up(1)
			down := c.Down(1)
			left := c.Left(1)
			right := c.Right(1)

			for _, p := range []utils.Cell{up, down, left, right} {
				if isInside(p, plots) && !isRock(p, plots) {
					next[key(p)] = p
				}
			}
		}

		var nextCurr []utils.Cell
		for _, cell := range next {
			nextCurr = append(nextCurr, cell)
		}
		curr = nextCurr

	}

	return len(curr)
}

func parsePlots(input string) ([][]rune, utils.Cell) {
	var matrix [][]rune
	start := utils.Cell{}
	for r, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		s := strings.IndexRune(line, 'S')
		if s > -1 {
			start = utils.Cell{
				R: r,
				C: s,
			}
		}
		matrix = append(matrix, []rune(line))
	}
	matrix[start.R][start.C] = '.'
	return matrix, start
}

func part2(input string) any {
	return _part2(input, steps2)

}
func _part2(input string, steps int) any {
	plots, start := parsePlots(input)

	quadrantsStart := map[string]int{}
	quadrantsCycleStart := map[string]int{}

	curr := []utils.Cell{start}
	for i := 0; i < 512; i++ {
		next := map[string]utils.Cell{}
		for _, c := range curr {

			up := c.Up(1)
			down := c.Down(1)
			left := c.Left(1)
			right := c.Right(1)

			for _, p := range []utils.Cell{up, down, left, right} {
				if !isRockWrap(p, plots) {
					next[key(p)] = p
				}
			}
		}

		var nextCurr []utils.Cell
		for _, cell := range next {
			nextCurr = append(nextCurr, cell)
		}
		curr = nextCurr

		quadrants := map[string]int{}
		for _, cell := range nextCurr {
			qr, qc := quadrant(cell, plots)
			qKey := strconv.Itoa(qr) + "_" + strconv.Itoa(qc)
			quadrants[qKey]++
			if _, found := quadrantsStart[qKey]; !found {
				quadrantsStart[qKey] = i
			}

			if quadrants[qKey] == 7520 {
				if _, found := quadrantsCycleStart[qKey]; !found {
					quadrantsCycleStart[qKey] = i
				}
			}
		}

		//for k, v := range quadrants {
		//	if k == "0_0" {
		//		//fmt.Printf("steps[%d]:%s: %d\n", i, k, v)
		//	}
		//}

	}

	fmt.Println("Starts:")
	var keys []string
	for k, _ := range quadrantsStart {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, k := range keys {
		qStart := quadrantsStart[k]
		qCycleStart := quadrantsCycleStart[k]

		if qCycleStart > 0 {
			fmt.Printf("%s: %d -> %d [diff: %d]\n", k, qStart, qCycleStart, qCycleStart-qStart)
		} else {
			fmt.Printf("%s: %d -> %d\n", k, qStart, qCycleStart)
		}
	}

	return len(curr)
}

func key(c utils.Cell) string {
	return strconv.Itoa(c.R) + "_" + strconv.Itoa(c.C)
}
func isRock(c utils.Cell, plots [][]rune) bool {
	return plots[c.R][c.C] == '#'
}

func quadrant(c utils.Cell, plots [][]rune) (int, int) {
	var rWrap int
	if c.R >= 0 {
		rWrap = c.R / len(plots)
	} else {
		rWrap = int(math.Floor(float64(c.R) / float64(len(plots))))
	}

	var cWrap int
	if c.C >= 0 {
		cWrap = c.C / len(plots[0])
	} else {
		cWrap = int(math.Floor(float64(c.C) / float64(len(plots[0]))))
	}

	return rWrap, cWrap
}
func isRockWrap(c utils.Cell, plots [][]rune) bool {
	var rWrap int
	if c.R >= 0 {
		rWrap = c.R % len(plots)
	} else {
		tmp := ((-c.R) - 1) % len(plots)
		rWrap = len(plots) - 1 - tmp
	}

	var cWrap int
	if c.C >= 0 {
		cWrap = c.C % len(plots[0])
	} else {
		tmp := ((-c.C) - 1) % len(plots[0])
		cWrap = len(plots[0]) - 1 - tmp
	}

	return plots[rWrap][cWrap] == '#'
}

func isInside(c utils.Cell, plots [][]rune) bool {
	return c.R >= 0 && c.R < len(plots) &&
		c.C >= 0 && c.C < len(plots[c.R])
}
