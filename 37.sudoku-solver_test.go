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

import (
	"reflect"
	"testing"
)

func Test_solveSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Test 1",
			args: args{[][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {
				'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {
				'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {
				'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {
				'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {
				'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {
				'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {
				'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {
				'.', '.', '.', '.', '8', '.', '.', '7', '9'}}},
			want: [][]byte{{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'}, {'1', '9', '8',
					'3', '4', '2', '5', '6', '7'}, {'8', '5', '9', '7', '6', '1',
					'4', '2', '3'}, {'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'}, {'9', '6', '1',
					'5', '3', '7', '2', '8', '4'}, {'2', '8', '7', '4', '1', '9',
					'6', '3', '5'}, {'3', '4', '5', '2', '8', '6', '1', '7', '9'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.args.board)
			if !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("\nsolveSudokuPos()\n%v\nwant\n%v", tt.args.board, tt.want)
			}
		})
	}
}

func Test_solveSudokuPos(t *testing.T) {
	type args struct {
		board [][]byte
		row   int
		col   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSudokuPos(tt.args.board, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("solveSudokuPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isVal(t *testing.T) {
	type args struct {
		board [][]byte
		row   int
		col   int
		val   byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidValue(tt.args.board, tt.args.val, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("isVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
