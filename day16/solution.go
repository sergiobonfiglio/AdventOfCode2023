package main

import (
	"slices"
	"strconv"
	"strings"
)

type Ray struct {
	r       int
	c       int
	dir     rune // >, <, v, ^
	visited int
	lastNew int
}

func (r *Ray) Right() {
	r.c++
	r.dir = '>'
}
func (r *Ray) Left() {
	r.c--
	r.dir = '<'
}
func (r *Ray) Up() {
	r.r--
	r.dir = '^'
}
func (r *Ray) Down() {
	r.r++
	r.dir = 'v'
}
func (r *Ray) Forward() {
	if r.dir == '>' {
		r.Right()
	} else if r.dir == '<' {
		r.Left()
	} else if r.dir == 'v' {
		r.Down()
	} else if r.dir == '^' {
		r.Up()
	}
}

func (r *Ray) IsInside(rowsLen, colsLen int) bool {
	return r.r < rowsLen && r.r >= 0 &&
		r.c < colsLen && r.c >= 0
}

func (r *Ray) reflect(s uint8) {
	if s == '/' {
		if r.dir == '^' {
			r.Right()
		} else if r.dir == '<' {
			r.Down()
		} else if r.dir == 'v' {
			r.Left()
		} else if r.dir == '>' {
			r.Up()
		}
	} else if s == '\\' {
		if r.dir == '^' {
			r.Left()
		} else if r.dir == '<' {
			r.Up()
		} else if r.dir == 'v' {
			r.Right()
		} else if r.dir == '>' {
			r.Down()
		}
	}
}

func (r *Ray) split(s uint8) *Ray {
	if s == '|' {
		if r.dir == '^' || r.dir == 'v' {
			r.Forward()
		} else if r.dir == '<' || r.dir == '>' {
			ray2 := &Ray{
				r:   r.r,
				c:   r.c,
				dir: r.dir,
			}
			r.Up()
			ray2.Down()
			return ray2
		}
	} else if s == '-' {
		if r.dir == '^' || r.dir == 'v' {
			ray2 := &Ray{
				r:   r.r,
				c:   r.c,
				dir: r.dir,
			}
			r.Left()
			ray2.Right()
			return ray2
		} else if r.dir == '<' || r.dir == '>' {
			r.Forward()
		}
	}
	return nil
}

func part1(input string) any {

	var layout []string
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		layout = append(layout, line)
	}

	startPos := Ray{
		r:   0,
		c:   0,
		dir: '>',
	}

	return getTotEnergy(layout, startPos)
}

func getTotEnergy(layout []string, startPos Ray) int {
	totEnergy := 0
	energy := map[int]map[int]int{}
	for r := range layout {
		energy[r] = map[int]int{}
	}

	visitedSplit := map[string]bool{}

	rowsLen := len(layout)
	colsLen := len(layout[0])
	rays := []*Ray{&startPos}
	round := 0
	for len(rays) > 0 {
		round++
		var toDel []int
		var toAdd []*Ray
		for i, ray := range rays {
			if !ray.IsInside(rowsLen, colsLen) || (round-ray.lastNew) > (rowsLen*colsLen) {
				// is outside
				toDel = append(toDel, i)
				continue
			}

			// inside
			if energy[ray.r][ray.c] == 0 {
				totEnergy++
				ray.lastNew = round
			}
			energy[ray.r][ray.c]++

			s := layout[ray.r][ray.c]

			if s == '.' {
				ray.Forward()
			} else if s == '/' || s == '\\' {
				ray.reflect(s)
			} else if s == '|' || s == '-' {

				key := strconv.Itoa(ray.r) + "_" + strconv.Itoa(ray.c) + "_" + string(ray.dir)
				if visitedSplit[key] {
					toDel = append(toDel, i)
					continue
				}

				visitedSplit[key] = true
				ray2 := ray.split(s)
				if ray2 != nil {
					toAdd = append(toAdd, ray2)
				}
			}
		}
		var tmpRays []*Ray
		for i, ray := range rays {
			if !slices.Contains(toDel, i) {
				tmpRays = append(tmpRays, ray)
			}
		}

		rays = append(tmpRays, toAdd...)

		//fmt.Printf("round: %d/%d, energy: %d [%d], rays: %d\n", round, rowsLen*colsLen, totEnergy, lastEnergyChange, len(rays))
	}

	return totEnergy
}

func part2(input string) any {
	var layout []string
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		layout = append(layout, line)
	}

	maxEnergy := 0
	for r := 0; r < len(layout); r++ {
		left := getTotEnergy(layout, Ray{
			r:   r,
			c:   0,
			dir: '>',
		})
		if left > maxEnergy {
			maxEnergy = left
		}

		right := getTotEnergy(layout, Ray{
			r:   r,
			c:   0,
			dir: '<',
		})
		if right > maxEnergy {
			maxEnergy = right
		}
	}

	for c := 0; c < len(layout[0]); c++ {
		down := getTotEnergy(layout, Ray{
			r:   0,
			c:   c,
			dir: 'v',
		})
		if down > maxEnergy {
			maxEnergy = down
		}

		up := getTotEnergy(layout, Ray{
			r:   0,
			c:   c,
			dir: '^',
		})
		if up > maxEnergy {
			maxEnergy = up
		}
	}

	return maxEnergy
}
