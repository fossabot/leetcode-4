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

import "fmt"

// @lc code=start
func longestValidParentheses(s string) int {
	maxLen := 0
	for i, open, l, left := 0, 0, 0, 0; i < len(s); i++ {
		l++
		// fmt.Printf("%c ", s[i])
		// fmt.Println(i, maxLen, l, open)
		if s[i] == '(' {
			open++
		} else {
			open--
			if open == 0 {
				if (l + left) > maxLen {
					maxLen = (l + left)
				}
			} else if open > 0 {
				if (l - open) > maxLen {
					maxLen = (l - open)
				}
				left = l
				l = 0
			} else {
				l = 0
				left = 0
				open = 0
			}
		}
		fmt.Printf("%c ", s[i])
		fmt.Println(maxLen, left, l, open)
	}
	return maxLen
}

// @lc code=end
