/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (25.66%)
 * Likes:    5664
 * Dislikes: 688
 * Total Accepted:    783.4K
 * Total Submissions: 3M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 *
 * Note:
 *
 * The solution set must not contain duplicate triplets.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 *
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{[]int{}},
			want: [][]int{},
		},
		{
			name: "Test 2",
			args: args{[]int{0}},
			want: [][]int{},
		},
		{
			name: "Test 3",
			args: args{[]int{-1, 0, 1, 0}},
			want: [][]int{{-1, 0, 1}},
		},
		{
			name: "Test 4",
			args: args{[]int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "Test 5",
			args: args{[]int{0, 0, 0}},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Test 6",
			args: args{[]int{0, 0, 0, 0}},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Test 7",
			args: args{[]int{1, 2, -2, -1}},
			want: [][]int{},
		},
		{
			name: "Test 8",
			args: args{[]int{-2, 0, 1, 1, 2}},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			name: "Test 9",
			args: args{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}},
			want: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
