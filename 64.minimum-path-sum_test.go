/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (49.83%)
 * Likes:    2389
 * Dislikes: 52
 * Total Accepted:    341.4K
 * Total Submissions: 656.6K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right which minimizes the sum of all numbers along its path.
 *
 * Note: You can only move either down or right at any point in time.
 *
 * Example:
 *
 *
 * Input:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * Output: 7
 * Explanation: Because the path 1→3→1→1→1 minimizes the sum.
 *
 *
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    nodeHeap
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("intHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intHeap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    nodeHeap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("intHeap.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intHeap_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    nodeHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_intHeap_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		h    *nodeHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.x)
		})
	}
}

func Test_intHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *nodeHeap
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dp(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dp(tt.args.grid); got != tt.want {
				t.Errorf("dp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dijkstra(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dijkstra(tt.args.grid); got != tt.want {
				t.Errorf("dijkstra() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodeHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    nodeHeap
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("nodeHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodeHeap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    nodeHeap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("nodeHeap.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodeHeap_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    nodeHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_nodeHeap_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		h    *nodeHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.x)
		})
	}
}

func Test_nodeHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *nodeHeap
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodeHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
