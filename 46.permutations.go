/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (60.04%)
 * Likes:    3155
 * Dislikes: 94
 * Total Accepted:    525.9K
 * Total Submissions: 868.4K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a collection of distinct integers, return all possible permutations.
 *
 * Example:
 *
 *
 * Input: [1,2,3]
 * Output:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 *
 */

package leetcode

// @lc code=start
func permute(nums []int) [][]int {
	result := [][]int{{}}
	for _, num := range nums {
		// New slice to hold results from current step
		curResult := [][]int{}
		// Get all the permutations in current result
		for _, per := range result {
			for i := 0; i < len(per)+1; i++ {
				// Make a slice one element bigger
				item := make([]int, len(per)+1)
				item[i] = num
				copy(item[:i], per[:i])
				copy(item[i+1:], per[i:])
				curResult = append(curResult, item)
			}
		}
		// Make current result the new result
		result = curResult
	}
	return result
}

// @lc code=end
