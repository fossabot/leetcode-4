/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 *
 * https://leetcode.com/problems/single-number/description/
 *
 * algorithms
 * Easy (62.24%)
 * Likes:    3712
 * Dislikes: 145
 * Total Accepted:    679.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '[2,2,1]'
 *
 * Given a non-empty array of integers, every element appears twice except for
 * one. Find that single one.
 *
 * Note:
 *
 * Your algorithm should have a linear runtime complexity. Could you implement
 * it without using extra memory?
 *
 * Example 1:
 *
 *
 * Input: [2,2,1]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: [4,1,2,1,2]
 * Output: 4
 *
 *
 */

package leetcode

// @lc code=start
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// @lc code=end
