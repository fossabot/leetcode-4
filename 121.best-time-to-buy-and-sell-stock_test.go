/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (48.88%)
 * Likes:    4382
 * Dislikes: 195
 * Total Accepted:    780.3K
 * Total Submissions: 1.6M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * Say you have an array for which the i^th element is the price of a given
 * stock on day i.
 *
 * If you were only permitted to complete at most one transaction (i.e., buy
 * one and sell one share of the stock), design an algorithm to find the
 * maximum profit.
 *
 * Note that you cannot sell a stock before you buy one.
 *
 * Example 1:
 *
 *
 * Input: [7,1,5,3,6,4]
 * Output: 5
 * Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit
 * = 6-1 = 5.
 * Not 7-1 = 6, as selling price needs to be larger than buying price.
 *
 *
 * Example 2:
 *
 *
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 *
 *
 */

package leetcode

import "testing"

func Test_maxProfitI(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[]int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{[]int{7, 6, 4, 3, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitI(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitI() = %v, want %v", got, tt.want)
			}
		})
	}
}
