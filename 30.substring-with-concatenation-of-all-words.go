/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 *
 * https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (24.49%)
 * Likes:    725
 * Dislikes: 1118
 * Total Accepted:    162.3K
 * Total Submissions: 655.6K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * You are given a string, s, and a list of words, words, that are all of the
 * same length. Find all starting indices of substring(s) in s that is a
 * concatenation of each word in words exactly once and without any intervening
 * characters.
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 * ⁠ s = "barfoothefoobarman",
 * ⁠ words = ["foo","bar"]
 * Output: [0,9]
 * Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar"
 * respectively.
 * The output order does not matter, returning [9,0] is fine too.
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ⁠ s = "wordgoodgoodgoodbestword",
 * ⁠ words = ["word","good","best","word"]
 * Output: []
 *
 *
 */

package leetcode

// @lc code=start
func findSubstring(s string, words []string) []int {
	lenWords, lenStr := len(words), len(s)
	if lenWords == 0 || lenStr == 0 {
		return []int{}
	}
	lenWord := len(words[0])
	if lenWord == 0 {
		return []int{}
	}
	solution := make([]int, 0)
	for i := 0; i < lenStr; i++ {
		wordMap := make(map[string]int)
		for _, word := range words {
			wordMap[word]++
		}
		k := 0
		for j := i; j <= (lenStr-lenWord) && k < lenWords; k++ {
			curWord := s[j : j+lenWord]
			if count, ok := wordMap[curWord]; ok {
				if count == 1 {
					delete(wordMap, curWord)
				} else {
					wordMap[curWord]--
				}
				j = j + lenWord
			} else {
				break
			}
		}
		if k == lenWords {
			solution = append(solution, i)
		}
	}
	return solution
}

// @lc code=end
