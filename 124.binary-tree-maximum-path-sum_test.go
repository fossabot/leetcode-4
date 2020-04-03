/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 *
 * https://leetcode.com/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (31.79%)
 * Likes:    2742
 * Dislikes: 231
 * Total Accepted:    286.5K
 * Total Submissions: 881.1K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty binary tree, find the maximum path sum.
 *
 * For this problem, a path is defined as any sequence of nodes from some
 * starting node to any node in the tree along the parent-child connections.
 * The path must contain at least one node and does not need to go through the
 * root.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   3
 *
 * Output: 6
 *
 *
 * Example 2:
 *
 *
 * Input: [-10,9,20,null,null,15,7]
 *
 * -10
 * / \
 * 9  20
 * /  \
 * 15   7
 *
 * Output: 42
 *
 *
 */

package leetcode

import "testing"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rootSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name     string
		args     args
		wantISum int
		wantGSum int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotISum, gotGSum := rootSum(tt.args.root)
			if gotISum != tt.wantISum {
				t.Errorf("rootSum() gotISum = %v, want %v", gotISum, tt.wantISum)
			}
			if gotGSum != tt.wantGSum {
				t.Errorf("rootSum() gotGSum = %v, want %v", gotGSum, tt.wantGSum)
			}
		})
	}
}

func Test_curMax(t *testing.T) {
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
			if got := curMax(tt.args.nums); got != tt.want {
				t.Errorf("curMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
