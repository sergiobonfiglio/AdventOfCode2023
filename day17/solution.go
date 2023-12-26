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

	graph := buildGraph(matrix)

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

	graph := buildGraph(matrix)

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

func getPrevDirs(curr *Vertex, prev map[*Vertex]*Edge) map[rune]int {
	dirMap := map[rune]int{}
	var prevEdges []*Edge

	currPrev := prev[curr]
	if currPrev != nil {
		for i := 0; i < 3; i++ {
			e := prev[currPrev.Src]
			if e == nil {
				break
			}
			dirMap[e.Dir]++
			prevEdges = append(prevEdges, e)
			currPrev = e
		}

	}

	return dirMap
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
					path := getPathFromEdge(v, prev)
					if !path.isValidForCrucibles() {
						prev[v.Key()] = nil
						continue
					}
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
}

type VertexKey string

func (v Vertex) Key() VertexKey {
	return VertexKey(strconv.Itoa(v.R) + "_" + strconv.Itoa(v.C))
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
func buildGraph(matrix [][]int) *Graph {

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
