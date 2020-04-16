/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (57.75%)
 * Likes:    4042
 * Dislikes: 346
 * Total Accepted:    432.8K
 * Total Submissions: 732.4K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an array nums of n integers where n > 1,  return an array output such
 * that output[i] is equal to the product of all the elements of nums except
 * nums[i].
 *
 * Example:
 *
 *
 * Input:  [1,2,3,4]
 * Output: [24,12,8,6]
 *
 *
 * Constraint: It's guaranteed that the product of the elements of any prefix
 * or suffix of the array (including the whole array) fits in a 32 bit
 * integer.
 *
 * Note: Please solve it without division and in O(n).
 *
 * Follow up:
 * Could you solve it with constant space complexity? (The output array does
 * not count as extra space for the purpose of space complexity analysis.)
 *
 */

package leetcode

// @lc code=start
func productExceptSelf(nums []int) []int {
	N := len(nums)
	out := make([]int, N)
	out[0] = 1
	for i := 1; i < N; i++ {
		out[i] = (out[i-1] * nums[i-1])
	}

	for i, right := N-1, 1; i >= 0; i-- {
		out[i] *= right
		right *= nums[i]
	}

	return out
}

// @lc code=end
