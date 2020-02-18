/*
 * @lc app=leetcode id=1319 lang=golang
 *
 * [1319] Number of Operations to Make Network Connected
 *
 * https://leetcode.com/problems/number-of-operations-to-make-network-connected/description/
 *
 * algorithms
 * Medium (50.53%)
 * Likes:    162
 * Dislikes: 4
 * Total Accepted:    7.8K
 * Total Submissions: 15.5K
 * Testcase Example:  '4\n[[0,1],[0,2],[1,2]]'
 *
 * There are n computers numbered from 0 to n-1 connected by ethernet cables
 * connections forming a network where connections[i] = [a, b] represents a
 * connection between computers a and b. Any computer can reach any other
 * computer directly or indirectly through the network.
 *
 * Given an initial computer network connections. You can extract certain
 * cables between two directly connected computers, and place them between any
 * pair of disconnected computers to make them directly connected. Return the
 * minimum number of times you need to do this in order to make all the
 * computers connected. If it's not possible, return -1.
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: n = 4, connections = [[0,1],[0,2],[1,2]]
 * Output: 1
 * Explanation: Remove cable between computer 1 and 2 and place between
 * computers 1 and 3.
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
 * Output: 2
 *
 *
 * Example 3:
 *
 *
 * Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
 * Output: -1
 * Explanation: There are not enough cables.
 *
 *
 * Example 4:
 *
 *
 * Input: n = 5, connections = [[0,1],[0,2],[3,4],[2,3]]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 10^5
 * 1 <= connections.length <= min(n*(n-1)/2, 10^5)
 * connections[i].length == 2
 * 0 <= connections[i][0], connections[i][1] < n
 * connections[i][0] != connections[i][1]
 * There are no repeated connections.
 * No two computers are connected by more than one cable.
 *
 */

package leetcode

// @lc code=start
type UnionFind struct {
	root []int
	size []int
}

// New returns an initialized list of size
func New(size int) *UnionFind {
	return new(UnionFind).init(size)
}

// Constructor initializes root and size arrays
func (uf *UnionFind) init(size int) *UnionFind {
	uf = new(UnionFind)
	uf.root = make([]int, size)
	uf.size = make([]int, size)

	for i := 0; i < size; i++ {
		uf.root[i] = i
		uf.size[i] = 1
	}

	return uf
}

// Union connects p and q by finding their roots and comparing their respective
// size arrays to keep the tree flat
func (uf *UnionFind) Union(p int, q int) {
	qRoot := uf.Root(q)
	pRoot := uf.Root(p)

	if uf.size[qRoot] < uf.size[pRoot] {
		uf.root[qRoot] = uf.root[pRoot]
		uf.size[pRoot] += uf.size[qRoot]
	} else {
		uf.root[pRoot] = uf.root[qRoot]
		uf.size[qRoot] += uf.size[pRoot]
	}
}

// Root or Find traverses each parent element while compressing the
// levels to find the root element of p
// If we attempt to access an element outside the array it returns -1
func (uf *UnionFind) Root(p int) int {
	if p > len(uf.root)-1 {
		return -1
	}

	for uf.root[p] != p {
		uf.root[p] = uf.root[uf.root[p]]
		p = uf.root[p]
	}

	return p
}

// Root or Find
func (uf *UnionFind) Find(p int) int {
	return uf.Root(p)
}

// Check if items p,q are connected
func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.Root(p) == uf.Root(q)
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < (n - 1) {
		return -1
	}

	uf := New(n)
	var count int
	for _, edge := range connections {
		vertex1 := edge[0]
		vertex2 := edge[1]
		if uf.Connected(vertex1, vertex2) {
			count++
		}
		uf.Union(vertex1, vertex2)
	}
	return count - (len(connections) - n + 1)
}

// @lc code=end
