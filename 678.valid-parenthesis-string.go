/*
 * @lc app=leetcode id=678 lang=golang
 *
 * [678] Valid Parenthesis String
 *
 * https://leetcode.com/problems/valid-parenthesis-string/description/
 *
 * algorithms
 * Medium (33.62%)
 * Likes:    1111
 * Dislikes: 33
 * Total Accepted:    52.5K
 * Total Submissions: 163.3K
 * Testcase Example:  '"()"'
 *
 *
 * Given a string containing only three types of characters: '(', ')' and '*',
 * write a function to check whether this string is valid. We define the
 * validity of a string by these rules:
 *
 * Any left parenthesis '(' must have a corresponding right parenthesis ')'.
 * Any right parenthesis ')' must have a corresponding left parenthesis '('.
 * Left parenthesis '(' must go before the corresponding right parenthesis ')'.
 * '*' could be treated as a single right parenthesis ')' or a single left
 * parenthesis '(' or an empty string.
 * An empty string is also valid.
 *
 *
 *
 * Example 1:
 *
 * Input: "()"
 * Output: True
 *
 *
 *
 * Example 2:
 *
 * Input: "(*)"
 * Output: True
 *
 *
 *
 * Example 3:
 *
 * Input: "(*))"
 * Output: True
 *
 *
 *
 * Note:
 *
 * The string size will be in the range [1, 100].
 *
 *
 */

package leetcode

// @lc code=start
func checkValidString(s string) bool {
	low, high := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			low++
			high++
		} else if s[i] == ')' {
			if low > 0 {
				low--
			}
			high--

		} else {
			if low > 0 {
				low--
			}
			high++
		}
		if high < 0 {
			return false
		}
	}
	return low == 0
}

// @lc code=end
