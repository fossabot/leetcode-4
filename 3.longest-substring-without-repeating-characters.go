/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (29.26%)
 * Likes:    7220
 * Dislikes: 425
 * Total Accepted:    1.2M
 * Total Submissions: 4.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 *
 *
 *
 *
 *
 */

package leetcode

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]int)
	startingIndex, maxLength := 0, 0
	for index, runeChar := range s {
		char := byte(runeChar)
		value, ok := charSet[char]
		if ok && value >= startingIndex {
			uniqueLength := (index - startingIndex)
			if uniqueLength > maxLength {
				maxLength = uniqueLength
			}
			startingIndex = charSet[char] + 1
		}
		charSet[char] = index
	}
	if (len(s) - startingIndex) > maxLength {
		maxLength = (len(s) - startingIndex)
	}
	return maxLength
}

// @lc code=end
