/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 *
 * https://leetcode.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (47.94%)
 * Likes:    2421
 * Dislikes: 151
 * Total Accepted:    241K
 * Total Submissions: 498.5K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 *
 * Given a binary tree, you need to compute the length of the diameter of the
 * tree. The diameter of a binary tree is the length of the longest path
 * between any two nodes in a tree. This path may or may not pass through the
 * root.
 *
 *
 *
 * Example:
 * Given a binary tree
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \
 * ⁠     4   5
 *
 *
 *
 * Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
 *
 *
 * Note:
 * The length of path between two nodes is represented by the number of edges
 * between them.
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
func diameterOfBinaryTree(root *TreeNode) int {
	_, bothside := depthOfTree(root)
	return bothside
}

func depthOfTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	left, bothsideLeft := depthOfTree(root.Left)
	right, bothsideRight := depthOfTree(root.Right)
	singleSide := max(left, right) + 1
	bothSide := max(bothsideLeft, bothsideRight)
	bothSide = max(bothSide, left+right)

	return singleSide, bothSide
}

// @lc code=end
