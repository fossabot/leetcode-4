/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 *
 * https://leetcode.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (52.17%)
 * Likes:    1181
 * Dislikes: 141
 * Total Accepted:    139.1K
 * Total Submissions: 261.4K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The most significant digit comes first and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 * Follow up:
 * What if you cannot modify the input lists? In other words, reversing the
 * lists is not allowed.
 *
 *
 *
 * Example:
 *
 * Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 8 -> 0 -> 7
 *
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Stack := sliceFromList(l1)
	l2Stack := sliceFromList(l2)

	l1Len, l2Len := len(l1Stack), len(l2Stack)
	if l1Len < l2Len {
		l1Stack, l2Stack = l2Stack, l1Stack
		l1Len, l2Len = l2Len, l1Len
	}

	carry := 0
	var pre *ListNode
	i := l1Len - 1
	for j := l2Len - 1; j >= 0; i, j = i-1, j-1 {
		carry, pre = sumToList(carry, l1Stack[i], l2Stack[j], pre)
	}

	for ; i >= 0; i-- {
		carry, pre = sumToList(carry, l1Stack[i], 0, pre)
	}

	if carry != 0 {
		cur := &ListNode{}
		cur.Next = pre
		cur.Val = carry
		pre = cur
	}

	return pre
}

func sumToList(carry int, num1 int, num2 int, pre *ListNode) (int, *ListNode) {
	cur := &ListNode{}
	cur.Next = pre
	cur.Val = (num1 + num2 + carry)
	carry = cur.Val / 10
	cur.Val %= 10
	pre = cur
	return carry, pre
}

func sliceFromList(list *ListNode) []int {
	slice := make([]int, 0)
	for iter := list; iter != nil; iter = iter.Next {
		slice = append(slice, iter.Val)
	}
	return slice
}

// @lc code=end
