package main

import (
	"AdventOfCode2022/day1"
	"AdventOfCode2022/day10"
	"AdventOfCode2022/day11"
	"AdventOfCode2022/day2"
	"AdventOfCode2022/day3"
	"AdventOfCode2022/day4"
	"AdventOfCode2022/day5"
	"AdventOfCode2022/day6"
	"AdventOfCode2022/day7"
	"AdventOfCode2022/day8"
	"AdventOfCode2022/day9"
	"fmt"
)

func main() {

	var solvers = []func(){
		day1.SolveDay,
		day2.SolveDay,
		day3.SolveDay,
		day4.SolveDay,
		day5.SolveDay,
		day6.SolveDay,
		day7.SolveDay,
		day8.SolveDay,
		day9.SolveDay,
		day10.SolveDay,
		day11.SolveDay,
	}

	for i, solver := range solvers {
		fmt.Printf("\n== Day %d ==\n", i+1)
		solver()
	}
}
