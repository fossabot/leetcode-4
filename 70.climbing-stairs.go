/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 *
 * https://leetcode.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (46.34%)
 * Likes:    3598
 * Dislikes: 118
 * Total Accepted:    600.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '2'
 *
 * You are climbing a stair case. It takes n steps to reach to the top.
 *
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can
 * you climb to the top?
 *
 * note: Given n will be a positive integer.
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: 2
 * Explanation: There are two ways to climb to the top.
 * 1. 1 step + 1 step
 * 2. 2 steps
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: 3
 * Explanation: There are three ways to climb to the top.
 * 1. 1 step + 1 step + 1 step
 * 2. 1 step + 2 steps
 * 3. 2 steps + 1 step
 *
 *
 */

package leetcode

// @lc code=start

var steps = map[int]int{}

func climbStairs(n int) int {
	if n <= 2 {
		_, found := steps[n]
		if !found {
			steps[n] = n
		}
		return steps[n]
	}
	step1, found := steps[n-1]
	if !found {
		step1 = climbStairs(n - 1)
	}
	step2, found := steps[n-2]
	if !found {
		step2 = climbStairs(n - 2)
	}
	steps[n] = step1 + step2
	return steps[n]
}

// @lc code=end
