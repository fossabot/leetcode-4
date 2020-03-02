/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (34.65%)
 * Likes:    2651
 * Dislikes: 192
 * Total Accepted:    533.3K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 *
 * Example:
 *
 *
 * Given linked list: 1->2->3->4->5, and n = 2.
 *
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 *
 *
 * Note:
 *
 * Given n will always be valid.
 *
 * Follow up:
 *
 * Could you do this in one pass?
 *
 */

package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curNode := head
	for i := 0; i < n; i++ {
		curNode = curNode.Next
	}
	if curNode == nil {
		return head.Next
	}
	rNode := head
	curNode = curNode.Next
	for curNode != nil {
		rNode, curNode = rNode.Next, curNode.Next
	}
	rNode.Next = rNode.Next.Next
	return head
}

// @lc code=end
