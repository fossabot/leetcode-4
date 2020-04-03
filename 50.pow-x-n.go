/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (29.18%)
 * Likes:    1231
 * Dislikes: 2738
 * Total Accepted:    422.3K
 * Total Submissions: 1.4M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 *
 * Example 1:
 *
 *
 * Input: 2.00000, 10
 * Output: 1024.00000
 *
 *
 * Example 2:
 *
 *
 * Input: 2.10000, 3
 * Output: 9.26100
 *
 *
 * Example 3:
 *
 *
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 * Note:
 *
 *
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 *
 *
 */

package leetcode

// @lc code=start
func myPow(x float64, n int) float64 {
	res := 1.0
	if n < 0 {
		x = (1 / x)
		n = -n
	}
	for i := n; i > 0; i /= 2 {
		if i&1 == 1 {
			res *= x
		}
		x *= x
	}
	return res
}

func myPowTail(cur, x float64, n int) float64 {
	if n == 0 {
		return cur
	} else if n < 0 {
		return myPowTail(cur/x, x, n+1)
	}
	return myPowTail(cur*x, x, n-1)
}

// @lc code=end
