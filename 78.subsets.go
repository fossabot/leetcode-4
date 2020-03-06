/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (57.70%)
 * Likes:    2936
 * Dislikes: 70
 * Total Accepted:    488.8K
 * Total Submissions: 842.2K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct integers, nums, return all possible subsets (the
 * power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: nums = [1,2,3]
 * Output:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

package leetcode

// @lc code=start
func subsets(nums []int) [][]int {
	result := [][]int{{}}
	for _, num := range nums {
		for _, set := range result {
			s := make([]int, len(set)+1)
			copy(s, set)
			s[len(set)] = num
			result = append(result, s)
		}
		// fmt.Println(result)
	}
	return result
}

// @lc code=end
