package leetcode

import (
	"container/heap"
	"fmt"
)

// return graph for problem 743
func getGraph(edges [][]int, N int) [][][]int {
	graph := make([][][]int, N)
	for i := 0; i < N; i++ {
		graph[i] = [][]int{}
	}
	for _, edge := range edges {
		source := edge[0] - 1
		dest, time := edge[1]-1, edge[2]
		e := []int{dest, time}
		graph[source] = append(graph[source], e)
	}
	return graph
}

// Print graph for problem 743
func printGraph(graph [][][]int) {
	for v, edges := range graph {
		fmt.Println(v)
		for _, edge := range edges {
			fmt.Println("   ", edge)
		}
	}
}

// pqGraphNode is something we manage in a priority queue.
type pqGraphNode struct {
	dest   int // The value of the item; arbitrary.
	weight int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A pqGraph implements heap.Interface and holds Items.
type pqGraph []*pqGraphNode

func (pq pqGraph) Len() int { return len(pq) }

func (pq pqGraph) Less(i, j int) bool {
	// We want Pop to give us the lowest weight node
	return pq[i].weight < pq[j].weight
}

func (pq pqGraph) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pqGraph) Push(x interface{}) {
	n := len(*pq)
	node := x.(*pqGraphNode)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *pqGraph) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

// update modifies the priority and value of an Item in the queue.
func (pq *pqGraph) update(item *pqGraphNode, priority int) {
	item.weight = priority
	heap.Fix(pq, item.index)
}
