/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (33.10%)
 * Likes:    2886
 * Dislikes: 148
 * Total Accepted:    402.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given a 2D board and a word, find if the word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cell,
 * where "adjacent" cells are those horizontally or vertically neighboring. The
 * same letter cell may not be used more than once.
 *
 * Example:
 *
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 *
 *
 *
 * Constraints:
 *
 *
 * board and word consists only of lowercase and uppercase English letters.
 * 1 <= board.length <= 200
 * 1 <= board[i].length <= 200
 * 1 <= word.length <= 10^3
 *
 *
 */

package leetcode

// @lc code=start
func exist(board [][]byte, word string) bool {
	X := len(board)
	if X == 0 {
		return false
	}
	Y, n := len(board[0]), len(word)
	if Y == 0 || n == 0 {
		return false
	}

	visited := make([][]bool, X)
	for x := 0; x < X; x++ {
		visited[x] = make([]bool, Y)
	}

	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			if board[x][y] == word[0] {
				st := byteStack{}
				st.Push(board[x][y])
				for !st.isEmpty(){

				}
			}
		}
	}
	return true
}

type byteStack []byte

func (s byteStack) Push(v byte) byteStack {
	return append(s, v)
}

func (s byteStack) Pop() (byteStack, byte) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s byteStack) Len() int {
	return len(s)
}

func (s byteStack) isEmpty() bool {
	return len(s) == 0
}

// @lc code=end
