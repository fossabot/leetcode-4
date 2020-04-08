/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (51.10%)
 * Likes:    1230
 * Dislikes: 62
 * Total Accepted:    265.9K
 * Total Submissions: 505K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 *
 * Example:
 *
 *
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 *
 */

package leetcode

// @lc code=start
func combine(n int, k int) [][]int {
	if k == 0 {
		return [][]int{{}}
	}
	res := [][]int{}
	for ; n >= k; n-- {
		tmp := []int{n}
		tmpRes := combine(n-1, k-1)
		for _, comb := range tmpRes {
			res = append(res, append(tmp, comb...))
		}
	}
	return res
}

// @lc code=end
