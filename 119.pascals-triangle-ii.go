/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 *
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (46.85%)
 * Likes:    647
 * Dislikes: 188
 * Total Accepted:    253.5K
 * Total Submissions: 538.6K
 * Testcase Example:  '3'
 *
 * Given a non-negative index k where k ≤ 33, return the k^th index row of the
 * Pascal's triangle.
 *
 * Note that the row index starts from 0.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 3
 * Output: [1,3,3,1]
 *
 *
 * Follow up:
 *
 * Could you optimize your algorithm to use only O(k) extra space?
 *
 */

package leetcode

// @lc code=start
func getRow(rowIndex int) []int {

	// Iterative solution
	// result := make([][]int, 0)
	// result = append(result, []int{1})
	// for i := 1; i <= rowIndex; i++ {
	// 	temp := make([]int, 0)
	// 	for j := 0; j <= i; j++ {
	// 		val := 0
	// 		if j == 0 || j == i {
	// 			val = 1
	// 		} else {
	// 			val = result[i-1][j-1] + result[i-1][j]
	// 		}
	// 		temp = append(temp, val)
	// 	}
	// 	result = append(result, temp)
	// }
	// return result[rowIndex]

	// Recursive solution
	if rowIndex == 0 {
		return []int{1}
	}
	pre := getRow(rowIndex - 1)
	result := make([]int, rowIndex+1)
	result[0], result[rowIndex] = 1, 1
	for i := 1; i < rowIndex; i++ {
		result[i] = pre[i-1] + pre[i]
	}
	return result
}

// @lc code=end
