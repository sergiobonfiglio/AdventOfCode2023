package main

import (
	"strings"
)

func part1(input string) any {
	var pattern []string

	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			pattern = append(pattern, line)
		}
	}

	//for each column
	stonesByRow := map[int]int{}
	for c := 0; c < len(pattern[0]); c++ {
		lastFree := 0
		for r := 0; r < len(pattern); r++ {
			curr := pattern[r][c]
			if curr == '#' {
				lastFree = r + 1
			} else if curr == '.' {
				//nothing
			} else {
				//rolling stone
				finalRow := min(r, lastFree)
				stonesByRow[finalRow]++
				lastFree++
			}
		}

	}

	sum := 0
	for row, stones := range stonesByRow {
		weight := len(pattern[0]) - row
		sum += weight * stones
	}

	return sum
}

func part2(input string) any {
	var pattern []string
	stones := map[int]map[int]bool{}

	for r, line := range strings.Split(input, "\n") {
		if line != "" {
			ll := strings.Replace(line, "O", ".", -1)
			pattern = append(pattern, ll)

			stones[r] = map[int]bool{}
			for c := 0; c < len(line); c++ {
				if line[c] == 'O' {
					stones[r][c] = true
				}
			}
		}
	}

	hashes := map[string]int{}
	hashes[stringMap(pattern, stones)] = -1
	startCycle := -1

	i := 0
	MaxCycles := 1000000000
	for ; i < MaxCycles; i++ {
		rollNorth(pattern, stones)
		rollWest(pattern, stones)
		rollSouth(pattern, stones)
		rollEast(pattern, stones)

		str := stringMap(pattern, stones)
		if hashes[str] > 0 {
			//fmt.Printf("Repetition after %d, from %d\n", i, hashes[str])
			startCycle = hashes[str]
			break
		}

		hashes[str] = i
	}

	diff := (MaxCycles-startCycle)%(i-startCycle) - 1
	for j := 0; j < diff; j++ {
		rollNorth(pattern, stones)
		rollWest(pattern, stones)
		rollSouth(pattern, stones)
		rollEast(pattern, stones)
	}

	sum := getLoad(pattern, stones)

	return sum

}
func getLoad(pattern []string, stones map[int]map[int]bool) int {
	sum := 0

	for row, stCols := range stones {
		rowCount := 0
		for _, ss := range stCols {
			if ss {
				rowCount++
			}
		}
		weight := len(pattern[0]) - row
		sum += weight * rowCount
	}

	return sum
}

func stringMap(pattern []string, stones map[int]map[int]bool) string {

	tot := ""
	for r := 0; r < len(pattern); r++ {
		row := ""
		for c := 0; c < len(pattern[0]); c++ {
			if stones[r][c] {
				row += "O"
			} else {
				row += string(pattern[r][c])
			}

		}
		tot += row + "\n"
	}
	return tot
}

func rollNorth(pattern []string, stones map[int]map[int]bool) {
	for c := 0; c < len(pattern[0]); c++ {
		lastFree := 0
		for r := lastFree; r < len(pattern); r++ {
			curr := pattern[r][c]
			if curr == '#' {
				lastFree = r + 1
			} else if stones[r][c] {
				//rolling stone
				finalRow := min(r, lastFree)
				stones[r][c] = false
				stones[finalRow][c] = true
				lastFree++
			}
		}

	}
}

func rollSouth(pattern []string, stones map[int]map[int]bool) {
	for c := 0; c < len(pattern[0]); c++ {
		lastFree := len(pattern) - 1
		for r := lastFree; r >= 0; r-- {
			curr := pattern[r][c]
			if curr == '#' {
				lastFree = r - 1
			} else if stones[r][c] {
				//rolling stone
				finalRow := max(r, lastFree)
				stones[r][c] = false
				stones[finalRow][c] = true
				lastFree--
			}
		}

	}
}

func rollWest(pattern []string, stones map[int]map[int]bool) {
	for r := 0; r < len(pattern); r++ {
		lastFree := 0
		for c := lastFree; c < len(pattern[0]); c++ {
			curr := pattern[r][c]
			if curr == '#' {
				lastFree = c + 1
			} else if stones[r][c] {
				//rolling stone
				finalCol := min(c, lastFree)
				stones[r][c] = false
				stones[r][finalCol] = true
				lastFree++
			}
		}

	}
}

func rollEast(pattern []string, stones map[int]map[int]bool) {
	for r := 0; r < len(pattern); r++ {
		lastFree := len(pattern[0]) - 1
		for c := lastFree; c >= 0; c-- {
			curr := pattern[r][c]
			if curr == '#' {
				lastFree = c - 1
			} else if stones[r][c] {
				//rolling stone
				finalCol := max(c, lastFree)
				stones[r][c] = false
				stones[r][finalCol] = true
				lastFree--
			}
		}
	}
}
