/*
 * @lc app=leetcode id=779 lang=golang
 *
 * [779] K-th Symbol in Grammar
 *
 * https://leetcode.com/problems/k-th-symbol-in-grammar/description/
 *
 * algorithms
 * Medium (36.97%)
 * Likes:    346
 * Dislikes: 121
 * Total Accepted:    25.7K
 * Total Submissions: 69.6K
 * Testcase Example:  '1\n1'
 *
 * On the first row, we write a 0. Now in every subsequent row, we look at the
 * previous row and replace each occurrence of 0 with 01, and each occurrence
 * of 1 with 10.
 *
 * Given row N and index K, return the K-th indexed symbol in row N. (The
 * values of K are 1-indexed.) (1 indexed).
 *
 *
 * Examples:
 * Input: N = 1, K = 1
 * Output: 0
 *
 * Input: N = 2, K = 1
 * Output: 0
 *
 * Input: N = 2, K = 2
 * Output: 1
 *
 * Input: N = 4, K = 5
 * Output: 1
 *
 * Explanation:
 * row 1: 0
 * row 2: 01
 * row 3: 0110
 * row 4: 01101001
 *
 *
 * Note:
 *
 *
 * N will be an integer in the range [1, 30].
 * K will be an integer in the range [1, 2^(N-1)].
 *
 *
 */

package leetcode

import (
	"testing"
)

func Test_kthGrammar(t *testing.T) {
	type args struct {
		N int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{1, 1},
			want: 0,
		},
		{
			name: "Test 2",
			args: args{2, 1},
			want: 0,
		},
		{
			name: "Test 3",
			args: args{2, 2},
			want: 1,
		},
		{
			name: "Test 4",
			args: args{4, 5},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthGrammar(tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("kthGrammar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inverse(t *testing.T) {
	type args struct {
		n int
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
			if got := inverse(tt.args.n); got != tt.want {
				t.Errorf("inverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
