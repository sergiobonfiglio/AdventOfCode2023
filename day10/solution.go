package main

import (
	"AdventOfCode2023/utils"
	"slices"
	"strings"
)

type Pos struct {
	r    int
	c    int
	maze *[]string
	dist int
}

func (p *Pos) N() *Pos {
	if p.r == 0 {
		return nil
	}
	return &Pos{
		r:    p.r - 1,
		c:    p.c,
		maze: p.maze,
	}
}

func (p *Pos) S() *Pos {
	if p.r == len(*p.maze)-1 {
		return nil
	}
	return &Pos{
		r:    p.r + 1,
		c:    p.c,
		maze: p.maze,
	}
}

func (p *Pos) W() *Pos {
	if p.c == 0 {
		return nil
	}
	return &Pos{
		r:    p.r,
		c:    p.c - 1,
		maze: p.maze,
	}
}

func (p *Pos) E() *Pos {
	if p.c == len((*p.maze)[0])-1 {
		return nil
	}
	return &Pos{
		r:    p.r,
		c:    p.c + 1,
		maze: p.maze,
	}
}

func (p *Pos) getConnecting() []*Pos {
	cand := p.getCandidates()
	var connecting []*Pos
	for _, cc := range cand {
		if p.connectsTo(cc) {
			connecting = append(connecting, cc)
		}
	}
	return connecting
}

func (p *Pos) connectsTo(p2 *Pos) bool {
	return slices.ContainsFunc(p.getCandidates(), func(pos *Pos) bool {
		return pos.r == p2.r && pos.c == p2.c
	}) && slices.ContainsFunc(p2.getCandidates(), func(pos *Pos) bool {
		return pos.r == p.r && pos.c == p.c
	})
}

func (p *Pos) getCandidates() []*Pos {
	currPipe := (*p.maze)[p.r][p.c]

	if currPipe == 'S' {

		return utils.FilterNil([]*Pos{p.N(), p.S(), p.W(), p.E()})

	} else if currPipe == '-' {

		return utils.FilterNil([]*Pos{p.W(), p.E()})

	} else if currPipe == '|' {

		return utils.FilterNil([]*Pos{p.N(), p.S()})

	} else if currPipe == 'L' {

		return utils.FilterNil([]*Pos{p.N(), p.E()})

	} else if currPipe == 'J' {

		return utils.FilterNil([]*Pos{p.N(), p.W()})

	} else if currPipe == '7' {

		return utils.FilterNil([]*Pos{p.S(), p.W()})

	} else if currPipe == 'F' {

		return utils.FilterNil([]*Pos{p.S(), p.E()})

	}
	//| is a vertical pipe connecting north and south.
	//- is a horizontal pipe connecting east and west.
	//L is a 90-degree bend connecting north and east.
	//J is a 90-degree bend connecting north and west.
	//7 is a 90-degree bend connecting south and west.
	//F is a 90-degree bend connecting south and east.
	return []*Pos{}
}

func (p *Pos) isStart() bool {
	return (*p.maze)[p.r][p.c] == 'S'
}

func part1(input string) any {
	var maze = &[]string{}
	var startPos *Pos
	for r, line := range strings.Split(input, "\n") {
		if line != "" {
			tmp := append(*maze, line)
			*maze = tmp
			sPos := strings.Index(line, "S")
			if sPos != -1 {
				startPos = &Pos{
					r:    r,
					c:    sPos,
					maze: maze,
				}
			}
		}
	}

	candidates := startPos.getConnecting()

	for _, cand := range candidates {
		currPos := cand
		var prevPos *Pos = startPos
		dist := 1
		for currPos != nil && !currPos.isStart() {
			currPos.dist = dist
			dist++

			nextPoss := currPos.getConnecting()
			if nextPoss == nil || len(nextPoss) == 0 {
				//dead end
				currPos = nil
				break
			}

			//filter prev pos
			notPrev := nextPoss
			if prevPos != nil {
				notPrev = []*Pos{}
				for _, poss := range nextPoss {
					if poss.r != prevPos.r || poss.c != prevPos.c {
						notPrev = append(notPrev, poss)
					}
				}
			}

			if len(notPrev) == 0 {
				//dead end
				currPos = nil
				break
			}
			if len(notPrev) != 1 {
				panic("err")
			}
			prevPos = currPos
			currPos = notPrev[0]
		}

		if currPos != nil {
			//loop!
			return dist / 2
		}

	}

	return nil
}

func part2(input string) any {
	var maze = &[]string{}
	var startPos *Pos
	for r, line := range strings.Split(input, "\n") {
		if line != "" {
			tmp := append(*maze, line)
			*maze = tmp
			sPos := strings.Index(line, "S")
			if sPos != -1 {
				startPos = &Pos{
					r:    r,
					c:    sPos,
					maze: maze,
				}
			}
		}
	}

	candidates := startPos.getConnecting()

	for _, cand := range candidates {
		currPos := cand
		var prevPos *Pos = startPos
		dist := 1
		edges := []*Pos{startPos}
		for currPos != nil && !currPos.isStart() {
			edges = append(edges, currPos)
			currPos.dist = dist
			dist++

			nextPoss := currPos.getConnecting()
			if nextPoss == nil || len(nextPoss) == 0 {
				//dead end
				currPos = nil
				break
			}

			//filter prev pos
			notPrev := nextPoss
			if prevPos != nil {
				notPrev = []*Pos{}
				for _, poss := range nextPoss {
					if poss.r != prevPos.r || poss.c != prevPos.c {
						notPrev = append(notPrev, poss)
					}
				}
			}

			if len(notPrev) == 0 {
				//dead end
				currPos = nil
				break
			}
			if len(notPrev) != 1 {
				panic("err")
			}
			prevPos = currPos
			currPos = notPrev[0]
		}

		if currPos != nil {
			//loop!
			//fmt.Printf("start reached in %d steps\n", dist)
			//fmt.Printf("edges: %d\n", len(edges))

			//replace S with real pipe (needed to simplify to edge logic)
			var realSPipe rune
			realSPipe = getRealPipe(startPos, edges)
			//fmt.Printf("real S: %s\n", string(realSPipe))
			(*maze)[startPos.r] = strings.Replace((*maze)[startPos.r], "S", string(realSPipe), -1)

			edgeMap := map[int]map[int]*uint8{}
			for _, edge := range edges {
				tmp := (*maze)[edge.r][edge.c]
				if edgeMap[edge.r] == nil {
					edgeMap[edge.r] = map[int]*uint8{}
				}
				edgeMap[edge.r][edge.c] = &tmp
			}

			count := 0
			for r, row := range *maze {
				for c, _ := range row {
					isEdge := edgeMap[r][c] != nil
					if !isEdge {
						condHit := '*' //fake val
						edgesHit := 0
						for i := c + 1; i < len(row); i++ {
							currIsEdge := edgeMap[r][i] != nil
							if currIsEdge {
								val := *edgeMap[r][i]

								if val == '|' {
									edgesHit++
								}
								if rune(val) == condHit {
									edgesHit++
								}

								// only consider corners to hit if they bend in a way that is equivalent
								// to a vertical wall. Also, we only go from left to right
								if val == 'F' {
									condHit = 'J'
								} else if val == 'L' {
									condHit = '7'
								}

							}
						}
						if edgesHit%2 != 0 {
							//inside
							//fmt.Printf("inside: (%d, %d)\n", r, c)
							count++
						}
					}
				}
			}

			return count
		}

	}

	return nil
}

func getRealPipe(pos *Pos, edges []*Pos) rune {
	conn := pos.getConnecting()
	var connEdge []*Pos
	for _, c := range conn {
		if slices.ContainsFunc(edges, func(pp *Pos) bool {
			return pp.r == c.r && pp.c == c.c
		}) {
			connEdge = append(connEdge, c)
		}
	}

	if len(connEdge) != 2 {
		panic(1)
	}

	haveN := slices.ContainsFunc(connEdge, func(pp *Pos) bool {
		return pp.r == pos.r-1
	})

	haveS := slices.ContainsFunc(connEdge, func(pp *Pos) bool {
		return pp.r == pos.r+1
	})

	haveW := slices.ContainsFunc(connEdge, func(pp *Pos) bool {
		return pp.c == pos.c-1
	})
	haveE := slices.ContainsFunc(connEdge, func(pp *Pos) bool {
		return pp.c == pos.c+1
	})

	if haveN {
		if haveS {
			return '|'
		} else if haveW {
			return 'J'
		} else if haveE {
			return 'L'
		}
	} else if haveS {
		if haveW {
			return '7'
		} else if haveE {
			return 'F'
		}
	}
	return '-'

}
