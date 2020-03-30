/*
 * @lc app=leetcode id=132 lang=golang
 *
 * [132] Palindrome Partitioning II
 *
 * https://leetcode.com/problems/palindrome-partitioning-ii/description/
 *
 * algorithms
 * Hard (28.82%)
 * Likes:    878
 * Dislikes: 29
 * Total Accepted:    122.4K
 * Total Submissions: 415.8K
 * Testcase Example:  '"aab"'
 *
 * Given a string s, partition s such that every substring of the partition is
 * a palindrome.
 *
 * Return the minimum cuts needed for a palindrome partitioning of s.
 *
 * Example:
 *
 *
 * Input: "aab"
 * Output: 1
 * Explanation: The palindrome partitioning ["aa","b"] could be produced using
 * 1 cut.
 *
 *
 */

package leetcode

import "math"

// @lc code=start
func minCut(s string) int {
	lenStr := len(s)
	if lenStr == 0 {
		return 0
	}

	dp := make([][2]int, lenStr)
	minC := math.MaxInt64

	dp[0][0] = 1
	dp[0][1] = 0

	// for i := 1; i < lenStr; i++ {
	// 	backIndex := i - dp[i-1][0] - 1
	// 	if[backIndex] == s[i] {

	// 	}
	// 	if dp[i][0] < minC {
	// 		minC = dp[i][0]
	// 	}
	// }
	return minC
}

// @lc code=end
