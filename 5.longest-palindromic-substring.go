/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (28.43%)
 * Likes:    5008
 * Dislikes: 441
 * Total Accepted:    747.1K
 * Total Submissions: 2.6M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: "cbbd"
 * Output: "bb"
 *
 *
 */

package leetcode

// @lc code=start

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getInnerPalndromeLength(s string, left, right int) int {
	for ; left >= 0 && right < len(s); left, right = left-1, right+1 {
		if s[left] != s[right] {
			break
		}
	}
	return right - left - 1
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	maxLength := 1
	start := 0
	end := 1
	for i := 0; i < len(s)-1; i++ {
		even := getInnerPalndromeLength(s, i, i+1)
		odd := getInnerPalndromeLength(s, i, i)
		curLen := max(even, odd)
		if curLen > maxLength {
			maxLength = curLen
			if curLen == even {
				start = i - curLen/2 + 1
			} else {
				start = i - curLen/2
			}
			end = i + curLen/2 + 1
		}
	}
	return s[start:end]
}

// @lc code=end
