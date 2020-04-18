/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (49.83%)
 * Likes:    2389
 * Dislikes: 52
 * Total Accepted:    341.4K
 * Total Submissions: 656.6K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right which minimizes the sum of all numbers along its path.
 *
 * Note: You can only move either down or right at any point in time.
 *
 * Example:
 *
 *
 * Input:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * Output: 7
 * Explanation: Because the path 1→3→1→1→1 minimizes the sum.
 *
 *
 */

package leetcode

import "container/heap"

// @lc code=start
func minPathSum(grid [][]int) int {
	return dp(grid)
}

func dp(grid [][]int) int {
	X, Y := len(grid), len(grid[0])

	for i := 1; i < X; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < Y; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < X; i++ {
		for j := 1; j < Y; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[X-1][Y-1]
}

func dijkstra(grid [][]int) int {
	X, Y := len(grid), len(grid[0])
	visited := make([][]bool, X)
	for i := 0; i < X; i++ {
		visited[i] = make([]bool, Y)
	}

	target := [2]int{X - 1, Y - 1}
	path := 0
	h := &nodeHeap{[3]int{grid[0][0], 0, 0}}
	heap.Init(h)

	for {
		node := heap.Pop(h).([3]int)
		path = node[0]
		curPos := [2]int{node[1], node[2]}
		if curPos == target {
			break
		}
		if node[1] < (X - 1) {
			x, y := node[1]+1, node[2]
			if visited[x][y] == false {
				visited[x][y] = true
				down := [3]int{path + grid[x][y], x, y}
				heap.Push(h, down)
			}
		}
		if node[2] < (Y - 1) {
			x, y := node[1], node[2]+1
			if visited[x][y] == false {
				right := [3]int{path + grid[x][y], x, y}
				heap.Push(h, right)
			}
		}
	}
	return path
}

type nodeHeap [][3]int

func (h nodeHeap) Len() int {
	return len(h)
}

// Min-heap implementation
func (h nodeHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h nodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *nodeHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// @lc code=end
