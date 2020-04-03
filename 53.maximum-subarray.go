/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (45.22%)
 * Likes:    5842
 * Dislikes: 245
 * Total Accepted:    722.7K
 * Total Submissions: 1.6M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */

package leetcode

// @lc code=start
func maxSubArray(nums []int) int {
	numLen := len(nums)
	if numLen == 0 {
		return 0
	}

	// SumInc holds the max sum including last val
	// gSum holds the global sum
	sumInc, gSum := nums[0], nums[0]
	for i := 1; i < numLen; i++ {
		sumInc = max(sumInc+nums[i], nums[i])
		gSum = max(gSum, sumInc)
	}
	return gSum
}

// @lc code=end
