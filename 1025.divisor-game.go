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

// @lc code=start
func divisorGame(N int) bool {
	// arr := make([]bool, N+1)
	// for i := 1; i <= N; i++ {
	// 	for j := 1; j <= i/2; j++ {
	// 		if (i%j) == 0 && !arr[i-j] {
	// 			arr[i] = true
	// 			break
	// 		}
	// 	}
	// }
	// return arr[N]
	return N%2 == 0
}

// @lc code=end
