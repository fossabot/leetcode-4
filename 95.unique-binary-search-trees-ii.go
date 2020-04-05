/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 *
 * https://leetcode.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (37.78%)
 * Likes:    1825
 * Dislikes: 146
 * Total Accepted:    175.6K
 * Total Submissions: 451.4K
 * Testcase Example:  '3'
 *
 * Given an integer n, generate all structurally unique BST's (binary search
 * trees) that store values 1 ... n.
 *
 * Example:
 *
 *
 * Input: 3
 * Output:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * Explanation:
 * The above output corresponds to the 5 unique BST's shown below:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTreesInRange(1, n)
}

func generateTreesInRange(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	res := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftRes := generateTreesInRange(start, i-1)
		rightRes := generateTreesInRange(i+1, end)
		for _, leftNode := range leftRes {
			for _, rightNode := range rightRes {
				// New Treenode with left pointer to leftNode and right pointer
				// to RightNode and i as value
				res = append(res, &TreeNode{i, leftNode, rightNode})
			}
		}
	}
	return res
}

// @lc code=end
