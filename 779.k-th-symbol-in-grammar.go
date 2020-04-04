/*
 * @lc app=leetcode id=779 lang=golang
 *
 * [779] K-th Symbol in Grammar
 *
 * https://leetcode.com/problems/k-th-symbol-in-grammar/description/
 *
 * algorithms
 * Medium (36.97%)
 * Likes:    346
 * Dislikes: 121
 * Total Accepted:    25.7K
 * Total Submissions: 69.6K
 * Testcase Example:  '1\n1'
 *
 * On the first row, we write a 0. Now in every subsequent row, we look at the
 * previous row and replace each occurrence of 0 with 01, and each occurrence
 * of 1 with 10.
 *
 * Given row N and index K, return the K-th indexed symbol in row N. (The
 * values of K are 1-indexed.) (1 indexed).
 *
 *
 * Examples:
 * Input: N = 1, K = 1
 * Output: 0
 *
 * Input: N = 2, K = 1
 * Output: 0
 *
 * Input: N = 2, K = 2
 * Output: 1
 *
 * Input: N = 4, K = 5
 * Output: 1
 *
 * Explanation:
 * row 1: 0
 * row 2: 01
 * row 3: 0110
 * row 4: 01101001
 *
 *
 * Note:
 *
 *
 * N will be an integer in the range [1, 30].
 * K will be an integer in the range [1, 2^(N-1)].
 *
 *
 */

package leetcode

// https://leetcode.com/problems/k-th-symbol-in-grammar/discuss/113697/My-3-lines-C%2B%2B-recursive-solution

// @lc code=start
func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	if K%2 == 0 {
		pre := kthGrammar(N-1, K/2)
		if pre == 0 {
			return 1
		}
		return 0
	}
	return kthGrammar(N-1, (K+1)/2)
}

// @lc code=end
