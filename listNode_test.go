package leetcode

import (
	"reflect"
	"testing"
)

func Test_newList(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newList(tt.args.vals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equalList(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalList(tt.args.l1, tt.args.l2); got != tt.want {
				t.Errorf("equalList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printList(t *testing.T) {
	type args struct {
		l *ListNode
	}
	tests := []struct {
		name  string
		args  args
		wantS string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := printList(tt.args.l); gotS != tt.wantS {
				t.Errorf("printList() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func Test_pqLinkedList_Len(t *testing.T) {
	tests := []struct {
		name string
		pq   pqLinkedList
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Len(); got != tt.want {
				t.Errorf("pqLinkedList.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pqLinkedList_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		pq   pqLinkedList
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("pqLinkedList.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pqLinkedList_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		pq   pqLinkedList
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

func Test_pqLinkedList_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		pq   *pqLinkedList
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

func Test_pqLinkedList_Pop(t *testing.T) {
	tests := []struct {
		name string
		pq   *pqLinkedList
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pqLinkedList.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
