/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 *
 * https://leetcode.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (29.77%)
 * Likes:    4638
 * Dislikes: 201
 * Total Accepted:    438.9K
 * Total Submissions: 1.5M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * Design and implement a data structure for Least Recently Used (LRU) cache.
 * It should support the following operations: get and put.
 *
 * get(key) - Get the value (will always be positive) of the key if the key
 * exists in the cache, otherwise return -1.
 * put(key, value) - Set or insert the value if the key is not already present.
 * When the cache reached its capacity, it should invalidate the least recently
 * used item before inserting a new item.
 *
 * The cache is initialized with a positive capacity.
 *
 * Follow up:
 * Could you do both operations in O(1) time complexity?
 *
 * Example:
 *
 *
 * LRUCache cache = new LRUCache( 2 );
 *
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // returns 1
 * cache.put(3, 3);    // evicts key 2
 * cache.get(2);       // returns -1 (not found)
 * cache.put(4, 4);    // evicts key 1
 * cache.get(1);       // returns -1 (not found)
 * cache.get(3);       // returns 3
 * cache.get(4);       // returns 4
 *
 *
 *
 *
 */

package leetcode

// @lc code=start
type LRUCache struct {
	capacity    int
	elements    map[int]int
	queue       []int
	queueExtra  map[int]int
	numElements int
}

func Constructor(capacity int) LRUCache {
	l := new(LRUCache)
	l.capacity = capacity
	l.queue = make([]int, 0)
	l.elements = make(map[int]int)
	l.queueExtra = make(map[int]int)
	return *l
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.elements[key]; ok {
		this.queue = append(this.queue, key)
		this.queueExtra[key]++
		return val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.elements[key]
	if ok {
		this.elements[key] = value
		this.queue = append(this.queue, key)
		this.queueExtra[key]++
		return
	}
	this.elements[key] = value
	this.queue = append(this.queue, key)
	if this.numElements == this.capacity {
		for i, key := range this.queue {
			if val, ok := this.queueExtra[key]; ok && val != 0 {
				this.queueExtra[key]--
			} else {
				delete(this.elements, key)
				delete(this.queueExtra, key)
				this.queue = this.queue[i+1:]
				break
			}
		}
	} else {
		this.numElements++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
