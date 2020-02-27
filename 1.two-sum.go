/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (44.85%)
 * Likes:    12797
 * Dislikes: 453
 * Total Accepted:    2.4M
 * Total Submissions: 5.3M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 */

package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	solution := make([]int, 0)
	numMap := make(map[int]int)
	for i, num := range nums {
		numMap[target-num] = i
	}

	for i, num := range nums {
		if val, ok := numMap[num]; ok && val != i {
			return []int{i, val}
		}
	}
	return solution
}

// @lc code=end
