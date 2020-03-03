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

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Item struct {
	node  *ListNode
	index int // The index of the item in the heap.
}

// type Heap []*Item

// func (h Heap) Len() int {
// 	return len(h)
// }

// func (h Heap) Less(i, j int) bool {
// 	return h[i].node.Val < h[j].node.Val
// }

// func (h Heap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// 	h[i].index = i
// 	h[j].index = j
// }

// func (h *Heap) Up(j int) {
// 	for {
// 		i := (j - 1) / 2 // parent
// 		if i == j || !h.Less(j, i) {
// 			break
// 		}
// 		h.Swap(i, j)
// 		j = i
// 	}
// }

// func (h *Heap) Down(i int, n int) {
// 	for {
// 		j1 := 2*i + 1
// 		if j1 >= n || j1 < 0 {
// 			break
// 		}
// 		j := j1 // left child
// 		j2 := j1 + 1
// 		if j2 < n && !h.Less(j1, j2) {
// 			j = j2 // = 2*i + 2 // right child
// 		}
// 		if !h.Less(j, i) {
// 			break
// 		}
// 		h.Swap(i, j)
// 		i = j
// 	}
// }

// func (h *Heap) Push(item *Item) {
// 	n := h.Len()
// 	item.index = n
// 	*h = append(*h, item)

// 	h.Up(h.Len() - 1)
// }

// func (h *Heap) Pop() *Item {
// 	n := h.Len() - 1

// 	h.Swap(0, n)
// 	h.Down(0, n)

// 	item := (*h)[n]
// 	item.index = -1 // for safety
// 	*h = (*h)[0:n]

// 	return item
// }

// func (h *Heap) Fix(i int) {
// 	h.Down(i, h.Len())
// 	h.Up(i)
// }

// func (h *Heap) Update(item *Item) {
// 	h.Fix(item.index)
// }

// func mergeKLists(lists []*ListNode) *ListNode {
// 	solution := &ListNode{}
// 	curNode := solution
// 	heapNodes := lists
// 	heap.Init(heapNodes)
// 	for {
// 		Node := heap.Pop(heapNodes)
// 		curNode.Next = Node
// 		curNode = curNode.Next
// 		heap.Push(heapNodes, Node.Next)
// 		if heapNodes.Len() == 0 {
// 			curNode.Next = nil
// 			break
// 		}
// 	}
// 	return solution.Next
// }

// func sorted(lists []*ListNode) []*ListNode {
// 	result := make([]*ListNode, 0)
// 	for _, Node := range lists {
// 		if Node != nil {
// 			result = append(result, Node)
// 		}
// 	}
// 	sort.SliceStable(result, func(i, j int) bool {
// 		return lists[i].Val < lists[j].Val
// 	})
// 	return result
// }

// @lc code=end
