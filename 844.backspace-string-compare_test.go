/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 *
 * https://leetcode.com/problems/backspace-string-compare/description/
 *
 * algorithms
 * Easy (46.87%)
 * Likes:    1194
 * Dislikes: 65
 * Total Accepted:    126.6K
 * Total Submissions: 267.3K
 * Testcase Example:  '"ab#c"\n"ad#c"'
 *
 * Given two strings S and T, return if they are equal when both are typed into
 * empty text editors. # means a backspace character.
 *
 *
 * Example 1:
 *
 *
 * Input: S = "ab#c", T = "ad#c"
 * Output: true
 * Explanation: Both S and T become "ac".
 *
 *
 *
 * Example 2:
 *
 *
 * Input: S = "ab##", T = "c#d#"
 * Output: true
 * Explanation: Both S and T become "".
 *
 *
 *
 * Example 3:
 *
 *
 * Input: S = "a##c", T = "#a#c"
 * Output: true
 * Explanation: Both S and T become "c".
 *
 *
 *
 * Example 4:
 *
 *
 * Input: S = "a#c", T = "b"
 * Output: false
 * Explanation: S becomes "c" while T becomes "b".
 *
 *
 * Note:
 *
 *
 * 1 <= S.length <= 200
 * 1 <= T.length <= 200
 * S and T only contain lowercase letters and '#' characters.
 *
 *
 * Follow up:
 *
 *
 * Can you solve it in O(N) time and O(1) space?
 *
 *
 *
 *
 *
 *
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{"a##c", "#a#c"},
			want: true,
		},
		{
			name: "Test 2",
			args: args{"ab##", "c#d#"},
			want: true,
		},
		{
			name: "Test 3",
			args: args{"a#c", "b"},
			want: false,
		},
		{
			name: "Test 4",
			args: args{"ab#c", "ad#c"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backspaceStr(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceStr(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backspaceStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkByte(t *testing.T) {
	type args struct {
		lastByte byte
		res      []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkByte(tt.args.lastByte, tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backspaceCompareIterative(t *testing.T) {
	type args struct {
		S string
		T string
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
			if got := backspaceCompareIterative(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompareIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextPos(t *testing.T) {
	type args struct {
		S   string
		pos int
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
			if got := nextPos(tt.args.S, tt.args.pos); got != tt.want {
				t.Errorf("nextPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backspaceCompareRecursive(t *testing.T) {
	type args struct {
		S string
		T string
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
			if got := backspaceCompareRecursive(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompareRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
