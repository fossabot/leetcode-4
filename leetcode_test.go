package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	str := "abc"
	val := lengthOfLongestSubstring(str)
	if val != 3 {
		for _, char := range str {
			fmt.Println(char)
		}
	}
}
