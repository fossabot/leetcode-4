/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 *
 * https://leetcode.com/problems/search-a-2d-matrix-ii/description/
 *
 * algorithms
 * Medium (42.11%)
 * Likes:    2608
 * Dislikes: 71
 * Total Accepted:    279.8K
 * Total Submissions: 657.7K
 * Testcase Example:  '[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n5'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 *
 *
 * Integers in each row are sorted in ascending from left to right.
 * Integers in each column are sorted in ascending from top to bottom.
 *
 *
 * Example:
 *
 * Consider the following matrix:
 *
 *
 * [
 * ⁠ [1,   4,  7, 11, 15],
 * ⁠ [2,   5,  8, 12, 19],
 * ⁠ [3,   6,  9, 16, 22],
 * ⁠ [10, 13, 14, 17, 24],
 * ⁠ [18, 21, 23, 26, 30]
 * ]
 *
 *
 * Given target = 5, return true.
 *
 * Given target = 20, return false.
 *
 */

package leetcode

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	for _, row := range matrix {
		if row[0] > target {
			return false
		}
		res := binarySearch(row, target, n)
		if res == true {
			return res
		}
	}
	return false
}

func binarySearch(row []int, target, length int) bool {
	for left, right := 0, length-1; left <= right; {
		mid := (left + right) / 2
		if row[mid] == target {
			return true
		} else if row[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// @lc code=end
