/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 *
 * https://leetcode.com/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (62.72%)
 * Likes:    309
 * Dislikes: 213
 * Total Accepted:    59.1K
 * Total Submissions: 94.1K
 * Testcase Example:  '[5,2,3,1]'
 *
 * Given an array of integers nums, sort the array in ascending order.
 *
 *
 * Example 1:
 * Input: nums = [5,2,3,1]
 * Output: [1,2,3,5]
 * Example 2:
 * Input: nums = [5,1,1,2,0,0]
 * Output: [0,0,1,1,2,5]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 50000
 * -50000 <= nums[i] <= 50000
 *
 *
 */

package leetcode

// @lc code=start
func sortArray(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	mid := length / 2
	leftArray := mergeSort(nums[0:mid])
	rightArray := mergeSort(nums[mid:])

	result := make([]int, length)
	for i, left, right := 0, 0, 0; i < length; i++ {
		if right == (length-mid) ||
			(left != mid && leftArray[left] <= rightArray[right]) {
			result[i] = leftArray[left]
			left++
		} else {
			result[i] = rightArray[right]
			right++
		}
	}
	return result
}

// @lc code=end
