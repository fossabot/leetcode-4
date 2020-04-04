/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 *
 * https://leetcode.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (55.79%)
 * Likes:    3181
 * Dislikes: 107
 * Total Accepted:    665.5K
 * Total Submissions: 1.2M
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * Given an array nums, write a function to move all 0's to the end of it while
 * maintaining the relative order of the non-zero elements.
 *
 * Example:
 *
 *
 * Input: [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 *
 * Note:
 *
 *
 * You must do this in-place without making a copy of the array.
 * Minimize the total number of operations.
 *
 */

package leetcode

// @lc code=start
func moveZeroes(nums []int) {
	numLen := len(nums)
	// nZ refers to the position of next non-zero element
	for nZ, i := 0, 0; i < numLen; i++ {
		if nums[i] != 0 {
			nums[i], nums[nZ] = nums[nZ], nums[i]
			nZ++
		}
	}
}

// @lc code=end
