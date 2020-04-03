/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 *
 * https://leetcode.com/problems/lexicographical-numbers/description/
 *
 * algorithms
 * Medium (49.34%)
 * Likes:    496
 * Dislikes: 66
 * Total Accepted:    49.1K
 * Total Submissions: 99K
 * Testcase Example:  '13'
 *
 * Given an integer n, return 1 - n in lexicographical order.
 *
 * For example, given 13, return: [1,10,11,12,13,2,3,4,5,6,7,8,9].
 *
 * Please optimize your algorithm to use less time and space. The input size
 * may be as large as 5,000,000.
 *
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_lexicalOrder(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lexicalOrder(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intStack_Push(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		s    intStack
		args args
		want intStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Push(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intStack.Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intStack_Pop(t *testing.T) {
	tests := []struct {
		name  string
		s     intStack
		want  intStack
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Pop()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intStack.Pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("intStack.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_intStack_Len(t *testing.T) {
	tests := []struct {
		name string
		s    intStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("intStack.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
