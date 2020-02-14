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
	strLength := len(s)
	halfStrLength := strLength / 2
	lowerHalf := lengthOfLongestSubstring(s[:halfStrLength])
	higherHalf := lengthOfLongestSubstring(s[(halfStrLength + 1):])
	highest := max(lowerHalf, higherHalf)

	subStr := s[halfStrLength-lowerHalf+1 : halfStrLength+higherHalf]
	var elementSet map[rune]bool

	for _, currentChar := range subStr {
		elementSet[currentChar] = true
	}
	return highest
}

func max(val1 int, val2 int) int {
	if val1 >= val2 {
		return val1
	} else {
		return val2
	}
}

// @lc code=end
