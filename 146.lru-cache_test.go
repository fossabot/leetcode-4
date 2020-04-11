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

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want LRUCache
	}{
		{
			name: "Test1",
			args: args{2},
			want: create(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := create(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Get(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		object *LRUCache
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.object.Get(tt.args.key); got != tt.want {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Put(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name   string
		object *LRUCache
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.object.Put(tt.args.key, tt.args.value)
		})
	}
}

// cache := Constructor(2)

// cache.Put(1, 1)
// cache.Put(2, 2)
// val := cache.Get(1) // returns 1
// if val != 1 {
// 	t.Errorf("Wrong Output. Expected 1, got %v", val)
// }
// cache.Put(3, 3)    // evicts key 2
// val = cache.Get(2) // returns -1 (not found)
// if val != -1 {
// 	t.Errorf("Wrong Output. Expected -1, got %v", val)
// }
// cache.Put(4, 4)    // evicts key 1
// val = cache.Get(1) // returns -1 (not found)
// if val != -1 {
// 	t.Errorf("Wrong Output. Expected -1, got %v", val)
// }
// val = cache.Get(3) // returns 3
// if val != 3 {
// 	t.Errorf("Wrong Output. Expected 3, got %v", val)
// }
// val = cache.Get(4) // returns 4
// if val != 4 {
// 	t.Errorf("Wrong Output. Expected 4, got %v", val)
// }
// cache.Put(2, 1)
// cache.Put(1, 1)
// cache.Put(2, 3)
// cache.Put(4, 1)     // evicts key 1
// val := cache.Get(1) // returns 1
// if val != -1 {
// 	t.Errorf("Wrong Output. Expected -1, got %v", val)
// }
// val = cache.Get(2) // returns -1 (not found)
// if val != 3 {
// 	t.Errorf("Wrong Output. Expected 3, got %v", val)
// }
// cache.Put(2, 1)
// cache.Put(2, 2)
// val := cache.Get(2)
// if val != 2 {
// 	t.Errorf("Wrong Output. Expected 2, got %v", val)
// }
// cache.Put(1, 1)
// cache.Put(4, 1)    // evicts key 1
// val = cache.Get(2) // returns 1
// if val != -1 {
// 	t.Errorf("Wrong Output. Expected -1, got %v", val)
// }
