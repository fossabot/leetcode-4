/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (60.62%)
 * Likes:    2510
 * Dislikes: 105
 * Total Accepted:    633.5K
 * Total Submissions: 1M
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the inorder traversal of its nodes' values.
 *
 * Example:
 *
 *
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * Output: [1,3,2]
 *
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 *
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recursiveInOrderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := recursiveInOrderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recursiveInOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_iterativeInOrderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := iterativeInOrderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("iterativeInOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stackPop(t *testing.T) {
	type args struct {
		stack []*TreeNode
	}
	tests := []struct {
		name  string
		args  args
		want  *TreeNode
		want1 []*TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := stackPop(tt.args.stack)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stackPop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("stackPop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
