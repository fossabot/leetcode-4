/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 *
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (44.67%)
 * Likes:    4218
 * Dislikes: 157
 * Total Accepted:    558.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * Given a 2d grid map of '1's (land) and '0's (water), count the number of
 * islands. An island is surrounded by water and is formed by connecting
 * adjacent lands horizontally or vertically. You may assume all four edges of
 * the grid are all surrounded by water.
 *
 * Example 1:
 *
 *
 * Input:
 * 11110
 * 11010
 * 11000
 * 00000
 *
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input:
 * 11000
 * 11000
 * 00100
 * 00011
 *
 * Output: 3
 *
 */

package leetcode

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	nRow, nCol := len(grid), len(grid[0])
	count := 1
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if grid[i][j] == '1' {
				// For assigning island value
				count++
				fill(grid, i, j, nRow, nCol, count)
			}
		}
	}

	// Started with 1
	return count - 1
}

func fill(grid [][]byte, x, y, X, Y, val int) {
	if x < 0 || y < 0 || x == X || y == Y {
		return
	}
	if grid[x][y] == '0' || grid[x][y] > '1' {
		return
	}
	grid[x][y] = byte(val) + '0'
	fill(grid, x-1, y, X, Y, val)
	fill(grid, x, y-1, X, Y, val)
	fill(grid, x+1, y, X, Y, val)
	fill(grid, x, y+1, X, Y, val)
}

// @lc code=end
