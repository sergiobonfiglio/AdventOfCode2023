package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type Path struct {
	edges []*Edge
	//vertices map[VertexKey]bool
}

func newPath(edges []*Edge) *Path {
	return &Path{
		edges: edges,
		//vertices: map[VertexKey]bool{},
	}
}

//func (p Path) addEdges(edges ...*Edge) *Path {
//	path := newPath()
//	path.edges = append(path.edges, edges...)
//	for _, e := range path.edges {
//		path.vertices[e.Src.Key()] = true
//	}
//	path.vertices[path.edges[len(path.edges)-1].Dst.Key()] = true
//
//	return path
//}

func (p Path) getKey() string {
	return p.toString()
}

func (p Path) Cost() int {
	sum := 0
	for _, edge := range p.edges {
		sum += edge.Weight
	}
	return sum
}

func (p Path) Dest() *Vertex {
	if len(p.edges) == 0 {
		return nil
	}
	return p.edges[len(p.edges)-1].Dst
}

func (p Path) toString() string {
	var n []string
	for _, e := range p.edges {
		n = append(n, fmt.Sprintf("(%d,%d) %s[%d]", e.Src.R, e.Src.C, string(e.Dir), e.Weight))
	}

	lastEdge := p.edges[len(p.edges)-1]
	n = append(n, fmt.Sprintf("(%d,%d)", lastEdge.Dst.R, lastEdge.Dst.C))
	pathStr := strings.Join(n, " ")
	return pathStr
}

func (p Path) ContainsVertex(v *Vertex) bool {
	//_, found := p.vertices[v.Key()]
	//return found

	for _, puEdge := range p.edges {
		if puEdge.Src.Equals(v) {
			return true
		}
	}

	return p.edges[len(p.edges)-1].Dst.Equals(v)

	//return false
}

func KShortestPaths(graph *Graph, K int, validateFn func(p *Path) bool) []*Path {

	var P []*Path

	//number of shortest paths found to node u
	counts := map[VertexKey]int{}
	for _, v := range graph.Vertices {
		counts[v.Key()] = 0
	}

	B := newPriorityQueue[Path]()

	startPath := newPath([]*Edge{
		{
			Src:    graph.Src,
			Dst:    graph.Src,
			Weight: 0,
			Dir:    0},
	})

	heap.Push(B, &Item[Path]{
		value:    startPath,
		priority: 0,
		index:    0,
	})

	minP := math.MaxInt
	for B.Len() > 0 && counts[graph.Dst.Key()] < K {

		puItem := heap.Pop(B).(*Item[Path])
		pu := puItem.value
		puDest := pu.Dest()
		puDestKey := puDest.Key()
		//counts[puDestKey]++

		if puDest.Equals(graph.Dst) {
			P = append(P, pu)
			puCost := pu.Cost()
			if puCost < minP {
				minP = puCost
			}
			//fmt.Printf("found path: cost = %d\n", pu.Cost())
		}

		if counts[puDestKey] < K {

			for i := 0; i < len(puDest.Edges); i++ {
				uEdge := puDest.Edges[i]

				if pu.ContainsVertex(uEdge.Dst) {
					continue
				}

				var pvEdges []*Edge
				pvEdges = append(pvEdges, pu.edges...)
				pvEdges = append(pvEdges, uEdge)

				pv := newPath(pvEdges)

				if validateFn(pv) {

					_, hasPath := B.lut[pv.getKey()]
					if !hasPath && pv.Cost() <= minP {
						counts[puDestKey]++

						heap.Push(B, &Item[Path]{
							value:    pv,
							priority: pv.Cost(),
							index:    B.Len(),
						})
					}

				}
			}
		}
	}

	return P
}
