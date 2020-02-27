/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (25.66%)
 * Likes:    5664
 * Dislikes: 688
 * Total Accepted:    783.4K
 * Total Submissions: 3M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 *
 * Note:
 *
 * The solution set must not contain duplicate triplets.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 *
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 */

package leetcode

import (
	"sort"
)

// @lc code=start

func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return [][]int{}
	}

	numMap := make(map[int]int)
	sort.Ints(nums)

	for i, num := range nums {
		numMap[num] = i
	}

	solution := make(map[[3]int]bool, 0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			twoSum := nums[i] + nums[j]
			checkVal := -twoSum
			if index, ok := numMap[checkVal]; ok && index > i && index > j {
				solution[[...]int{nums[i], nums[j], checkVal}] = true
			}
		}
	}

	sol := make([][]int, 0)
	for array := range solution {
		// Golang copies array to slice as address. This assignement is only for
		// creating new address
		arr := array
		sol = append(sol, arr[:])
	}
	return sol
}

// @lc code=end
