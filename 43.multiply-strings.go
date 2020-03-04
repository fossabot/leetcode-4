/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 *
 * https://leetcode.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (32.58%)
 * Likes:    1462
 * Dislikes: 673
 * Total Accepted:    259.7K
 * Total Submissions: 794.4K
 * Testcase Example:  '"2"\n"3"'
 *
 * Given two non-negative integers num1 and num2 represented as strings, return
 * the product of num1 and num2, also represented as a string.
 *
 * Example 1:
 *
 *
 * Input: num1 = "2", num2 = "3"
 * Output: "6"
 *
 * Example 2:
 *
 *
 * Input: num1 = "123", num2 = "456"
 * Output: "56088"
 *
 *
 * Note:
 *
 *
 * The length of both num1 and num2 is < 110.
 * Both num1 and num2 contain only digits 0-9.
 * Both num1 and num2 do not contain any leading zero, except the number 0
 * itself.
 * You must not use any built-in BigInteger library or convert the inputs to
 * integer directly.
 *
 *
 */

package leetcode

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	res := make([]byte, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			val := (num1[i] - '0') * (num2[j] - '0')
			val += (res[i+j+1])
			if val >= 10 {
				res[i+j] += (val / 10)
				val %= 10
			}
			res[i+j+1] = val
		}
	}
	if res[0] == 0 {
		res = res[1:]
	}
	for i := range res {
		res[i] += '0'
	}
	return string(res)
}

// @lc code=end
