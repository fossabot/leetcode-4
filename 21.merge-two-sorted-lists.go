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

package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNode := ListNode{0, nil}
	for {
		if l1.Val < l2.Val {
			resultNode.Next = ListNode{l1.Val,nil}
			l1 = l1.Next
		} else {
			resultNode.Next = ListNode{l2.Val,nil}
			l2 = l2.Next
		}
		if l1==nil || l2==nil {
			break
		}
	}
}

// @lc code=end
