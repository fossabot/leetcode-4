/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (50.37%)
 * Likes:    3046
 * Dislikes: 444
 * Total Accepted:    778.6K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNode := &ListNode{}
	curNode := resultNode
	for {
		if l1 == nil {
			curNode.Next = l2
			break
		} else if l2 == nil {
			curNode.Next = l1
			break
		} else {
			if l1.Val <= l2.Val {
				curNode.Next = l1
				l1, curNode = l1.Next, curNode.Next
			} else {
				curNode.Next = l2
				l2, curNode = l2.Next, curNode.Next
			}
		}

	}
	return resultNode.Next
}

// @lc code=end
