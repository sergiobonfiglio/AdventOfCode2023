package day8

import (
	"AdventOfCode2023/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(fileName string) []string {

	readFile, err := os.Open("day8/" + fileName)
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

	var input = readInput("input.txt")
	fmt.Println("Part 1:", solvePart1(input))

	fmt.Println("Part 2:", solvePart2(input))
}

type Node struct {
	Label    string
	Left     string
	Right    string
	isEnding bool
}

func parseNode(input string) *Node {
	parts := strings.Split(input, " = ")

	label := parts[0]
	pairs := strings.Trim(parts[1], "()")
	pairParts := strings.Split(pairs, ", ")

	return &Node{
		Label:    label,
		Left:     pairParts[0],
		Right:    pairParts[1],
		isEnding: strings.HasSuffix(label, "Z"),
	}
}

func solvePart1(input []string) int {

	path := map[string]*Node{}
	var moves string
	for i, row := range input {
		if i == 0 {
			moves = row
		} else {
			node := parseNode(row)
			path[node.Label] = node
		}
	}

	currNode := path["AAA"]
	endReached := false
	count := 0
	for !endReached {

		for _, m := range moves {
			if m == 'L' {
				currNode = path[currNode.Left]
			} else {
				currNode = path[currNode.Right]
			}
		}
		count++
		endReached = currNode.Label == "ZZZ"
	}

	return count * len(moves)
}

func solvePart2(input []string) int64 {
	path := map[string]*Node{}
	var moves string
	var startingNodes []*Node
	for i, row := range input {
		if i == 0 {
			moves = row
		} else {
			node := parseNode(row)
			path[node.Label] = node

			if strings.HasSuffix(node.Label, "A") {
				startingNodes = append(startingNodes, node)
			}
		}
	}

	currNodes := startingNodes
	count := 0

	var maxFinals []int64
	cycles := map[string]int{}
	for _, curr := range currNodes {
		var next = curr

		visited := map[string]int{}

		var finals []int
		count = 0
		for cycles[curr.Label] == 0 {
			for i, m := range moves {
				if m == 'L' {
					next = path[next.Left]
				} else {
					next = path[next.Right]
				}
				count++

				if next.isEnding {
					finals = append(finals, count)
				}

				mvIx := strconv.Itoa(i)
				key := next.Label + mvIx
				if visited[key] > 0 {
					//fmt.Printf("found cycle %s => %s, %d: %d; first visit: %d\n", curr.Label, next.Label, i, count, visited[key])
					//fmt.Printf("endings: %v\n", finals)
					cycles[curr.Label] = count
					break
				}
				visited[key] = count
			}
		}
		maxFinals = append(maxFinals, int64(finals[len(finals)-1]))
	}

	lcm := utils.LCM(maxFinals[0], maxFinals[1], maxFinals[2:]...)
	return lcm
}
