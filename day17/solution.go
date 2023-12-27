package main

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func part1(input string) any {
	var matrix [][]int
	r := 0
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		matrix = append(matrix, []int{})
		for _, c := range line {
			digit, err := strconv.Atoi(string(c))
			if err != nil {
				panic(-1)
			}
			matrix[r] = append(matrix[r], digit)
		}
		r++
	}

	graph := buildGraph2(matrix, 1, 3)

	minPath := Dijkstra(graph)

	if minPath == nil {
		return nil
	}

	fmt.Printf("Path [%d]: %s\n\n%s\n", minPath.Cost(), minPath.toString2(), minPath.toString())
	return minPath.Cost()
}

func part1_old(input string) any {
	var matrix [][]int
	r := 0
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		matrix = append(matrix, []int{})
		for _, c := range line {
			digit, err := strconv.Atoi(string(c))
			if err != nil {
				panic(-1)
			}
			matrix[r] = append(matrix[r], digit)
		}
		r++
	}

	graph := buildGraphOld(matrix)

	//minPath := A_Star(graph)
	//minPath := Dijkstra(graph)

	paths := KShortestPaths(graph, 128, func(p *Path) bool {
		return p.isValidForCrucibles()
	})

	minPath := paths[0]
	for _, path := range paths {
		if path.Cost() < minPath.Cost() {
			tmp := path
			minPath = tmp
		}
	}

	if minPath == nil {
		return nil
	}

	fmt.Printf("Path [%d]: %s\n", minPath.Cost(), minPath.toString())
	return minPath.Cost()
}

func areOpposite(a, b rune) bool {
	return (a == '>' && b == '<') || (a == '<' && b == '>') || (a == 'v' && b == '^') || (a == '^' && b == 'v')
}
func (p Path) isValidForCrucibles() bool {

	consecutive := 0
	lastDir := '*'
	for _, edge := range p.edges {
		if lastDir == edge.Dir {
			consecutive++
		} else {
			consecutive = 0

			if areOpposite(lastDir, edge.Dir) {
				return false
			}

			lastDir = edge.Dir
		}

		if consecutive > 2 {
			return false
		}
	}

	return true
}
func (p Path) isValidForUltraCrucibles(dest *Vertex) bool {

	if !p.edges[len(p.edges)-1].Dst.Equals(dest) {
		return true
	}

	consecutive := math.MaxInt
	lastDir := '*'
	//lastConsecutive := 3
	minConsecutive := math.MaxInt
	for _, edge := range p.edges {
		if lastDir == edge.Dir {
			consecutive++
		} else {
			//lastConsecutive = consecutive

			if consecutive < minConsecutive {
				minConsecutive = consecutive
			}

			consecutive = 0

			if areOpposite(lastDir, edge.Dir) {
				return false
			}

			lastDir = edge.Dir
		}

		if consecutive > 9 {
			return false
		}
	}

	if minConsecutive < 3 {
		return false
	}

	return true
}

func part2(input string) any {
	var matrix [][]int
	r := 0
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		matrix = append(matrix, []int{})
		for _, c := range line {
			digit, err := strconv.Atoi(string(c))
			if err != nil {
				panic(-1)
			}
			matrix[r] = append(matrix[r], digit)
		}
		r++
	}

	graph := buildGraph2(matrix, 4, 10)

	minPath := Dijkstra(graph)

	if minPath == nil {
		return nil
	}

	fmt.Printf("Path [%d]: %s\n\n%s\n", minPath.Cost(), minPath.toString2(), minPath.toString())
	return minPath.Cost()

}
func part2Old(input string) any {
	var matrix [][]int
	r := 0
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		matrix = append(matrix, []int{})
		for _, c := range line {
			digit, err := strconv.Atoi(string(c))
			if err != nil {
				panic(-1)
			}
			matrix[r] = append(matrix[r], digit)
		}
		r++
	}

	graph := buildGraphOld(matrix)

	//minPath := A_Star(graph)
	//minPath := Dijkstra(graph)

	paths := KShortestPaths(graph, 128, func(p *Path) bool {
		return p.isValidForUltraCrucibles(graph.Dst)
	})

	minPath := paths[0]
	for _, path := range paths {
		if path.Cost() < minPath.Cost() {
			tmp := path
			minPath = tmp
		}
	}

	if minPath == nil {
		return nil
	}

	fmt.Printf("Path [%d]: %s\n", minPath.Cost(), minPath.toString())
	return minPath.Cost()

}

func getPathFromEdge(v *Vertex, prev map[VertexKey]*Edge) *Path {

	var edges []*Edge
	curr := v
	for curr != nil {
		currEdge, prevFound := prev[curr.Key()]

		if prevFound {
			edges = append(edges, currEdge)
			curr = currEdge.Src
		} else {
			curr = nil
		}
	}
	slices.Reverse(edges)
	return &Path{edges: edges}
}

func Dijkstra(graph *Graph) *Path {

	dist := map[VertexKey]int{}
	prev := map[VertexKey]*Edge{}

	pq := newPriorityQueue[Vertex]()

	for i, vertex := range graph.Vertices {
		pr := math.MaxInt
		if vertex == graph.Src {
			pr = 0
		}
		dist[vertex.Key()] = pr
		heap.Push(pq, &Item[Vertex]{
			value:    vertex,
			priority: pr,
			index:    i,
		})
	}

	for pq.Len() > 0 {

		uItem := heap.Pop(pq).(*Item[Vertex])
		u := uItem.value

		for _, edge := range u.Edges {
			v := edge.Dst

			qq, inQueue := pq.lut[v.getKey()]
			if inQueue && qq != nil {
				alt := edge.Weight + dist[u.Key()]
				if alt < dist[v.Key()] {

					prev[v.Key()] = edge

					vItem := pq.lut[v.getKey()]
					vItem.priority = alt
					heap.Fix(pq, vItem.index)
					dist[v.Key()] = alt

				}
			}
		}
	}

	return getPathFromEdge(graph.Dst, prev)
}

type Graph struct {
	Src      *Vertex
	Dst      *Vertex
	Vertices []*Vertex
	//Edges []Edge
}

type Vertex struct {
	R     int
	C     int
	Edges []*Edge
	Label string
	Z     int
	D     rune
}

type VertexKey string

func (v Vertex) Key() VertexKey {
	return VertexKey(v.Label)
}
func (v Vertex) Equals(v2 *Vertex) bool {
	return v.R == v2.R && v.C == v2.C
}

func (v Vertex) AddEdge(v2 *Vertex, weight int, dir rune) Vertex {
	edges := append(v.Edges, &Edge{
		Src:    &v,
		Dst:    v2,
		Weight: weight,
		Dir:    dir,
	})
	return Vertex{
		R:     v.R,
		C:     v.C,
		Edges: edges,
		Label: rczKey(v.R, v.C, v.Z, v.D),
		Z:     v.Z,
		D:     v.D,
	}
}

type Edge struct {
	Src    *Vertex
	Dst    *Vertex
	Weight int
	Dir    rune
}

func rcKey(r, c int) string {
	return strconv.Itoa(r) + "_" + strconv.Itoa(c)
}
func buildGraphOld(matrix [][]int) *Graph {

	vMap := map[string]*Vertex{}

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			k := rcKey(r, c)
			vMap[k] = &Vertex{R: r, C: c}
		}
	}

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			k := rcKey(r, c)
			currV := vMap[k]

			if upV, found := vMap[rcKey(r-1, c)]; found {
				*currV = currV.AddEdge(upV, matrix[r-1][c], '^')
			}
			if downV, found := vMap[rcKey(r+1, c)]; found {
				*currV = currV.AddEdge(downV, matrix[r+1][c], 'v')
			}
			if leftV, found := vMap[rcKey(r, c-1)]; found {
				*currV = currV.AddEdge(leftV, matrix[r][c-1], '<')
			}
			if rightV, found := vMap[rcKey(r, c+1)]; found {
				*currV = currV.AddEdge(rightV, matrix[r][c+1], '>')
			}

		}
	}
	var vertices []*Vertex
	for _, v := range vMap {
		vertices = append(vertices, v)
	}

	return &Graph{
		Src:      vMap[rcKey(0, 0)],
		Dst:      vMap[rcKey(len(matrix)-1, len(matrix[0])-1)],
		Vertices: vertices,
	}
}

func rczKey(r, c, z int, p rune) string {
	return strconv.Itoa(r) + "_" + strconv.Itoa(c) + "_" + strconv.Itoa(z) + "_" + string(p)
}

var DIRS = []rune{'>', '<', '^', 'v'}

func buildGraph2(matrix [][]int, minStraight, maxStraight int) *Graph {

	vMap := map[string]*Vertex{}

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if r != 0 || c != 0 {
				for z := minStraight; z <= maxStraight; z++ {
					for _, d := range DIRS {
						k := rczKey(r, c, z, d)
						vMap[k] = &Vertex{R: r, C: c, Label: k, Z: z, D: d}
					}
				}
			}
		}
	}

	lastR := len(matrix) - 1
	lastC := len(matrix[0]) - 1
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			isStart := r == 0 && c == 0
			isEnd := r == lastR && c == lastC
			if !isStart && !isEnd {
				for z := minStraight; z <= maxStraight; z++ {
					for _, d := range DIRS {
						k := rczKey(r, c, z, d)
						currV := vMap[k]

						addEdges(matrix, vMap, currV)
					}

				}

			}

		}
	}
	//connect starting node (special case)
	startV := &Vertex{R: 0, C: 0, Label: rczKey(0, 0, 0, '*')}
	addEdges(matrix, vMap, startV)

	//add fake end node reachable with 0 weight from all other directional ending nodes
	endV := &Vertex{R: lastR, C: lastC,
		Label: rczKey(lastR, lastC, 0, '*')}
	for z := minStraight; z <= maxStraight; z++ {
		for _, d := range DIRS {
			k := rczKey(lastR, lastC, z, d)
			currV := vMap[k]
			*currV = currV.AddEdge(endV, 0, '#')
		}
	}

	vertices := []*Vertex{startV, endV}
	for _, v := range vMap {
		vertices = append(vertices, v)
	}

	return &Graph{
		Src:      startV,
		Dst:      endV,
		Vertices: vertices,
	}
}

func addEdges(matrix [][]int, vMap map[string]*Vertex, currV *Vertex) {

	r, c, d, z := currV.R, currV.C, currV.D, currV.Z

	for _, d2 := range DIRS {

		if !areOpposite(d, d2) {

			nextZ := 1
			if d == d2 {
				nextZ = z + 1 // will not exist if it exceeds maxStraight
			}

			var nextVKey string
			if d2 == '^' {
				nextVKey = rczKey(r-1, c, nextZ, d2)
			} else if d2 == 'v' {
				nextVKey = rczKey(r+1, c, nextZ, d2)
			} else if d2 == '>' {
				nextVKey = rczKey(r, c+1, nextZ, d2)
			} else if d2 == '<' {
				nextVKey = rczKey(r, c-1, nextZ, d2)
			}

			if nextV, found := vMap[nextVKey]; found {
				*currV = currV.AddEdge(nextV, matrix[nextV.R][nextV.C], d2)
			}
		}

	}
}

type Cell struct {
	R int
	C int
}

func (p *Cell) Up() {
	p.R--
}
func (p *Cell) Down() {
	p.R++
}
func (p *Cell) Left() {
	p.C--
}
func (p *Cell) Right() {
	p.C++
}

func (p *Cell) IsInside(rowLen, colLen int) bool {
	return p.R >= 0 && p.R < rowLen && p.C >= 0 && p.C < colLen
}
