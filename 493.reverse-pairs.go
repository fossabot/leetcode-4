/*
 * @lc app=leetcode id=493 lang=golang
 *
 * [493] Reverse Pairs
 *
 * https://leetcode.com/problems/reverse-pairs/description/
 *
 * algorithms
 * Hard (24.11%)
 * Likes:    737
 * Dislikes: 108
 * Total Accepted:    33.6K
 * Total Submissions: 137.6K
 * Testcase Example:  '[1,3,2,3,1]'
 *
 * Given an array nums, we call (i, j) an important reverse pair if i < j and
 * nums[i] > 2*nums[j].
 *
 * You need to return the number of important reverse pairs in the given
 * array.
 *
 * Example1:
 *
 * Input: [1,3,2,3,1]
 * Output: 2
 *
 *
 * Example2:
 *
 * Input: [2,4,3,5,1]
 * Output: 3
 *
 *
 * Note:
 *
 * The length of the given array will not exceed 50,000.
 * All the numbers in the input array are in the range of 32-bit integer.
 *
 *
 */

package leetcode

import "sort"

// @lc code=start
func reversePairs(nums []int) int {
	return indexedReversePairs(nums, 0, len(nums)-1)
}

func indexedReversePairs(nums []int, s, e int) int {
	if s >= e {
		return 0
	}
	mid := s + (e-s)/2
	cnt := indexedReversePairs(nums, s, mid) + indexedReversePairs(nums, mid+1, e)
	for i, j := s, mid+1; i <= mid; i++ {
		for j <= e && nums[i] > nums[j]*2 {
			j++
		}
		cnt += j - (mid + 1)
	}
	sort.Ints(nums[s : e+1])
	return cnt
}

// @lc code=end
