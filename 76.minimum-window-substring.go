/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 *
 * https://leetcode.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (33.26%)
 * Likes:    3584
 * Dislikes: 252
 * Total Accepted:    337.7K
 * Total Submissions: 1M
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * Given a string S and a string T, find the minimum window in S which will
 * contain all the characters in T in complexity O(n).
 *
 * Example:
 *
 *
 * Input: S = "ADOBECODEBANC", T = "ABC"
 * Output: "BANC"
 *
 *
 * Note:
 *
 *
 * If there is no such window in S that covers all characters in T, return the
 * empty string "".
 * If there is such window, you are guaranteed that there will always be only
 * one unique minimum window in S.
 *
 *
 */

package leetcode

// @lc code=start

// // An Item is something we manage in a priority queue.
// type Item struct {
// 	value    rune // The value of the item; arbitrary.
// 	priority int  // The priority of the item in the queue.
// 	// The index is needed by update and is maintained by the heap.Interface methods.
// 	index int // The index of the item in the heap.
// }

// // A PriorityQueue implements heap.Interface and holds Items.
// type PriorityQueue []*Item

// func (pq PriorityQueue) Len() int { return len(pq) }

// func (pq PriorityQueue) Less(i, j int) bool {
// 	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
// 	return pq[i].priority < pq[j].priority
// }

// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// 	pq[i].index = i
// 	pq[j].index = j
// }

// // Push implements pushing a new element in Priority queue
// func (pq *PriorityQueue) Push(x interface{}) {
// 	n := len(*pq)
// 	item := x.(*Item)
// 	item.index = n
// 	*pq = append(*pq, item)
// }

// // Pop extracts the lowest element from Priority Queue
// func (pq *PriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	old[n-1] = nil  // avoid memory leak
// 	item.index = -1 // for safety
// 	*pq = old[0 : n-1]
// 	return item
// }

// // update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, value string, priority int) {
// 	item.value = value
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }

type char struct {
	ch    rune
	index int
}

func minWindow(s string, t string) string {
	tMap := make(map[rune]int)
	lst := make([]char, 0)
	left, right := 0, 0
	// min := 0

	for _, ch := range t {
		tMap[ch] = 0
	}

	for i, ch := range s {
		if _, ok := tMap[ch]; ok {
			tMap[ch]++
			lst = append(lst, char{ch, i})
		}
		for tMap[lst[0].ch] != 1 {
			lst = lst[1:]
		}
	}
	if right == 0 {
		return ""
	}
	return s[left:right]
}

// @lc code=end
