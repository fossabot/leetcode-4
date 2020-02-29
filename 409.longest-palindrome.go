/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 *
 * https://leetcode.com/problems/longest-palindrome/description/
 *
 * algorithms
 * Easy (49.44%)
 * Likes:    764
 * Dislikes: 74
 * Total Accepted:    125.6K
 * Total Submissions: 253.7K
 * Testcase Example:  '"abccccdd"'
 *
 * Given a string which consists of lowercase or uppercase letters, find the
 * length of the longest palindromes that can be built with those letters.
 *
 * This is case sensitive, for example "Aa" is not considered a palindrome
 * here.
 *
 * Note:
 * Assume the length of given string will not exceed 1,010.
 *
 *
 * Example:
 *
 * Input:
 * "abccccdd"
 *
 * Output:
 * 7
 *
 * Explanation:
 * One longest palindrome that can be built is "dccaccd", whose length is 7.
 *
 *
 */

package leetcode

// @lc code=start
func longestPalindromeLength(s string) int {
	// charArray := rune(s)
	charMap := make(map[rune]int)
	for _, ch := range s {
		charMap[ch]++
	}

	result := 0
	odd := false
	for _, count := range charMap {
		if count%2 == 1 {
			odd = true
		}
		result += ((count / 2) * 2)
	}
	if odd {
		result++
	}
	return result
}

// @lc code=end
