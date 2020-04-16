/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (52.14%)
 * Likes:    3220
 * Dislikes: 102
 * Total Accepted:    490.1K
 * Total Submissions: 906.7K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given a set of candidate numbers (candidates) (without duplicates) and a
 * target number (target), find all unique combinations in candidates where the
 * candidate numbers sums to target.
 *
 * The same repeated number may be chosen from candidates unlimited number of
 * times.
 *
 * Note:
 *
 *
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [2,3,6,7], target = 7,
 * A solution set is:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8,
 * A solution set is:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 */

package leetcode

import "sort"

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return genCombSum(candidates, target)
}

func genCombSum(nums []int, target int) [][]int {
	solution := [][]int{}
	for i, num := range nums {
		if num > target {
			break
		} else if num < target {
			tmp := genCombSum(nums[i:], target-num)
			for _, sol := range tmp {
				sol = append(sol, num)
				solution = append(solution, sol)
			}
		} else {
			solution = append(solution, []int{num})
			break
		}
	}
	return solution
}

// @lc code=end
