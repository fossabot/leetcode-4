/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (37.72%)
 * Likes:    3803
 * Dislikes: 190
 * Total Accepted:    792.8K
 * Total Submissions: 2.1M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */

package leetcode

// @lc code=start

// Stack is a data type for Stack daata structuire
type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func stackNew() *Stack {
	return &Stack{nil, 0}
}

// Len returns the number of items in the stack
func (object *Stack) Len() int {
	return object.length
}

// Peek shows the top item on the stack
func (object *Stack) Peek() interface{} {
	if object.length == 0 {
		return nil
	}
	return object.top.value
}

// Pop the top item of the stack and return it
func (object *Stack) Pop() interface{} {
	if object.length == 0 {
		return nil
	}

	n := object.top
	object.top = n.prev
	object.length--
	return n.value
}

// Push a value onto the top of the stack
func (object *Stack) Push(value interface{}) {
	n := &node{value, object.top}
	object.top = n
	object.length++
}

func isValid(s string) bool {
	data := stackNew()
	for _, char := range s {
		switch char {
		case '(':
			data.Push(')')
		case '{':
			data.Push('}')
		case '[':
			data.Push(']')

		case ')':
			fallthrough
		case '}':
			fallthrough
		case ']':
			if data.Len() == 0 {
				return false
			}
			ch := data.Pop()
			if ch != char {
				return false
			}
		}
	}
	if data.Len() != 0 {
		return false
	}
	return true
}

// @lc code=end
