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

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	return iterativeInOrderTraversal(root)
}

func recursiveInOrderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(inorderTraversal(root.Left), root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

func iterativeInOrderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	visited := make(map[*TreeNode]bool)
	for len(stack) != 0 {
		var curNode *TreeNode
		curNode, stack = stackPop(stack)
		if curNode == nil {
			continue
		}
		yes, _ := visited[curNode]
		if yes {
			result = append(result, curNode.Val)
		} else {
			visited[curNode] = true
			stack = append(stack, curNode.Right, curNode, curNode.Left)
		}
	}
	return result
}

func stackPop(stack []*TreeNode) (*TreeNode, []*TreeNode) {
	end := len(stack) - 1
	num := stack[end]
	return num, stack[:end]
}

// @lc code=end
