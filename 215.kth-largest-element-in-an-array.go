/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (51.31%)
 * Likes:    3160
 * Dislikes: 223
 * Total Accepted:    555.3K
 * Total Submissions: 1M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Find the kth largest element in an unsorted array. Note that it is the kth
 * largest element in the sorted order, not the kth distinct element.
 *
 * Example 1:
 *
 *
 * Input: [3,2,1,5,6,4] and k = 2
 * Output: 5
 *
 *
 * Example 2:
 *
 *
 * Input: [3,2,3,1,2,4,5,5,6] and k = 4
 * Output: 4
 *
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ array's length.
 *
 */

package leetcode

// @lc code=start
func findKthLargest(nums []int, k int) int {
	j := part(nums)

	// Assuming k is always valid
	if j == (k - 1) {
		return nums[j]
	} else if j > (k - 1) {
		return findKthLargest(nums[:j], k)
	}
	return findKthLargest(nums[j+1:], k-j-1)
}

// Part partitions the array around a pivot and returns the index
func part(nums []int) int {
	arrLen := len(nums)
	if arrLen == 1 {
		return 0
	}
	j := 0
	for pivot, i := nums[0], 1; i < arrLen; i++ {
		if nums[i] > pivot {
			j++
			swap(nums, i, j)
		}
	}
	swap(nums, 0, j)
	return j
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// @lc code=end
