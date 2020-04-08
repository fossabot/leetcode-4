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

import (
	"reflect"
	"testing"
)

func Test_totalNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{4},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{10},
			want: 724,
		},
		{
			name: "Test 3",
			args: args{14},
			want: 365596,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.args.n); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeBoard(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeBoard(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSolve(t *testing.T) {
	type args struct {
		board [][]bool
		n     int
		row   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSolve(tt.args.board, tt.args.n, tt.args.row); got != tt.want {
				t.Errorf("countSolve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSafe(t *testing.T) {
	type args struct {
		board [][]bool
		n     int
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
			if got := isSafe(tt.args.board, tt.args.n, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("isSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_placeQueen(t *testing.T) {
	type args struct {
		board [][]bool
		n     int
		row   int
		col   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			placeQueen(tt.args.board, tt.args.n, tt.args.row, tt.args.col)
		})
	}
}

func Test_removeQueen(t *testing.T) {
	type args struct {
		board [][]bool
		n     int
		row   int
		col   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removeQueen(tt.args.board, tt.args.n, tt.args.row, tt.args.col)
		})
	}
}
