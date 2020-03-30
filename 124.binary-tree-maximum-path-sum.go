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

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	_, curGSum := rootSum(root)
	return curGSum
}

func rootSum(root *TreeNode) (iSum, gSum int) {
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	} else if root.Left == nil {
		rightISum, rightGSum := rootSum(root.Right)
		iSum = curMax([]int{rightISum, 0}) + root.Val
		gSum = curMax([]int{iSum, rightGSum})
		return iSum, gSum
	} else if root.Right == nil {
		leftISum, leftGSum := rootSum(root.Left)
		iSum = curMax([]int{leftISum, 0}) + root.Val
		gSum = curMax([]int{iSum, leftGSum})
		return iSum, gSum
	}
	leftISum, leftGSum := rootSum(root.Left)
	rightISum, rightGSum := rootSum(root.Right)

	iSum = curMax([]int{leftISum, rightISum, 0}) + root.Val
	gSum = curMax([]int{leftGSum, rightGSum, leftISum + root.Val + rightISum, iSum})
	return iSum, gSum
}

func curMax(nums []int) int {
	var max int
	for i, val := range nums {
		// Returning value from slice rather than pre-defined min
		if i == 0 || val > max {
			max = val
		}
	}
	return max
}

// @lc code=end
