package main

import (
	"math"
	"strings"
)

type Pos struct {
	r int
	c int
}

func (p *Pos) DistManhattan(p2 *Pos) int {
	return int(math.Abs(float64(p.r-p2.r)) + math.Abs(float64(p.c-p2.c)))
}

func part1(input string) any {
	var stars []*Pos
	colsWithStars := map[int]bool{}
	rowsWithStars := map[int]bool{}

	for r, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		starsCount := 0
		for c, s := range line {
			if s == '#' {
				starsCount++
				stars = append(stars, &Pos{
					r: r,
					c: c,
				})
				rowsWithStars[r] = true
				colsWithStars[c] = true
			}
		}

	}
	var expanded []string
	//expand rows
	for r, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		expanded = append(expanded, line)
		if !rowsWithStars[r] {
			expanded = append(expanded, line)
		}

	}

	//expand cols
	maxCols := len(expanded[0])
	shift := 0
	for c := 0; c < maxCols; c++ {
		if !colsWithStars[c] {
			for r, _ := range expanded {
				expanded[r] = expanded[r][:c+shift] + "." + expanded[r][c+shift:]
			}
			shift++
		}
	}

	//update stars the lazy way
	stars = []*Pos{}
	for r, row := range expanded {
		for c, s := range row {
			if s == '#' {
				stars = append(stars, &Pos{
					r: r,
					c: c,
				})
			}
		}
	}

	sum := 0
	pairs := 0
	for i := 0; i < len(stars)-1; i++ {
		for j := i + 1; j < len(stars); j++ {

			pairs++
			dist := stars[i].DistManhattan(stars[j])
			sum += dist
		}
	}

	return sum
}

func part2(input string) any {
	return _part2(input, 1000000)
}
func _part2(input string, expansion int) any {
	var stars []*Pos
	colsWithStars := map[int]bool{}
	rowsWithStars := map[int]bool{}

	for r, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		starsCount := 0
		for c, s := range line {
			if s == '#' {
				starsCount++
				stars = append(stars, &Pos{
					r: r,
					c: c,
				})
				rowsWithStars[r] = true
				colsWithStars[c] = true
			}
		}

	}

	sum := 0
	pairs := 0

	shiftAmt := max(1, expansion-1)

	for i := 0; i < len(stars)-1; i++ {
		for j := i + 1; j < len(stars); j++ {

			pairs++
			star1 := stars[i]
			star2 := stars[j]
			dist := star1.DistManhattan(star2)

			minR, maxR := min(star1.r, star2.r), max(star1.r, star2.r)
			minC, maxC := min(star1.c, star2.c), max(star1.c, star2.c)

			distShift := 0
			for r := minR + 1; r < maxR; r++ {
				if !rowsWithStars[r] {
					distShift += shiftAmt
				}
			}

			for c := minC + 1; c < maxC; c++ {
				if !colsWithStars[c] {
					distShift += shiftAmt
				}
			}

			sum += dist + distShift
		}
	}

	return sum

}
