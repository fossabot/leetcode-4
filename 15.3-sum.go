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

	solution := make([][]int, 0)
	for i, num1 := range nums[:len(nums)-1] {
		twoSumSol := twoSum(nums[i+1:], -num1)
		solution = append(solution, []int{num1, nums[i+twoSumSol[0]], nums[twoSumSol[1]]})
		// smallSlice := nums[i+1:]
		// for j, num2 := range smallSlice {
		// 	twoSum := num1 + num2
		// 	checkVal := -twoSum
		// 	if index, ok := numMap[checkVal]; ok && index > i && index > (j+i+1) {
		// 		if j > 0 && smallSlice[j] == smallSlice[j-1] {
		// 			continue
		// 		}
		// 		fmt.Println(i, i+j+1, index)
		// 		fmt.Println(num1, num2, checkVal)
		// 		solution = append(solution, []int{num1, num2, checkVal})
		// 	}
		// }
	}
	return solution
	// sort.Ints(nums)
	// var solus = make([][]int, 0)
	// for lo := 0; lo < len(nums)-2; lo++ {
	// 	mid, hi, target := lo+1, len(nums)-1, -nums[lo]
	// 	for mid < hi {
	// 		sum := nums[mid] + nums[hi]
	// 		switch {
	// 		case sum < target:
	// 			mid++
	// 		case sum > target:
	// 			hi--
	// 		default: // find one
	// 			solus = append(solus, []int{nums[lo], nums[mid], nums[hi]})
	// 			for mid+1 < hi && nums[mid] == nums[mid+1] { // move next to the right-most same number
	// 				mid++
	// 			}
	// 			for hi-1 > mid && nums[hi] == nums[hi-1] { // move back to the left-most same number
	// 				hi--
	// 			}
	// 			mid++
	// 			hi--
	// 		}
	// 	}
	// 	for lo+1 < len(nums) && nums[lo] == nums[lo+1] { // move next to the right-most same number
	// 		lo++
	// 	}
	// }
	// return solus
}

// @lc code=end
