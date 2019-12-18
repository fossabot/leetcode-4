/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.57%)
 * Likes:    2680
 * Dislikes: 4173
 * Total Accepted:    891.1K
 * Total Submissions: 3.5M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 *
 */

package reverse

import "strconv"
import "math"

// @lc code=start
func reverse(x int) int {
	negative := false
	pow2_31 := (math.Pow(2, 31) - 1)
	if math.Abs(float64(x)) > pow2_31 {
		return 0
	}
	if x < 0 {
		x = (0 - x)
		negative = true
	}
	inputIntToString := strconv.Itoa(x)
	runes := []rune(inputIntToString)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reverseInt, _ := strconv.Atoi(string(runes))
	if negative {
		reverseInt = (0 - reverseInt)
	}
	if math.Abs(float64(reverseInt)) > pow2_31 {
		return 0
	}
	return reverseInt
}

// @lc code=end
