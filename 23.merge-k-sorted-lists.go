/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (38.05%)
 * Likes:    3813
 * Dislikes: 248
 * Total Accepted:    555.9K
 * Total Submissions: 1.5M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * Merge k sorted linked lists and return it as one sorted list. Analyze and
 * describe its complexity.
 *
 * Example:
 *
 *
 * Input:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * Output: 1->1->2->3->4->4->5->6
 *
 *
 */

package leetcode

import (
	"container/heap"
)

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	q := pqLinkedList{}
	for _, list := range lists {
		if list != nil {
			heap.Push(&q, list)
		}
	}

	solution := &ListNode{}
	head := solution

	for q.Len() != 0 {
		it := heap.Pop(&q).(*ListNode)
		solution.Next = it
		if it.Next != nil {
			heap.Push(&q, it.Next)
		}
		solution = solution.Next
	}
	return head.Next
}

// @lc code=end
