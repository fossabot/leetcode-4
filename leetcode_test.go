package leetcode

import "testing"

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
