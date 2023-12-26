package main

import "container/heap"

func newPriorityQueue[T Keyable]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	heap.Init(pq)
	return pq
}

type Keyable interface {
	getKey() string
}

func (v Vertex) getKey() string {
	return string(v.Key())
}

// An Item is something we manage in a priority queue.
type Item[T Keyable] struct {
	value    *T  // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T Keyable] struct {
	items []*Item[T]
	lut   map[string]*Item[T]
}

func (pq *PriorityQueue[T]) Len() int {
	return len((*pq).items)
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	pqv := (*pq).items
	return pqv[i].priority < pqv[j].priority
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	pqv := (*pq).items
	pqv[i], pqv[j] = pqv[j], pqv[i]
	pqv[i].index = i
	pqv[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len((*pq).items)
	item := x.(*Item[T])
	item.index = n
	pq.items = append(pq.items, item)

	if pq.lut == nil {
		pq.lut = map[string]*Item[T]{}
	}

	key := (*item.value).getKey()
	pq.lut[key] = item
}

func (pq *PriorityQueue[T]) Pop() any {
	old := (*pq).items
	n := len(old)
	item := old[n-1]

	pq.lut[(*item.value).getKey()] = nil

	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	pq.items = old[0 : n-1]
	return item
}

var _ heap.Interface = &PriorityQueue[Vertex]{}
