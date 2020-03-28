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

// @lc code=start
func minCut(s string) int {
	lenStr := len(s)
	if lenStr == 0 {
		return 0
	}
	leftStart, rightStart := 0, lenStr-1
	leftPartition := 0

leftToRight:
	for {
		left, right := leftStart, rightStart
		for s[left] == s[right] {
			if (right - left) <= 1 {
				rightStart = leftStart - 1
				if rightStart < 0 {
					break leftToRight
				}
				leftPartition++
				leftStart = 0
				left = leftStart
				right = rightStart
			} else {
				left++
				right--
			}
		}
		leftStart++
	}

	leftStart, rightStart = 0, lenStr-1
	rightPartition := 0
rightToLeft:
	for {
		left, right := leftStart, rightStart
		for s[left] == s[right] {
			if (right - left) <= 1 {
				leftStart = rightStart + 1
				if leftStart >= lenStr {
					break rightToLeft
				}
				rightPartition++
				rightStart = lenStr - 1
				left = leftStart
				right = rightStart
			} else {
				left++
				right--
			}
		}
		rightStart--
	}

	if leftPartition < rightPartition {
		return leftPartition
	}
	return rightPartition
}

// @lc code=end
