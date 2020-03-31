/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 *
 * https://leetcode.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (32.80%)
 * Likes:    3211
 * Dislikes: 107
 * Total Accepted:    338.2K
 * Total Submissions: 996.8K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * You are given coins of different denominations and a total amount of money
 * amount. Write a function to compute the fewest number of coins that you need
 * to make up that amount. If that amount of money cannot be made up by any
 * combination of the coins, return -1.
 *
 * Example 1:
 *
 *
 * Input: coins = [1, 2, 5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 *
 * Example 2:
 *
 *
 * Input: coins = [2], amount = 3
 * Output: -1
 *
 *
 * Note:
 * You may assume that you have an infinite number of each kind of coin.
 *
 */

package leetcode

import (
	"sort"
)

// @lc code=start
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	minCoin := make([]int, amount+1)

	for _, coin := range coins {
		if coin > amount {
			break
		}
		minCoin[coin] = 1
		for i := 1; i <= amount; i++ {
			if minCoin[i] == 0 {
				continue
			}
			newAmount := i + coin
			if newAmount > amount {
				break
			}
			coinCount := minCoin[i] + 1
			if minCoin[newAmount] == 0 || minCoin[newAmount] > coinCount {
				minCoin[newAmount] = coinCount
			}
		}
	}

	if amount!=0 && minCoin[amount] == 0 {
		return -1
	}

	return minCoin[amount]
}

// @lc code=end
