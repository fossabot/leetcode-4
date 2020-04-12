/*
 * @lc app=leetcode id=1046 lang=golang
 *
 * [1046] Last Stone Weight
 *
 * https://leetcode.com/problems/last-stone-weight/description/
 *
 * algorithms
 * Easy (62.53%)
 * Likes:    506
 * Dislikes: 24
 * Total Accepted:    63.7K
 * Total Submissions: 100.6K
 * Testcase Example:  '[2,7,4,1,8,1]'
 *
 * We have a collection of stones, each stone has a positive integer weight.
 *
 * Each turn, we choose the two heaviest stones and smash them together.
 * Suppose the stones have weights x and y with x <= y.  The result of this
 * smash is:
 *
 *
 * If x == y, both stones are totally destroyed;
 * If x != y, the stone of weight x is totally destroyed, and the stone of
 * weight y has new weight y-x.
 *
 *
 * At the end, there is at most 1 stone left.  Return the weight of this stone
 * (or 0 if there are no stones left.)
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [2,7,4,1,8,1]
 * Output: 1
 * Explanation:
 * We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
 * we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
 * we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
 * we combine 1 and 1 to get 0 so the array converts to [1] then that's the
 * value of last stone.
 *
 *
 *
 * Note:
 *
 *
 * 1 <= stones.length <= 30
 * 1 <= stones[i] <= 1000
 *
 *
 */

package leetcode

import (
	"container/heap"
)

// @lc code=start
func lastStoneWeight(stones []int) int {
	h := toIntHeap(stones)
	heap.Init(h)
	for h.Len() > 1 {
		big, small := heap.Pop(h), heap.Pop(h)
		diff := big.(int) - small.(int)
		if diff != 0 {
			heap.Push(h, diff)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// Maxheap implementation
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push uses pointer receivers because they modify the slice's length,
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop uses pointer receivers because they modify the slice's length,
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func toIntHeap(nums []int) *IntHeap {
	c := make(IntHeap, len(nums))
	for i, v := range nums {
		c[i] = v
	}
	return &c
}

// @lc code=end
