package main

import (
	"slices"
	"strconv"
	"strings"
)

func part1(input string) any {
	hashSum := int32(0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ",")

		for _, part := range parts {
			hashSum += hash(part)
		}
	}

	return hashSum
}

type lens struct {
	label string
	focal int
}

func part2(input string) any {
	boxes := map[int32][]*lens{}
	for _, line := range strings.Split(input, "\n") {
		steps := strings.Split(line, ",")

		for _, step := range steps {
			if step[len(step)-1] == '-' {
				label := step[:len(step)-1]
				boxNum := hash(label)

				lensIx := slices.IndexFunc(boxes[boxNum], func(l *lens) bool {
					return l.label == label
				})

				if lensIx >= 0 {
					boxes[boxNum] = append(boxes[boxNum][:lensIx], boxes[boxNum][lensIx+1:]...)
				}

			} else {
				pp := strings.Split(step, "=")
				label, focalStr := pp[0], pp[1]
				boxNum := hash(label)
				focal, err := strconv.Atoi(focalStr)
				if err != nil {
					panic(-1)
				}

				lensIx := slices.IndexFunc(boxes[boxNum], func(l *lens) bool {
					return l.label == label
				})

				if lensIx >= 0 {
					boxes[boxNum][lensIx].focal = focal
				} else {
					boxes[boxNum] = append(boxes[boxNum], &lens{
						label: label,
						focal: focal,
					})
				}
			}
		}
	}

	sum := 0
	for k, lenses := range boxes {
		for i, l := range lenses {
			focPow := (1 + int(k)) * (i + 1) * l.focal
			sum += focPow
		}
	}

	return sum
}

func hash(s string) int32 {

	res := int32(0)
	for _, c := range s {
		res += c
		res *= 17
		res = res % 256
	}
	return res
}
