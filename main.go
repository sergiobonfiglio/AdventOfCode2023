package main

import (
	"AdventOfCode2023/day1"
	"AdventOfCode2023/day2"
	"AdventOfCode2023/day3"
	"AdventOfCode2023/day4"
	"AdventOfCode2023/day5"
	"AdventOfCode2023/day6"
	"AdventOfCode2023/day7"
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
	}

	for i, solver := range solvers {
		fmt.Printf("\n== Day %d ==\n", i+1)
		solver()
	}
}
