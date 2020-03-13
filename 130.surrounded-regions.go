/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 *
 * https://leetcode.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (25.26%)
 * Likes:    1199
 * Dislikes: 548
 * Total Accepted:    188.3K
 * Total Submissions: 739K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * Given a 2D board containing 'X' and 'O' (the letter O), capture all regions
 * surrounded by 'X'.
 *
 * A region is captured by flipping all 'O's into 'X's in that surrounded
 * region.
 *
 * Example:
 *
 *
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 *
 *
 * After running your function, the board should be:
 *
 *
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 *
 *
 * Explanation:
 *
 * Surrounded regions shouldn’t be on the border, which means that any 'O' on
 * the border of the board are not flipped to 'X'. Any 'O' that is not on the
 * border and it is not connected to an 'O' on the border will be flipped to
 * 'X'. Two cells are connected if they are adjacent cells connected
 * horizontally or vertically.
 *
 */

package leetcode

// @lc code=start
func solve(board [][]byte) {
	nX := len(board)
	if nX == 0 {
		return
	}
	nY := len(board[0])
	visited := make([][]bool, nX)
	for x := 0; x < nX; x++ {
		visited[x] = make([]bool, nY)
	}
	for x := 0; x < nX; x++ {
		for y := 0; y < nY; y++ {
			if !visited[x][y] {
				if board[x][y] == 'X' {
					visited[x][y] = true
				} else {
					visit(board, visited, x, y, nX, nY)
					visited[x][y] = true
				}
			}
		}
	}
}

func visit(board [][]byte, visited [][]bool, x, y, nX, nY int) {
	if visited[x][y] || x == 0 || x == nX-1 || y == 0 || y == nY-1 {
		visited[x][y] = true
		return
	}

	visit(board, visited, x+1, y, nX, nY)
	visit(board, visited, x, y+1, nX, nY)
	if board[x-1][y] != 'O' && board[x+1][y] != 'O' && board[x][y-1] != 'O' && board[x][y+1] != 'O' {
		board[x][y] = 'X'
	}
	visited[x][y] = true
}

// @lc code=end
