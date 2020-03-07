package contest

import "sort"

func sortString(s string) string {
	b := []byte(s)
	// Reverse sort the string array from big to small
	sort.Slice(b, func(i int, j int) bool { return b[i] > b[j] })

	leftStack := stack{}
	left := 0

	for i := 1; i < len(b); i++ {
		if b[i] != b[left] {
			// Put the byte array with smallest value on top
			leftStack = leftStack.Push(b[left:i])
			left = i
		}
	}

	leftStack = leftStack.Push(b[left:])

	res := []byte{}
	for leftStack.Len() != 0 {
		rightStack := stack{}
		for leftStack.Len() != 0 {
			x := []byte{}
			leftStack, x = leftStack.Pop()
			res = append(res, x[0])
			if len(x) != 1 {
				rightStack = rightStack.Push(x[1:])
			}
		}
		for rightStack.Len() != 0 {
			x := []byte{}
			rightStack, x = rightStack.Pop()
			res = append(res, x[0])
			if len(x) != 1 {
				leftStack = leftStack.Push(x[1:])
			}
		}
	}
	return string(res)
}

type stack [][]byte

func (s stack) Push(v []byte) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, []byte) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Len() int {
	return len(s)
}
