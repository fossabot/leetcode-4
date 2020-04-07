/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 *
 * https://leetcode.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (43.84%)
 * Likes:    1502
 * Dislikes: 58
 * Total Accepted:    209K
 * Total Submissions: 461.2K
 * Testcase Example:  '"aab"'
 *
 * Given a string s, partition s such that every substring of the partition is
 * a palindrome.
 *
 * Return all possible palindrome partitioning of s.
 *
 * Example:
 *
 *
 * Input: "aab"
 * Output:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 *
 *
 */

package leetcode

// @lc code=start
func partition(s string) [][]string {
	lenS := len(s)
	if lenS == 0 {
		return [][]string{}
	}
	return generatePartition(s, 0, lenS-1)
}

func generatePartition(s string, start, end int) [][]string {
	if start == end {
		return [][]string{{string(s[start])}}
	}

	res := [][]string{}
	for i := start; i < end; i++ {
		leftRes := generatePartition(s, start, i)
		rightRes := generatePartition(s, i, end)
		for _, left := range leftRes {
			for _, right := range rightRes {
				sol := append(left, right...)
				res = append(res, sol)
			}
		}
	}
	return res
}

// @lc code=end
