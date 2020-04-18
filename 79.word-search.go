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
	X, Y := len(board), len(board[0])
	wordLen := len(word)
	if X == 0 || Y == 0 || wordLen == 0 {
		return false
	}
	visited := make([][]bool, X)
	for i := 0; i < X; i++ {
		visited[i] = make([]bool, Y)
	}
	bWord := []byte(word)

	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if chkExist(board, visited, i, j, X, Y, bWord) {
				return true
			}
		}
	}

	return false
}

func chkExist(board [][]byte, visited [][]bool, x, y, X, Y int,
	byteArr []byte) bool {
	posValid := chkValidPos(visited, x, y, X, Y)
	if len(byteArr) == 1 {
		if posValid {
			return board[x][y] == byteArr[0]
		}
	}

	if posValid {
		visited[x][y] = true
		if board[x][y] == byteArr[0] {
			right := chkExist(board, visited, x, y+1, X, Y, byteArr[1:])
			left := chkExist(board, visited, x, y-1, X, Y, byteArr[1:])
			down := chkExist(board, visited, x+1, y, X, Y, byteArr[1:])
			up := chkExist(board, visited, x-1, y, X, Y, byteArr[1:])
			if right || left || up || down {
				return true
			}
		}
		visited[x][y] = false
	}
	return false
}

func chkValidPos(visited [][]bool, x, y, X, Y int) bool {
	if x >= 0 && x < X && y >= 0 && y < Y {
		// Visited value is negated because, a pos is valid only if it wasn't visited before
		return !visited[x][y]
	}
	return false
}

// @lc code=end
