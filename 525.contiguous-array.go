/*
 * @lc app=leetcode id=525 lang=golang
 *
 * [525] Contiguous Array
 *
 * https://leetcode.com/problems/contiguous-array/description/
 *
 * algorithms
 * Medium (44.32%)
 * Likes:    1294
 * Dislikes: 63
 * Total Accepted:    71.4K
 * Total Submissions: 168.8K
 * Testcase Example:  '[0,1]'
 *
 * Given a binary array, find the maximum length of a contiguous subarray with
 * equal number of 0 and 1.
 *
 *
 * Example 1:
 *
 * Input: [0,1]
 * Output: 2
 * Explanation: [0, 1] is the longest contiguous subarray with equal number of
 * 0 and 1.
 *
 *
 *
 * Example 2:
 *
 * Input: [0,1,0]
 * Output: 2
 * Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal
 * number of 0 and 1.
 *
 *
 *
 * Note:
 * The length of the given binary array will not exceed 50,000.
 *
 */

package leetcode

// @lc code=start
func findMaxLength(nums []int) int {
	sumIndex := make(map[int]int)
	sumIndex[0] = 0
	max := 0
	for i, state := 0, 0; i < len(nums); i++ {
		if nums[i] == 0 {
			state--
		} else {
			state++
		}
		index, found := sumIndex[state]
		if !found {
			sumIndex[state] = (i + 1)
		} else if max < (i + 1 - index) {
			max = (i + 1 - index)
		}
	}
	return max
}

// @lc code=end
