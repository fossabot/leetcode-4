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

import "math"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidBSTInRange(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTInRange(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	} else if root.Val <= low || root.Val >= high {
		return false
	}
	return checkSubTree(root, low, high)
}

func checkSubTree(root *TreeNode, low, high int) bool {
	leftRes := isValidBSTInRange(root.Left, low, root.Val)
	rightRes := isValidBSTInRange(root.Right, root.Val, high)

	if leftRes == false || rightRes == false {
		return false
	}
	return true
}

// @lc code=end
