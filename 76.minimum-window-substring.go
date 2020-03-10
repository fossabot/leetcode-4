/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 *
 * https://leetcode.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (33.26%)
 * Likes:    3584
 * Dislikes: 252
 * Total Accepted:    337.7K
 * Total Submissions: 1M
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * Given a string S and a string T, find the minimum window in S which will
 * contain all the characters in T in complexity O(n).
 *
 * Example:
 *
 *
 * Input: S = "ADOBECODEBANC", T = "ABC"
 * Output: "BANC"
 *
 *
 * Note:
 *
 *
 * If there is no such window in S that covers all characters in T, return the
 * empty string "".
 * If there is such window, you are guaranteed that there will always be only
 * one unique minimum window in S.
 *
 *
 */

package leetcode

// @lc code=start
type char struct {
	ch    rune
	index int
}

func minWindow(s string, t string) string {
	// Return if any of the strings are empty
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	left, right := 0, 0

	// left & right index initialized to 0
	// Minimum length to length of string + 1
	// ValidLength is the number of character to match
	// Current Map to increase value when right includes a new character and
	// decreases when left passes a new character
	for l, r, minLen, valL, curMap := 0, 0, len(s)+1, len(t),
		make(map[byte]int); r < len(s); r++ {
		rch := s[r]

		if _, ok := tMap[rch]; ok {
			curMap[rch]++
			if curMap[rch] <= tMap[rch] {
				valL--
			}
		}

		// Enter the loop when all the characters are found
		for ; valL == 0; l++ {
			lch := s[l]
			val, ok := tMap[lch]
			if !ok || curMap[lch] > val {
				if curMap[lch] > val {
					curMap[lch]--
				}
			} else {
				length := r - l + 1
				if length < minLen {
					left, right = l, r+1
					minLen = length
				}
				// decreasing current character and exiting loop by incrementing
				// loop condition
				curMap[lch]--
				valL++
			}
		}
	}

	return s[left:right]
}

// @lc code=end
