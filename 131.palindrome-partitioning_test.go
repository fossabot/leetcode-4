/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 *
 * https://leetcode.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (43.84%)
 * Likes:    1502
 * Dislikes: 58
 * Total Accepted:    209K
 * Total Submissions: 461.2K
 * Testcase Example:  '"aab"'
 *
 * Given a string s, partition s such that every substring of the partition is
 * a palindrome.
 *
 * Return all possible palindrome partitioning of s.
 *
 * Example:
 *
 *
 * Input: "aab"
 * Output:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 *
 *
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test 1",
			args: args{"aab"},
			want: [][]string{{"aa", "b"}, {"a", "a", "b"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatePartition(t *testing.T) {
	type args struct {
		s     string
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePartition(tt.args.s, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
