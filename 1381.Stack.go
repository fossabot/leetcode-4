// package leetcode

// type CustomStack struct {
// 	maxSize int
// 	size    int
// 	stack   []int
// }

// func Constructor(maxSize int) CustomStack {
// 	st := CustomStack{}
// 	st.stack = make([]int, maxSize)
// 	st.size = 0
// 	st.maxSize = maxSize
// 	return st
// }

// func (this *CustomStack) Push(x int) {
// 	if this.size == this.maxSize {
// 		return
// 	}
// 	this.stack[this.size] = x
// 	this.size++
// }

// func (this *CustomStack) Pop() int {
// 	if this.size == 0 {
// 		return -1
// 	}
// 	this.size--
// 	return this.stack[this.size]
// }

// func (this *CustomStack) Increment(k int, val int) {
// 	for i := 0; i < k && i < this.size; i++ {
// 		this.stack[i] += val
// 	}
// }

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
