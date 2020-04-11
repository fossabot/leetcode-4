/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (34.71%)
 * Likes:    2843
 * Dislikes: 127
 * Total Accepted:    447K
 * Total Submissions: 1.3M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * If the target is not found in the array, return [-1, -1].
 *
 * Example 1:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 *
 * Example 2:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 *
 */

package leetcode

// @lc code=start
func searchRange(nums []int, target int) []int {
	rng := bSearch(nums, target)
	if rng[0] == -1 {
		return rng
	}

	return findRange(nums, rng, target)
}

func bSearch(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return []int{-1, -1}
	}

	for left, right := 0, length-1; left <= right; {
		mid := (left + right) / 2
		if nums[mid] == target {
			return []int{left, right}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}

func findRange(nums, limits []int, target int) []int {
	left, right := limits[0], limits[1]
	mid := (left + right) / 2

	for midL, midR := (left+mid)/2, (mid+right)/2; ; {
		if midL == target {
			left = midL
		}
	}
}

// @lc code=end
