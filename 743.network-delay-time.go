/*
 * @lc app=leetcode id=743 lang=golang
 *
 * [743] Network Delay Time
 *
 * https://leetcode.com/problems/network-delay-time/description/
 *
 * algorithms
 * Medium (44.84%)
 * Likes:    1045
 * Dislikes: 202
 * Total Accepted:    66.3K
 * Total Submissions: 146.6K
 * Testcase Example:  '[[2,1,1],[2,3,1],[3,4,1]]\n4\n2'
 *
 * There are N network nodes, labelled 1 to N.
 *
 * Given times, a list of travel times as directed edges times[i] = (u, v, w),
 * where u is the source node, v is the target node, and w is the time it takes
 * for a signal to travel from source to target.
 *
 * Now, we send a signal from a certain node K. How long will it take for all
 * nodes to receive the signal? If it is impossible, return -1.
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: times = [[2,1,1],[2,3,1],[3,4,1]], N = 4, K = 2
 * Output: 2
 *
 *
 *
 *
 * Note:
 *
 *
 * N will be in the range [1, 100].
 * K will be in the range [1, N].
 * The length of times will be in the range [1, 6000].
 * All edges times[i] = (u, v, w) will have 1 <= u, v <= N and 0 <= w <= 100.
 *
 *
 */

package leetcode

import (
	"container/heap"
)

// @lc code=start
func networkDelayTime(times [][]int, N int, K int) int {
	// Converting the graph to slice of slice consisting of dest and weights
	network := getGraph(times, N)

	// Slice of nodes in the graph. This will hold all the nodes in graph
	nodeSlice := make([]*pqGraphNode, N)
	nodeSlice[K-1] = &pqGraphNode{dest: K - 1,
		weight: 0,
		index:  0}
	//Priority queue for all the nodes in graph
	pq := make(pqGraph, 1)
	pq[0] = nodeSlice[K-1]
	heap.Init(&pq)

	visited, visitTime := 0, 0

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*pqGraphNode)
		source, time := node.dest, node.weight
		visited++
		visitTime = time
		for _, edge := range network[source] {
			dest, delay := edge[0], edge[1]+time
			newNode := nodeSlice[dest]
			if newNode == nil {
				nodeSlice[dest] = &pqGraphNode{dest: dest,
					weight: delay}
				heap.Push(&pq, nodeSlice[dest])
			} else if newNode.weight > delay {
				newNode.weight = delay
				pq.update(newNode, delay)
			}
		}
	}
	if visited == N {
		return visitTime
	}
	return -1
}

// @lc code=end
