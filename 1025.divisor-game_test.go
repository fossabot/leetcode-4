/*
 * @lc app=leetcode id=1025 lang=golang
 *
 * [1025] Divisor Game
 *
 * https://leetcode.com/problems/divisor-game/description/
 *
 * algorithms
 * Easy (65.45%)
 * Likes:    298
 * Dislikes: 850
 * Total Accepted:    44.9K
 * Total Submissions: 68.5K
 * Testcase Example:  '2'
 *
 * Alice and Bob take turns playing a game, with Alice starting first.
 *
 * Initially, there is a number N on the chalkboard.  On each player's turn,
 * that player makes a move consisting of:
 *
 *
 * Choosing any x with 0 < x < N and N % x == 0.
 * Replacing the number N on the chalkboard with N - x.
 *
 *
 * Also, if a player cannot make a move, they lose the game.
 *
 * Return True if and only if Alice wins the game, assuming both players play
 * optimally.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: true
 * Explanation: Alice chooses 1, and Bob has no more moves.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: 3
 * Output: false
 * Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more
 * moves.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= N <= 1000
 *
 *
 *
 */

package leetcode

import "testing"

func Test_divisorGame(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{2},
			want: true,
		},
		{
			name: "Test 2",
			args: args{3},
			want: false,
		},
		{
			name: "Test 3",
			args: args{4},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorGame(tt.args.N); got != tt.want {
				t.Errorf("divisorGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
