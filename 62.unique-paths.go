/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 *
 * https://leetcode.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (50.38%)
 * Likes:    2664
 * Dislikes: 187
 * Total Accepted:    419K
 * Total Submissions: 807.4K
 * Testcase Example:  '3\n2'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 *
 * How many possible unique paths are there?
 *
 *
 * Above is a 7 x 3 grid. How many possible unique paths are there?
 *
 *
 * Example 1:
 *
 *
 * Input: m = 3, n = 2
 * Output: 3
 * Explanation:
 * From the top-left corner, there are a total of 3 ways to reach the
 * bottom-right corner:
 * 1. Right -> Right -> Down
 * 2. Right -> Down -> Right
 * 3. Down -> Right -> Right
 *
 *
 * Example 2:
 *
 *
 * Input: m = 7, n = 3
 * Output: 28
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= m, n <= 100
 * It's guaranteed that the answer will be less than or equal to 2 * 10 ^ 9.
 *
 *
 */

package leetcode

// @lc code=start
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	path := make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				path[i][j] = 1
			} else {
				path[i][j] = (path[i-1][j] + path[i][j-1])
			}
		}
	}
	return path[m-1][n-1]
}

// @lc code=end
