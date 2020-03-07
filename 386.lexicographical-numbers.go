/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 *
 * https://leetcode.com/problems/lexicographical-numbers/description/
 *
 * algorithms
 * Medium (49.34%)
 * Likes:    496
 * Dislikes: 66
 * Total Accepted:    49.1K
 * Total Submissions: 99K
 * Testcase Example:  '13'
 *
 * Given an integer n, return 1 - n in lexicographical order.
 *
 * For example, given 13, return: [1,10,11,12,13,2,3,4,5,6,7,8,9].
 *
 * Please optimize your algorithm to use less time and space. The input size
 * may be as large as 5,000,000.
 *
 */

package leetcode

// @lc code=start
func lexicalOrder(n int) []int {
	sol := []int{}
	var limit int
	if n < 9 {
		limit = n
	} else {
		limit = 9
	}
	for i := 1; i <= limit; i++ {
		sol = append(sol, i)
		nextStart := (i * 10)
		nextStop := (nextStart + 10)
		for nextStart <= n {
			for j := nextStart; j < nextStop && j <= n; j++ {
				sol = append(sol, j)
			}
			nextStart *= 10
			nextStop = (nextStart * 2)
		}
	}
	return sol
}

// @lc code=end
