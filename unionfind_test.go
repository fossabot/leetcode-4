package leetcode

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *UnionFind
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_init(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		uf   *UnionFind
		args args
		want *UnionFind
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uf.init(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionFind.init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Union(t *testing.T) {
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name string
		uf   *UnionFind
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.uf.Union(tt.args.p, tt.args.q)
		})
	}
}

func TestUnionFind_Root(t *testing.T) {
	type args struct {
		p int
	}
	tests := []struct {
		name string
		uf   *UnionFind
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uf.Root(tt.args.p); got != tt.want {
				t.Errorf("UnionFind.Root() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Find(t *testing.T) {
	type args struct {
		p int
	}
	tests := []struct {
		name string
		uf   *UnionFind
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uf.Find(tt.args.p); got != tt.want {
				t.Errorf("UnionFind.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Connected(t *testing.T) {
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name string
		uf   *UnionFind
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uf.Connected(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("UnionFind.Connected() = %v, want %v", got, tt.want)
			}
		})
	}
}
