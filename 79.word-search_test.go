/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (33.10%)
 * Likes:    2886
 * Dislikes: 148
 * Total Accepted:    402.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given a 2D board and a word, find if the word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cell,
 * where "adjacent" cells are those horizontally or vertically neighboring. The
 * same letter cell may not be used more than once.
 *
 * Example:
 *
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 *
 *
 *
 * Constraints:
 *
 *
 * board and word consists only of lowercase and uppercase English letters.
 * 1 <= board.length <= 200
 * 1 <= board[i].length <= 200
 * 1 <= word.length <= 10^3
 *
 *
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
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
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byteStack_Push(t *testing.T) {
	type args struct {
		v byte
	}
	tests := []struct {
		name string
		s    byteStack
		args args
		want byteStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Push(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("byteStack.Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byteStack_Pop(t *testing.T) {
	tests := []struct {
		name  string
		s     byteStack
		want  byteStack
		want1 byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Pop()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("byteStack.Pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("byteStack.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_byteStack_Len(t *testing.T) {
	tests := []struct {
		name string
		s    byteStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("byteStack.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byteStack_isEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    byteStack
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.isEmpty(); got != tt.want {
				t.Errorf("byteStack.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
