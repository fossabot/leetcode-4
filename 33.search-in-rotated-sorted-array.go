/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (33.45%)
 * Likes:    4162
 * Dislikes: 419
 * Total Accepted:    625.9K
 * Total Submissions: 1.8M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 *
 * You are given a target value to search. If found in the array return its
 * index, otherwise return -1.
 *
 * You may assume no duplicate exists in the array.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * Example 1:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 *
 */

package leetcode

// @lc code=start
func search(nums []int, target int) int {
	numLen := len(nums)
	if numLen == 0 {
		return -1
	}

	left, right := 0, numLen-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[right] >= nums[mid] {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target <= nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

// @lc code=end
