/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (51.28%)
 * Likes:    2705
 * Dislikes: 155
 * Total Accepted:    512.5K
 * Total Submissions: 962.4K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings, group anagrams together.
 *
 * Example:
 *
 *
 * Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * Output:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * Note:
 *
 *
 * All inputs will be in lowercase.
 * The order of your output does not matter.
 *
 *
 */

package leetcode

import "sort"

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// using string as key
	strMap := make(map[string][]string)
	// keys := make([]string, 0)

	for _, str := range strs {
		// Using sorted string as key
		key := sortString(str)

		// Using number as key
		// Bug: Can't handle character repeat
		// If array is used, memory usage increased
		// key := getKey(str)
		strSlice, found := strMap[key]
		if found {
			strMap[key] = append(strSlice, str)
		} else {
			// keys = append(keys, key)
			strMap[key] = []string{str}
		}
	}

	result := make([][]string, 0)

	for _, strArray := range strMap {
		result = append(result, strArray)
	}

	return result
}

type byteSlice []byte

func (a byteSlice) Len() int           { return len(a) }
func (a byteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byteSlice) Less(i, j int) bool { return a[i] < a[j] }

func sortString(s string) string {
	b := byteSlice(s)
	sort.Sort(b)
	return string(b)
}

func getKey(s string) int {
	lenS := len(s)
	key := 0

	for i := 0; i < lenS; i++ {
		key = key | (1 << (s[i] - 'a'))
	}

	return key
}

// @lc code=end
