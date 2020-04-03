/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 *
 * https://leetcode.com/problems/string-to-integer-atoi/description/
 *
 * algorithms
 * Medium (15.02%)
 * Likes:    1318
 * Dislikes: 8013
 * Total Accepted:    489.6K
 * Total Submissions: 3.3M
 * Testcase Example:  '"42"'
 *
 * Implement atoi which converts a string to an integer.
 *
 * The function first discards as many whitespace characters as necessary until
 * the first non-whitespace character is found. Then, starting from this
 * character, takes an optional initial plus or minus sign followed by as many
 * numerical digits as possible, and interprets them as a numerical value.
 *
 * The string can contain additional characters after those that form the
 * integral number, which are ignored and have no effect on the behavior of
 * this function.
 *
 * If the first sequence of non-whitespace characters in str is not a valid
 * integral number, or if no such sequence exists because either str is empty
 * or it contains only whitespace characters, no conversion is performed.
 *
 * If no valid conversion could be performed, a zero value is returned.
 *
 * Note:
 *
 *
 * Only the space character ' ' is considered as whitespace character.
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. If the numerical
 * value is out of the range of representable values, INT_MAX (2^31 − 1) or
 * INT_MIN (−2^31) is returned.
 *
 *
 * Example 1:
 *
 *
 * Input: "42"
 * Output: 42
 *
 *
 * Example 2:
 *
 *
 * Input: "   -42"
 * Output: -42
 * Explanation: The first non-whitespace character is '-', which is the minus
 * sign.
 * Then take as many numerical digits as possible, which gets 42.
 *
 *
 * Example 3:
 *
 *
 * Input: "4193 with words"
 * Output: 4193
 * Explanation: Conversion stops at digit '3' as the next character is not a
 * numerical digit.
 *
 *
 * Example 4:
 *
 *
 * Input: "words and 987"
 * Output: 0
 * Explanation: The first non-whitespace character is 'w', which is not a
 * numerical
 * digit or a +/- sign. Therefore no valid conversion could be performed.
 *
 * Example 5:
 *
 *
 * Input: "-91283472332"
 * Output: -2147483648
 * Explanation: The number "-91283472332" is out of the range of a 32-bit
 * signed integer.
 * Thefore INT_MIN (−2^31) is returned.
 *
 */

package leetcode

import "testing"

func Test_parseAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseAtoi(tt.args.str); got != tt.want {
				t.Errorf("parseAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{"42"},
			want: 42,
		},
		{
			name: "Test 2",
			args: args{"   -42"},
			want: -42,
		},
		{
			name: "Test 3",
			args: args{"4193 with words"},
			want: 4193,
		},
		{
			name: "Test 4",
			args: args{"words and 987"},
			want: 0,
		},
		{
			name: "Test 5",
			args: args{"-91283472332"},
			want: -2147483648,
		},
		{
			name: "Test 6",
			args: args{"+1"},
			want: 1,
		},
		{
			name: "Test 7",
			args: args{"+-1"},
			want: 0,
		},
		{
			name: "Test 8",
			args: args{"+"},
			want: 0,
		},
		{
			name: "Test 9",
			args: args{"9223372036854775808"},
			want: 2147483647,
		},
		{
			name: "Test 10",
			args: args{"18446744073709551617"},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
