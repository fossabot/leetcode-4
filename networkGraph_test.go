package leetcode

import (
	"reflect"
	"testing"
)

func Test_getGraph(t *testing.T) {
	type args struct {
		edges [][]int
		N     int
	}
	tests := []struct {
		name string
		args args
		want [][][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGraph(tt.args.edges, tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printGraph(t *testing.T) {
	type args struct {
		graph [][][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printGraph(tt.args.graph)
		})
	}
}

func Test_pqGraph_Len(t *testing.T) {
	tests := []struct {
		name string
		pq   pqGraph
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Len(); got != tt.want {
				t.Errorf("pqGraph.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pqGraph_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		pq   pqGraph
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("pqGraph.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pqGraph_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		pq   pqGraph
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_pqGraph_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		pq   *pqGraph
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Push(tt.args.x)
		})
	}
}

func Test_pqGraph_Pop(t *testing.T) {
	tests := []struct {
		name string
		pq   *pqGraph
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pqGraph.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pqGraph_update(t *testing.T) {
	type args struct {
		item     *pqGraphNode
		priority int
	}
	tests := []struct {
		name string
		pq   *pqGraph
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.update(tt.args.item, tt.args.priority)
		})
	}
}
