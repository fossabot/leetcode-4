package leetcode

import "testing"

func TestLongestSubstring(t *testing.T) {
	str := "abc"
	val := lengthOfLongestSubstring(str)
	if val != 3 {
		t.Error("Value not matched, got value ", val)
	}
}
