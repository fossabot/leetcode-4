/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (34.46%)
 * Likes:    1819
 * Dislikes: 1573
 * Total Accepted:    599.9K
 * Total Submissions: 1.7M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */

package leetcode

import "math"

// @lc code=start
func longestCommonPrefix(strs []string) string {
	commonPrefix := ""
	if len(strs) == 0 {
		return commonPrefix
	}
	minLengthStr := math.MaxInt32
	for _, str := range strs {
		if len(str) < minLengthStr {
			minLengthStr = len(str)
		}
	}
	for charIndex := 0; charIndex < minLengthStr; charIndex++ {
		currentChar := strs[0][charIndex]
		for strIndex := 0; strIndex < len(strs); strIndex++ {
			if currentChar != strs[strIndex][charIndex] {
				return commonPrefix
			}
		}
		commonPrefix += string(currentChar)
	}
	return commonPrefix
}

// @lc code=end
