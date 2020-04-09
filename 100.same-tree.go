/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 *
 * https://leetcode.com/problems/same-tree/description/
 *
 * algorithms
 * Easy (51.38%)
 * Likes:    1778
 * Dislikes: 53
 * Total Accepted:    502.4K
 * Total Submissions: 965.4K
 * Testcase Example:  '[1,2,3]\n[1,2,3]'
 *
 * Given two binary trees, write a function to check if they are the same or
 * not.
 *
 * Two binary trees are considered the same if they are structurally identical
 * and the nodes have the same value.
 *
 * Example 1:
 *
 *
 * Input:     1         1
 * ⁠         / \       / \
 * ⁠        2   3     2   3
 *
 * ⁠       [1,2,3],   [1,2,3]
 *
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input:     1         1
 * ⁠         /           \
 * ⁠        2             2
 *
 * ⁠       [1,2],     [1,null,2]
 *
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input:     1         1
 * ⁠         / \       / \
 * ⁠        2   1     1   2
 *
 * ⁠       [1,2,1],   [1,1,2]
 *
 * Output: false
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Iterative implementation
	// if p.Val == q.Val {
	// 	if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
	// 		return true
	// 	}
	// }
	// return false

	pQ := []*TreeNode{p}
	qQ := []*TreeNode{q}

	for {
		// Pop operation
		p, pQ = pop(pQ)
		q, qQ = pop(qQ)

		if p == nil && q == nil {
			if len(pQ) == 0 {
				return true
			}
			continue
		} else if p == nil || q == nil {
			return false
		} else if p.Val != q.Val {
			return false
		}
		pQ = append(pQ, p.Left, p.Right)
		qQ = append(qQ, q.Left, q.Right)
	}
}

func pop(queue []*TreeNode) (*TreeNode, []*TreeNode) {
	if len(queue) == 1 {
		return queue[0], []*TreeNode{}
	}
	return queue[0], queue[1:]
}

// @lc code=end
