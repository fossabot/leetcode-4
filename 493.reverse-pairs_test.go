/*
 * @lc app=leetcode id=493 lang=golang
 *
 * [493] Reverse Pairs
 *
 * https://leetcode.com/problems/reverse-pairs/description/
 *
 * algorithms
 * Hard (24.11%)
 * Likes:    737
 * Dislikes: 108
 * Total Accepted:    33.6K
 * Total Submissions: 137.6K
 * Testcase Example:  '[1,3,2,3,1]'
 *
 * Given an array nums, we call (i, j) an important reverse pair if i < j and
 * nums[i] > 2*nums[j].
 *
 * You need to return the number of important reverse pairs in the given
 * array.
 *
 * Example1:
 *
 * Input: [1,3,2,3,1]
 * Output: 2
 *
 *
 * Example2:
 *
 * Input: [2,4,3,5,1]
 * Output: 3
 *
 *
 * Note:
 *
 * The length of the given array will not exceed 50,000.
 * All the numbers in the input array are in the range of 32-bit integer.
 *
 *
 */

package leetcode

import (
	"testing"
)

func Test_reversePairs(t *testing.T) {
	type args struct {
		nums []int
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
			if got := reversePairs(tt.args.nums); got != tt.want {
				t.Errorf("reversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexedReversePairs(t *testing.T) {
	type args struct {
		nums []int
		s    int
		e    int
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
			if got := indexedReversePairs(tt.args.nums, tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("indexedReversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
