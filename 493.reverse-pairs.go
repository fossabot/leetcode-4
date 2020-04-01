/*
 * @lc app=leetcode id=493 lang=golang
 *
 * [493] Reverse Pairs
 *
 * https://leetcode.com/problems/reverse-pairs/description/
 *
 * algorithms
 * Hard (24.11%)
 * Likes:    737
 * Dislikes: 108
 * Total Accepted:    33.6K
 * Total Submissions: 137.6K
 * Testcase Example:  '[1,3,2,3,1]'
 *
 * Given an array nums, we call (i, j) an important reverse pair if i < j and
 * nums[i] > 2*nums[j].
 *
 * You need to return the number of important reverse pairs in the given
 * array.
 *
 * Example1:
 *
 * Input: [1,3,2,3,1]
 * Output: 2
 *
 *
 * Example2:
 *
 * Input: [2,4,3,5,1]
 * Output: 3
 *
 *
 * Note:
 *
 * The length of the given array will not exceed 50,000.
 * All the numbers in the input array are in the range of 32-bit integer.
 *
 *
 */

package leetcode

import "sort"

// @lc code=start
func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, s, e int) int {
	if s >= e {
		return 0
	}
	mid := s + (e-s)/2
	cnt := mergeSort(nums, s, mid) + mergeSort(nums, mid+1, e)
	for i, j := s, mid+1; i <= mid; i++ {
		for j <= e && nums[i] > nums[j]*2 {
			j++
		}
		cnt += j - (mid + 1)
	}
	sort.Ints(nums[s : e+1])
	return cnt
}

// func countReversePairs(nums []int) (int, map[int]int, []int) {
// 	numsLen := len(nums)
// 	if numsLen <= 1 {
// 		keys, keyMap := bucketMap(nums)
// 		return 0, keys, keyMap
// 	}

// 	mid := numsLen / 2
// 	leftCount, leftKeyMap, leftKeys := countReversePairs(nums[:mid])
// 	rightCount, rightKeyMap, rightKeys := countReversePairs(nums[mid:])

// 	count := mergeCount(leftKeyMap, rightKeyMap, leftKeys, rightKeys)
// 	count += (leftCount + rightCount)
// 	mergedKeyMap, mergedKeys := mergeKeyMap(leftKeyMap, rightKeyMap,
// 		leftKeys, rightKeys)
// 	// fmt.Println(leftCount, leftKeyMap, leftKeys)
// 	// fmt.Println(rightCount, rightKeyMap, rightKeys)
// 	// fmt.Println(count, mergedKeyMap, mergedKeys)
// 	// fmt.Println("step")

// 	return count, mergedKeyMap, mergedKeys
// }

// func bucketMap(nums []int) (map[int]int, []int) {
// 	keyMap := make(map[int]int)
// 	keys := make([]int, 0)
// 	for _, num := range nums {
// 		keyMap[num]++
// 		keys = []int{num}
// 	}
// 	return keyMap, keys
// }

// func mergeCount(leftKeyMap, rightKeyMap map[int]int, leftKeys,
// 	rightKeys []int) int {
// 	count := 0
// 	for _, key := range rightKeys {
// 		i := sort.SearchInts(leftKeys, 2*key+1)
// 		for _, leftKey := range leftKeys[i:] {
// 			count += leftKeyMap[leftKey]
// 			fmt.Println(key, i, leftKey, count, leftKeyMap[leftKey])
// 		}
// 	}
// 	return count
// }

// func mergeKeyMap(leftKeyMap, rightKeyMap map[int]int, leftKeys,
// 	rightKeys []int) (map[int]int, []int) {
// 	leftKeyLen, rightKeyLen := len(leftKeys), len(rightKeys)
// 	mergedKeyMap := make(map[int]int)
// 	mergedKeys := make([]int, 0)

// 	i, j := 0, 0

// 	for i < leftKeyLen && j < rightKeyLen {
// 		leftKey, rightKey := leftKeys[i], rightKeys[j]
// 		if leftKey < rightKey {
// 			mergedKeyMap[leftKey] = leftKeyMap[leftKey]
// 			mergedKeys = append(mergedKeys, leftKey)
// 			i++
// 		} else if leftKey == rightKey {
// 			// rightKey == leftKey. So they can be used interchangibly
// 			mergedKeyMap[leftKey] = leftKeyMap[leftKey] +
// 				rightKeyMap[leftKey]
// 			mergedKeys = append(mergedKeys, leftKey)
// 			i++
// 			j++
// 		} else {
// 			mergedKeyMap[rightKey] = rightKeyMap[rightKey]
// 			mergedKeys = append(mergedKeys, rightKey)
// 			j++
// 		}
// 	}

// 	if i < leftKeyLen {
// 		for _, key := range leftKeys[i:] {
// 			mergedKeyMap[key] = leftKeyMap[key]
// 			mergedKeys = append(mergedKeys, key)
// 		}
// 	}

// 	if j < rightKeyLen {
// 		for _, key := range rightKeys[j:] {
// 			mergedKeyMap[key] = rightKeyMap[key]
// 			mergedKeys = append(mergedKeys, key)
// 		}
// 	}

// 	return mergedKeyMap, mergedKeys
// }

// @lc code=end
