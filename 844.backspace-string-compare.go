/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 *
 * https://leetcode.com/problems/backspace-string-compare/description/
 *
 * algorithms
 * Easy (46.87%)
 * Likes:    1194
 * Dislikes: 65
 * Total Accepted:    126.6K
 * Total Submissions: 267.3K
 * Testcase Example:  '"ab#c"\n"ad#c"'
 *
 * Given two strings S and T, return if they are equal when both are typed into
 * empty text editors. # means a backspace character.
 *
 *
 * Example 1:
 *
 *
 * Input: S = "ab#c", T = "ad#c"
 * Output: true
 * Explanation: Both S and T become "ac".
 *
 *
 *
 * Example 2:
 *
 *
 * Input: S = "ab##", T = "c#d#"
 * Output: true
 * Explanation: Both S and T become "".
 *
 *
 *
 * Example 3:
 *
 *
 * Input: S = "a##c", T = "#a#c"
 * Output: true
 * Explanation: Both S and T become "c".
 *
 *
 *
 * Example 4:
 *
 *
 * Input: S = "a#c", T = "b"
 * Output: false
 * Explanation: S becomes "c" while T becomes "b".
 *
 *
 * Note:
 *
 *
 * 1 <= S.length <= 200
 * 1 <= T.length <= 200
 * S and T only contain lowercase letters and '#' characters.
 *
 *
 * Follow up:
 *
 *
 * Can you solve it in O(N) time and O(1) space?
 *
 *
 *
 *
 *
 *
 */

package leetcode

// @lc code=start
func backspaceCompare(S string, T string) bool {
	return backspaceCompareIterative(S, T)
}

func backspaceCompareIterative(S string, T string) bool {
	posS, posT := len(S)-1, len(T)-1
	for {
		posS = nextPos(S, posS)
		posT = nextPos(T, posT)

		if posS < 0 && posT < 0 {
			return true
		} else if posS < 0 || posT < 0 {
			return false
		}

		if S[posS] != T[posT] {
			return false
		}
		posS--
		posT--
	}
}

func nextPos(S string, pos int) int {
	bsCount := 0

	for {
		if pos < 0 {
			return pos
		}

		if S[pos] == '#' {
			bsCount--
			pos--
		} else {
			if bsCount == 0 {
				return pos
			}
			bsCount++
			pos--
		}
	}
}

func backspaceCompareRecursive(S string, T string) bool {
	resS, resT := backspaceStr([]byte(S)), backspaceStr([]byte(T))
	return string(resS) == string(resT)
}

func backspaceStr(b []byte) []byte {
	byteLen := len(b)
	if byteLen == 0 {
		return b
	}

	lastByte := b[byteLen-1]
	res := backspaceStr(b[:byteLen-1])
	return checkByte(lastByte, res)
}

func checkByte(lastByte byte, res []byte) []byte {
	if lastByte == '#' {
		resLen := len(res)
		if resLen != 0 {
			res = res[:resLen-1]
		}
		return res
	}
	res = append(res, lastByte)
	return res
}

// @lc code=end
