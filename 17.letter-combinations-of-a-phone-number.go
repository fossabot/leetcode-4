/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (44.70%)
 * Likes:    3185
 * Dislikes: 366
 * Total Accepted:    529K
 * Total Submissions: 1.2M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent.
 *
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 *
 * Example:
 *
 *
 * Input: "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * Note:
 *
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 *
 */

package leetcode

// @lc code=start
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	solution := []string{""}
	numMap := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	for _, digit := range digits {
		temp := solution
		chars := numMap[digit]
		solution = []string{}
		for _, tmpStr := range temp {
			for _, char := range chars {
				solution = append(solution, tmpStr+string(char))
			}
		}
	}
	return solution
}

// @lc code=end
