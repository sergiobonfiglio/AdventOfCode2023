package main

import (
	"container/heap"
	"math"
	"slices"
)

func DistManhattan(p1 *Vertex, p2 *Vertex) int {
	return int(math.Abs(float64(p1.R-p2.R)) +
		math.Abs(float64(p1.C-p2.C)))
}

func getPath(v *Vertex, prev map[VertexKey]*Vertex) *Path {
	var edges []*Edge

	current := v
	prevV := prev[current.Key()]

	for prevV != nil {

		for _, edge := range prevV.Edges {
			if edge.Dst.Key() == current.Key() {
				tmp := edge
				edges = append(edges, tmp)
				current = prevV
				prevV = prev[current.Key()]
			}
		}
	}
	slices.Reverse(edges)

	p := newPath(edges)
	//p = p.addEdges(edges...)

	return p
}

func A_Star(graph *Graph) *Path {

	prev := map[VertexKey]*Vertex{}

	gScore := map[VertexKey]int{}
	gScore[graph.Src.Key()] = 0

	fScore := map[VertexKey]int{}
	//fScore[graph.Src.Key()] = DistManhattan(graph.Src, graph.Dst)
	fScore[graph.Src.Key()] = Dijkstra(graph).Cost()

	openSet := newPriorityQueue[Vertex]()

	heap.Push(openSet, &Item[Vertex]{
		value:    graph.Src,
		priority: fScore[graph.Src.Key()],
		index:    0,
	})
	destKey := graph.Dst.Key()

	for openSet.Len() > 0 {

		current := heap.Pop(openSet).(*Item[Vertex])

		if current.value.Key() == destKey {
			return getPath(current.value, prev)
		}

		for _, edge := range current.value.Edges {

			currPath := getPath(current.value, prev)
			currPath.edges = append(currPath.edges, edge)
			if !currPath.isValidForCrucibles() {
				continue
			}

			tentativeGScore := gScore[current.value.Key()] + edge.Weight
			currDstKey := edge.Dst.Key()
			currScore, found := gScore[currDstKey]
			if !found || tentativeGScore < currScore {
				prev[currDstKey] = current.value
				gScore[currDstKey] = tentativeGScore
				//fScore[currDstKey] = tentativeGScore + DistManhattan(edge.Dst, graph.Dst)
				hCost := Dijkstra(&Graph{
					Src:      edge.Dst,
					Dst:      graph.Dst,
					Vertices: graph.Vertices,
				}).Cost()
				fScore[currDstKey] = tentativeGScore + hCost

				if _, inSet := openSet.lut[string(currDstKey)]; !inSet {
					heap.Push(openSet, &Item[Vertex]{
						value:    edge.Dst,
						priority: fScore[currDstKey],
						index:    openSet.Len(),
					})
				}
			}

		}

	}

	return nil
}
