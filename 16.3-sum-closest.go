/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.72%)
 * Likes:    1654
 * Dislikes: 120
 * Total Accepted:    425.5K
 * Total Submissions: 930.7K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 *
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 */

package leetcode

import (
	"math"
	"sort"
)

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result, diff := 0, math.MaxInt64
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				val := nums[i] + nums[j] + nums[k]
				if val > target {
					if (val - target) < diff {
						diff = val - target
						result = val
					} else {
						break
					}
				} else {
					if (target - val) < diff {
						diff = target - val
						result = val
					}
				}
			}
		}
	}
	return result
}

// @lc code=end
