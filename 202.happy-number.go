/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 *
 * https://leetcode.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (47.70%)
 * Likes:    1603
 * Dislikes: 357
 * Total Accepted:    361.8K
 * Total Submissions: 734.2K
 * Testcase Example:  '19'
 *
 * Write an algorithm to determine if a number is "happy".
 *
 * A happy number is a number defined by the following process: Starting with
 * any positive integer, replace the number by the sum of the squares of its
 * digits, and repeat the process until the number equals 1 (where it will
 * stay), or it loops endlessly in a cycle which does not include 1. Those
 * numbers for which this process ends in 1 are happy numbers.
 *
 * Example:
 *
 *
 * Input: 19
 * Output: true
 * Explanation:
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 *
 */

package leetcode

// @lc code=start
func isHappy(n int) bool {
	for slow, fast := n, n; ; {
		slow = digitSquareSum(slow)
		fast = digitSquareSum(digitSquareSum(fast))
		if fast == 1 {
			return true
		} else if slow == fast {
			return false
		}
	}
}

func digitSquareSum(n int) (sum int) {
	for ; n > 0; n /= 10 {
		tmp := n % 10
		sum += (tmp * tmp)
	}
	return sum
}

// @lc code=end
