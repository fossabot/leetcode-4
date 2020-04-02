package leetcode

import "strconv"

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

func equalList(l1, l2 *ListNode) bool {
	for {
		if l1 == nil && l2 == nil {
			return true
		} else if l1 == nil || l2 == nil || l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
}

func printList(l *ListNode) (s string) {
	s = ""
	for l != nil {
		s += (strconv.Itoa(l.Val) + " ")
		l = l.Next
	}
	return s
}
