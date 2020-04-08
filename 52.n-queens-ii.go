/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 *
 * https://leetcode.com/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (54.77%)
 * Likes:    415
 * Dislikes: 138
 * Total Accepted:    124K
 * Total Submissions: 221.2K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 * such that no two queens attack each other.
 *
 *
 *
 * Given an integer n, return the number of distinct solutions to the n-queens
 * puzzle.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: 2
 * Explanation: There are two distinct solutions to the 4-queens puzzle as
 * shown below.
 * [
 * [".Q..",  // Solution 1
 * "...Q",
 * "Q...",
 * "..Q."],
 *
 * ["..Q.",  // Solution 2
 * "Q...",
 * "...Q",
 * ".Q.."]
 * ]
 *
 *
 */

package leetcode

// @lc code=start
func totalNQueens(n int) int {
	board := makeBoard(n)
	return backtrack(board, n, 0)
}

func makeBoard(n int) [][]bool {
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	return board
}

func backtrack(board [][]bool, n, row int) int {
	count := 0
	for col := 0; col < n; col++ {
		if isSafe(board, n, row, col) {
			placeQueen(board, n, row, col)
			if row == (n - 1) {
				count++
			} else {
				count += backtrack(board, n, row+1)
			}
			removeQueen(board, n, row, col)
		}
	}
	return count
}

func isSafe(board [][]bool, n, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] {
			return false
		}
	}

	// left diagonal
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}

	// right diagonal
	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] {
			return false
		}
	}
	return true
}

func placeQueen(board [][]bool, n, row, col int) {
	board[row][col] = true
}

func removeQueen(board [][]bool, n, row, col int) {
	board[row][col] = false
}

// @lc code=end
