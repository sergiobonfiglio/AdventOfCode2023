package main

import (
	"math"
	"strconv"
	"strings"
)

func part1(input string) any {

	start := Point[int]{
		Y: 0,
		X: 0,
	}
	points := []Point[int]{start}

	curr := start
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		dir := parts[0]
		meters, _ := strconv.Atoi(parts[1])
		color := parts[2]
		_ = color

		var next Point[int]
		if dir == "R" {
			next = curr.Right(meters)
		} else if dir == "L" {
			next = curr.Left(meters)
		} else if dir == "U" {
			next = curr.Up(meters)
		} else if dir == "D" {
			next = curr.Down(meters)
		}

		points = append(points, next)
		curr = next
	}

	perimeter := getPerimeter(points)
	rawArea := shoelaceArea(points)
	area := rawArea - (perimeter / 2) + 1

	total := area + perimeter
	return total
}

func getPerimeter[T int | int64](points []Point[T]) int {
	sum := 0
	closedPol := append(points, points[0])
	for i := 0; i < len(closedPol)-1; i++ {
		p1 := closedPol[i]
		p2 := closedPol[i+1]

		sum += int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
	}
	return sum
}
func shoelaceArea[T int | int64](points []Point[T]) int {

	sum1 := T(0)
	sum2 := T(0)
	closedPol := append(points, points[0])
	for i := 0; i < len(closedPol)-1; i++ {
		p1 := closedPol[i]
		p2 := closedPol[i+1]

		sum1 += p1.X * p2.Y
		sum2 += p1.Y * p2.X
	}

	return int(math.Abs(float64(sum1-sum2))) / 2
}

func part2(input string) any {
	start := Point[int64]{
		Y: 0,
		X: 0,
	}
	points := []Point[int64]{start}

	curr := start
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		parts := strings.Fields(line)

		encoded := parts[2]
		encoded = strings.Trim(encoded, "()")

		meters, err := strconv.ParseInt(encoded[1:6], 16, 64)
		if err != nil {
			panic(-1)
		}

		dir, err := strconv.Atoi(encoded[6:])
		if err != nil {
			panic(-1)
		}

		var next Point[int64]
		if dir == 0 {
			next = curr.Right(meters)
		} else if dir == 2 {
			next = curr.Left(meters)
		} else if dir == 3 {
			next = curr.Up(meters)
		} else if dir == 1 {
			next = curr.Down(meters)
		} else {
			panic(-1)
		}

		points = append(points, next)
		curr = next
	}

	perimeter := getPerimeter(points)
	rawArea := shoelaceArea(points)
	area := rawArea - (perimeter / 2) + 1
	//fmt.Printf("area: %d\nsum: %d\n", area, area+perimeter)

	total := area + perimeter
	return total
}

type Point[T int | int64] struct {
	Y T
	X T
}

func (p Point[T]) Up(d T) Point[T] {
	return Point[T]{
		Y: p.Y - d,
		X: p.X,
	}
}
func (p Point[T]) Down(d T) Point[T] {
	return Point[T]{
		Y: p.Y + d,
		X: p.X,
	}
}
func (p Point[T]) Left(d T) Point[T] {
	return Point[T]{
		Y: p.Y,
		X: p.X - d,
	}
}
func (p Point[T]) Right(d T) Point[T] {
	return Point[T]{
		Y: p.Y,
		X: p.X + d,
	}
}
