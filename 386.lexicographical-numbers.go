/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 *
 * https://leetcode.com/problems/lexicographical-numbers/description/
 *
 * algorithms
 * Medium (49.34%)
 * Likes:    496
 * Dislikes: 66
 * Total Accepted:    49.1K
 * Total Submissions: 99K
 * Testcase Example:  '13'
 *
 * Given an integer n, return 1 - n in lexicographical order.
 *
 * For example, given 13, return: [1,10,11,12,13,2,3,4,5,6,7,8,9].
 *
 * Please optimize your algorithm to use less time and space. The input size
 * may be as large as 5,000,000.
 *
 */

package leetcode

// @lc code=start
func lexicalOrder(n int) []int {
	sol := make([]int, 0)
	st := intStack{}
	st = st.Push(1)
	for st.Len() != 0 {
		var x int
		st, x = st.Pop()
		sol = append(sol, x)
		if x*10 <= n {
			st = st.Push(x * 10)
		} else {
			for limit := x + 10; x < n && x < limit; {
				x++
				sol = append(sol, x)
			}
		}
	}
	return sol
}

type intStack []int

func (s intStack) Push(v int) intStack {
	return append(s, v)
}

func (s intStack) Pop() (intStack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s intStack) Len() int {
	return len(s)
}

// @lc code=end
