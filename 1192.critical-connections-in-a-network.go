/*
 * @lc app=leetcode id=1192 lang=golang
 *
 * [1192] Critical Connections in a Network
 *
 * https://leetcode.com/problems/critical-connections-in-a-network/description/
 *
 * algorithms
 * Hard (48.23%)
 * Likes:    599
 * Dislikes: 57
 * Total Accepted:    29.6K
 * Total Submissions: 61.1K
 * Testcase Example:  '4\n[[0,1],[1,2],[2,0],[1,3]]'
 *
 * There are n servers numbered from 0 to n-1 connected by undirected
 * server-to-server connections forming a network where connections[i] = [a, b]
 * represents a connection between servers a and b. Any server can reach any
 * other server directly or indirectly through the network.
 *
 * A critical connection is a connection that, if removed, will make some
 * server unable to reach some other server.
 *
 * Return all critical connections in the network in any order.
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
 * Output: [[1,3]]
 * Explanation: [[3,1]] is also accepted.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 10^5
 * n-1 <= connections.length <= 10^5
 * connections[i][0] != connections[i][1]
 * There are no repeated connections.
 *
 *
 */

package leetcode

import "sort"

// @lc code=start
// func criticalConnections(n int, connections [][]int) [][]int {
// 	var criticalConns [][]int
// 	disc := make([]int, n)
// 	low := make([]int, n)
// 	for _, connection := range connections {

// 	}
// 	return criticalConns
// }

// @lc code=end
