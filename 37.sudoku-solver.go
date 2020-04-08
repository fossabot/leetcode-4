/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 *
 * https://leetcode.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (39.95%)
 * Likes:    1470
 * Dislikes: 84
 * Total Accepted:    170.1K
 * Total Submissions: 410.4K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Write a program to solve a Sudoku puzzle by filling the empty cells.
 *
 * A sudoku solution must satisfy all of the following rules:
 *
 *
 * Each of the digits 1-9 must occur exactly once in each row.
 * Each of the digits 1-9 must occur exactly once in each column.
 * Each of the the digits 1-9 must occur exactly once in each of the 9 3x3
 * sub-boxes of the grid.
 *
 *
 * Empty cells are indicated by the character '.'.
 *
 *
 * A sudoku puzzle...
 *
 *
 * ...and its solution numbers marked in red.
 *
 * Note:
 *
 *
 * The given board contain only digits 1-9 and the character '.'.
 * You may assume that the given Sudoku puzzle will have a single unique
 * solution.
 * The given board size is always 9x9.
 *
 *
 */

package leetcode

// @lc code=start
func solveSudoku(board [][]byte) {
	solveSudokuPos(board, 0, 0)
}

func solveSudokuPos(board [][]byte, row, col int) bool {
	solved, row, col := getNextPos(board, row, col)
	if solved {
		return solved
	}

	for i := byte('1'); i <= '9'; i++ {
		if isValidValue(board, i, row, col) {
			putValue(board, i, row, col)
			solved := solveSudokuPos(board, row, col+1)
			if solved {
				return solved
			}
			removeValue(board, row, col)
		}
	}
	return false
}

func getNextPos(board [][]byte, row, col int) (bool, int, int) {
	for row < 9 {
		for col < 9 {
			if board[row][col] == '.' {
				return false, row, col
			}
			col++
		}
		row++
		col = 0
	}
	return true, row, col
}

func removeValue(board [][]byte, row, col int) {
	board[row][col] = '.'
}

func putValue(board [][]byte, val byte, row, col int) {
	board[row][col] = val
}

func isValidValue(board [][]byte, val byte, row, col int) bool {
	// Check for value in row
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
	}

	// Check for value in column
	for i := 0; i < 9; i++ {
		if board[i][col] == val {
			return false
		}
	}

	// Check for value in box
	rowStart, rowEnd := (row/3)*3, (row/3+1)*3
	colStart, colEnd := (col/3)*3, (col/3+1)*3
	for i := rowStart; i < rowEnd; i++ {
		for j := colStart; j < colEnd; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}

	return true
}

// @lc code=end
