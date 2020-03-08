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
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	result = append(inorderTraversal(root.Left), root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	// treeMap := make(map[*TreeNode]bool, 0)
	// stack:=stackNew()
	// if !treeMap[root] {
	// }
	return result
}

// var stack []*TreeNode

// stack = append(stack, "world!") // Push
// stack = append(stack, "Hello ")

// for len(stack) > 0 {
// 	n := len(stack) - 1 // Top element
// 	fmt.Print(stack[n])

// 	stack = stack[:n] // Pop
// }

// @lc code=end
