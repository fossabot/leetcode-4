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

// @lc code=start
func networkDelayTime(times [][]int, N int, K int) int {
	graph := make([][][]int, N)
	for _, edge := range times {
		// Changing index from 1 to N -> 0 to N-1
		edge := []int{edge[0] - 1, edge[1] - 1, edge[2]}
		graph[edge[0]] = append(graph[edge[0]], edge[1:])
	}
	visited := make(map[int]bool)
	visitTime := make([]int, N)
	for vertex, edges := range graph {
		for _, edge := range edges {
			curVisTime
		}
	}
	// return 0
}

// @lc code=end
