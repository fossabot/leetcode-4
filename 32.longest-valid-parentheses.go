/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 *
 * https://leetcode.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (27.27%)
 * Likes:    2839
 * Dislikes: 122
 * Total Accepted:    252.1K
 * Total Submissions: 921K
 * Testcase Example:  '"(()"'
 *
 * Given a string containing just the characters '(' and ')', find the length
 * of the longest valid (well-formed) parentheses substring.
 *
 * Example 1:
 *
 *
 * Input: "(()"
 * Output: 2
 * Explanation: The longest valid parentheses substring is "()"
 *
 *
 * Example 2:
 *
 *
 * Input: ")()())"
 * Output: 4
 * Explanation: The longest valid parentheses substring is "()()"
 *
 *
 */

package leetcode

// @lc code=start
func longestValidParentheses(s string) int {
	maxLen := 0
	for i, left, right := 0, 0, 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++ // Assuming only ( or )
		}
		if left == right {
			if maxLen < 2*left {
				maxLen = 2 * left
			}
		} else if left < right {
			left = 0
			right = 0
		}
	}
	for i, left, right := len(s)-1, 0, 0; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++ // Assuming only ( or )
		}
		if left == right {
			if maxLen < 2*left {
				maxLen = 2 * left
			}
		} else if left > right {
			left = 0
			right = 0
		}

	}
	return maxLen
}

// @lc code=end
