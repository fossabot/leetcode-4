/*
 * @lc app=leetcode id=1038 lang=golang
 *
 * [1038] Binary Search Tree to Greater Sum Tree
 *
 * https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/description/
 *
 * algorithms
 * Medium (78.52%)
 * Likes:    601
 * Dislikes: 86
 * Total Accepted:    38.7K
 * Total Submissions: 48.6K
 * Testcase Example:  '[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]'
 *
 * Given the root of a binary search tree with distinct values, modify it so
 * that every node has a new value equal to the sum of the values of the
 * original tree that are greater than or equal to node.val.
 *
 * As a reminder, a binary search tree is a tree that satisfies these
 * constraints:
 *
 *
 * The left subtree of a node contains only nodes with keys less than the
 * node's key.
 * The right subtree of a node contains only nodes with keys greater than the
 * node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
 * Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
 *
 *
 *
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is between 1 and 100.
 * Each node will have value between 0 and 100.
 * The given tree is a binary search tree.
 *
 *
 *
 *
 * Note: This question is the same as 538:
 * https://leetcode.com/problems/convert-bst-to-greater-tree/
 *
 *
 *
 */

package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	sumBST(root, 0)

	return root
}

func sumBST(root *TreeNode, sum int) int {
	if root.Right != nil {
		sum = sumBST(root.Right, sum)
	}

	root.Val += sum
	sum = root.Val

	if root.Left != nil {
		sum = sumBST(root.Left, sum)
	}

	return sum
}

// @lc code=end
