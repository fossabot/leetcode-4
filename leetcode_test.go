package leetcode

import (
	"reflect"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	str := "abc"
	val := lengthOfLongestSubstring(str)
	if val != 3 {
		t.Error("Value not matched, got value ", val)
	}
}

func TestFindRedundantConnection(t *testing.T) {
	input1 := [][]int{{1, 2}, {1, 3}, {2, 3}}
	expectedResult1 := []int{2, 3}
	input2 := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	expectedResult2 := []int{1, 4}
	actualResult1 := findRedundantConnection(input1)
	actualResult2 := findRedundantConnection(input2)
	if actualResult1[0] != expectedResult1[0] && actualResult1[1] !=
		expectedResult1[1] && actualResult2[0] != expectedResult2[0] &&
		actualResult2[1] != expectedResult2[1] {
		t.Error("Value not matched. Expected [2,3] Got Result", actualResult1)
		t.Error("Value not matched. Expected [1,4] Got Result", actualResult2)
	}
}

// func TestCriticalConnections(t *testing.T) {
// 	input1 := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}
// 	expectedResult1 := [][]int{{1, 3}}
// 	actualResult1 := criticalConnections(4, input1)
// 	if actualResult1[0][0] != expectedResult1[0][0] && actualResult1[0][1] !=
// 		expectedResult1[0][1] {
// 		t.Error("Value not matched. Expected [1,3] Got Result", actualResult1)
// 	}
// }

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input []int
		want  [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{0}, [][]int{}},
		{[]int{-1, 0, 1, 0}, [][]int{{-1, 0, 1}}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{1, 2, -2, -1}, [][]int{}},
		{[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, [][]int{
			{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2},
		}},
	}

	for _, test := range tests {
		if got := threeSum(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("threeSum(%v) = %v, want %v\n", test.input, got, test.want)
		}
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{}},
		{[]int{2, 7, 11, 15}, []int{0, 1}},
		// {[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		// {[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		// {[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		// {[]int{1, 2, -2, -1}, [][]int{}},
		// {[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		// {[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, [][]int{
		// 	{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2},
		// }},
	}

	for _, test := range tests {
		// test := tests[2]
		// {
		if got := twoSum(test.input, 9); !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum(%v) = %v, want %v\n", test.input, got, test.want)
		}
	}
}
