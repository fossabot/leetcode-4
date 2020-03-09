/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (32.37%)
 * Likes:    6592
 * Dislikes: 1714
 * Total Accepted:    1.1M
 * Total Submissions: 3.5M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Example:
 *
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 *
 *
 */

package leetcode

import "strconv"

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode is a data structure for Linked-List Data type
type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(vals []int) *ListNode {
	listNode := &ListNode{}
	original := listNode
	for _, v := range vals {
		temp := &ListNode{}
		temp.Val = v
		listNode.Next = temp
		listNode = listNode.Next
	}
	return original.Next
}

func equal(l1, l2 *ListNode) bool {
	for {
		if l1 == nil && l2 == nil {
			return true
		} else if l1 == nil || l2 == nil || l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
}

func print(l *ListNode) (s string) {
	s = ""
	for l != nil {
		s += (strconv.Itoa(l.Val) + " ")
		l = l.Next
	}
	return s
}

// @lc code=start

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNode := &ListNode{}
	curNode := resultNode
	leading := 0
	value := 0
	for {
		if l2 == nil && l1 != nil {
			value = l1.Val + leading
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			value = l2.Val + leading
			l2 = l2.Next
		} else if l1 == nil && l2 == nil {
			break
		} else {
			value = l1.Val + l2.Val + leading
			l1 = l1.Next
			l2 = l2.Next
		}

		curNode.Next = &ListNode{value % 10, nil}
		leading = value / 10
		curNode = curNode.Next
	}

	if leading != 0 {
		curNode.Next = &ListNode{1, nil}
	}

	return resultNode.Next
}

// @lc code=end
