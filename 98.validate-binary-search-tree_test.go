/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 *
 * https://leetcode.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (26.77%)
 * Likes:    3288
 * Dislikes: 467
 * Total Accepted:    612.8K
 * Total Submissions: 2.2M
 * Testcase Example:  '[2,1,3]'
 *
 * Given a binary tree, determine if it is a valid binary search tree (BST).
 *
 * Assume a BST is defined as follows:
 *
 *
 * The left subtree of a node contains only nodes with keys less than the
 * node's key.
 * The right subtree of a node contains only nodes with keys greater than the
 * node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 *
 * Input: [2,1,3]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 *
 * Input: [5,1,4,null,null,3,6]
 * Output: false
 * Explanation: The root node's value is 5 but its right child's value is 4.
 *
 *
 */

package leetcode

import (
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{&TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{4,
				&TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}}},
			want: false,
		},
		{
			name: "Test 2",
			args: args{&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}},
			want: true,
		},
		{
			name: "Test 3",
			args: args{&TreeNode{1, &TreeNode{1, nil, nil}, nil}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBSTInRange(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
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
			if got := isValidBSTInRange(tt.args.root, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("isValidBSTInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkSubTree(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
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
			if got := checkSubTree(tt.args.root, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("checkSubTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
