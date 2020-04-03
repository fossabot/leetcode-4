package leetcode

import "strconv"

// ListNode defines a singly-linked list.
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

func printList(l *ListNode) (s string) {
	s = ""
	for l != nil {
		s += (strconv.Itoa(l.Val) + " ")
		l = l.Next
	}
	return s
}

// An pqListNode is something we manage in a priority queue.
type pqListNode struct {
	node *ListNode // The value of the item; arbitrary.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A pqLinkedList implements heap.Interface and holds Items.
type pqLinkedList []*ListNode

func (pq pqLinkedList) Len() int { return len(pq) }

func (pq pqLinkedList) Less(i, j int) bool {
	// Get the lowest Node from queue
	return pq[i].Val < pq[j].Val
}

func (pq pqLinkedList) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pqLinkedList) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *pqLinkedList) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
