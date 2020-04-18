/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 *
 * https://leetcode.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (40.97%)
 * Likes:    754
 * Dislikes: 971
 * Total Accepted:    184.5K
 * Total Submissions: 432.1K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * Given an integer array nums, find the sum of the elements between indices i
 * and j (i ≤ j), inclusive.
 *
 * Example:
 *
 * Given nums = [-2, 0, 3, -5, 2, -1]
 *
 * sumRange(0, 2) -> 1
 * sumRange(2, 5) -> -1
 * sumRange(0, 5) -> -3
 *
 *
 *
 * Note:
 *
 * You may assume that the array does not change.
 * There are many calls to sumRange function.
 *
 *
 */

package leetcode

import (
	"reflect"
	"testing"
)

func TestNumArrayConstructor(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want NumArray
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumArrayConstructor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumArrayConstructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumArray_SumRange(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		this *NumArray
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.SumRange(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("NumArray.SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
