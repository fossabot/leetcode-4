/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 *
 * https://leetcode.com/problems/word-break/description/
 *
 * algorithms
 * Medium (37.57%)
 * Likes:    3675
 * Dislikes: 201
 * Total Accepted:    493.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * Given a non-empty string s and a dictionary wordDict containing a list of
 * non-empty words, determine if s can be segmented into a space-separated
 * sequence of one or more dictionary words.
 *
 * Note:
 *
 *
 * The same word in the dictionary may be reused multiple times in the
 * segmentation.
 * You may assume the dictionary does not contain duplicate words.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "leetcode", wordDict = ["leet", "code"]
 * Output: true
 * Explanation: Return true because "leetcode" can be segmented as "leet
 * code".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "applepenapple", wordDict = ["apple", "pen"]
 * Output: true
 * Explanation: Return true because "applepenapple" can be segmented as "apple
 * pen apple".
 * Note that you are allowed to reuse a dictionary word.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * Output: false
 *
 *
 */

package leetcode

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	return wordForm(s, "", wordDict)
}

func wordForm(s string, input string, wordDict []string) bool {
	if len(input) == len(s) {
		return true
	}
	for _, word := range wordDict {
		curWord := input + word
		res := chkStr(s, curWord)
		if res {
			return res
		}
	}
	return false
}

func chkStr(s, t string) bool {
	lenS, lenT := len(s), len(t)
	if lenS >= lenT {
		return s[:lenT] == t
	}
	return false
}

// @lc code=end
