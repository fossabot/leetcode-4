/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 *
 * https://leetcode.com/problems/fibonacci-number/description/
 *
 * algorithms
 * Easy (66.93%)
 * Likes:    486
 * Dislikes: 182
 * Total Accepted:    166K
 * Total Submissions: 248K
 * Testcase Example:  '2'
 *
 * The Fibonacci numbers, commonly denoted F(n) form a sequence, called the
 * Fibonacci sequence, such that each number is the sum of the two preceding
 * ones, starting from 0 and 1. That is,
 *
 *
 * F(0) = 0,   F(1) = 1
 * F(N) = F(N - 1) + F(N - 2), for N > 1.
 *
 *
 * Given N, calculate F(N).
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: 1
 * Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: 2
 * Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
 *
 *
 * Example 3:
 *
 *
 * Input: 4
 * Output: 3
 * Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
 *
 *
 *
 *
 * Note:
 *
 * 0 ≤ N ≤ 30.
 *
 */

package leetcode

// @lc code=start

var fibMap = map[int]int{}

func fib(N int) int {
	if N < 2 {
		_, found := fibMap[N]
		if !found {
			fibMap[N] = N
		}
		return fibMap[N]
	}
	fibN1, found := fibMap[N-1]
	if !found {
		fibN1 = fib(N - 1)
	}
	fibN2, found := fibMap[N-2]
	if !found {
		fibN2 = fib(N - 2)
	}
	fibMap[N] = fibN1 + fibN2
	return fibMap[N]
}

// @lc code=end
