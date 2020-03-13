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
	}

	for _, test := range tests {
		if got := twoSum(test.input, 9); !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum(%v) = %v, want %v\n", test.input, got, test.want)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"tababad", "ababa"},
		{"abbapa", "abba"},
		{"cbbd", "bb"},
	}

	for _, test := range tests {
		if result := longestPalindrome(test.input); result != test.expected {
			t.Errorf("longestPalindrome(%v) = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"[", false},
	}

	for _, test := range tests {
		if result := isValid(test.input); result != test.expected {
			t.Errorf("isValid(%v) = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"+1", 1},
		{"+-1", 0},
		{"+", 0},
		{"9223372036854775808", 2147483647},
		{"18446744073709551617", 2147483647},
	}

	for _, test := range tests {
		if result := myAtoi(test.input); result != test.expected {
			t.Errorf("myAtoi(%v) = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "bb", -1},
		{"bb", "", 0},
		{"mississippi", "issip", 4},
	}

	for _, test := range tests {
		if result := strStr(test.haystack, test.needle); test.expected != result {
			t.Errorf("strStr(%v, %v) = %v, expected %v\n", test.haystack, test.needle, result, test.expected)
		}
	}
}

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
	}

	for _, test := range tests {
		if result := countAndSay(test.input); test.expected != result {
			t.Errorf("countAndSay(%v) = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		// {[]int{}, 0},
		{[]int{1, 1}, 1},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
	}

	for _, test := range tests {
		if result := maxArea(test.input); test.expected != result {
			t.Errorf("maxArea(%v) = %v, want %v\n", test.input, result, test.expected)
		}
	}
}
func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		input       []int
		inputTarget int
		expected    int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
	}

	for _, test := range tests {
		if result := threeSumClosest(test.input, test.inputTarget); test.expected != result {
			t.Errorf("threeSumCLosest(%v, %v) = %v, want %v\n", test.input, test.inputTarget, result, test.expected)
		}
	}
}

// func TestLetterCombinations(t *testing.T) {
func TestFindSubtring(t *testing.T) {
	tests := []struct {
		inString string
		inWords  []string
		expected []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"", []string{"word"}, []int{}},
		{"word", []string{}, []int{}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
	}

	for _, test := range tests {
		if result := findSubstring(test.inString, test.inWords); !reflect.DeepEqual(test.expected, result) {
			t.Errorf("findSubstring(%v, %v) = %v, expected %v\n", test.inString, test.inWords, result, test.expected)
		}
	}
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected string
	}{
		{"", "ABC", ""},
		{"ADOBECODEBANC", "", ""},
		{"ADOBECODEBANC", "AABC", "ADOBECODEBA"},
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"ADOBECODEBANCA", "AABC", "BANCA"},
		{"a", "aa", ""},
	}

	for _, test := range tests {
		if result := minWindow(test.s, test.t); test.expected != result {
			t.Errorf("minWindow(%v, %v) = %v, want %v\n", test.s, test.t, result, test.expected)
		}
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		input    [][]byte
		expected [][]byte
	}{
		{
			[][]byte{
				{'X', 'X'},
				{'X', 'O'},
			},
			[][]byte{
				{'X', 'X'},
				{'X', 'O'},
			},
		},
		{
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
	}

	for _, test := range tests {
		if solve(test.input); !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("solve(%v)", test.input)
		}
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

func TestIsMatch(t *testing.T) {
	tests := []struct {
		str      string
		pattern  string
		expected bool
	}{
		{"mississippi", "mis*is*p*.", false},
		// {"aaaaa", "bba", -1},
		// {"", "bb", -1},
		// {"bb", "", 0},
		// {"mississippi", "issip", 4},
	}

	for _, test := range tests {
		if result := isMatch(test.str, test.pattern); test.expected != result {
			t.Errorf("isMatch(%v, %v) = %v, expected %v\n", test.str, test.pattern, result, test.expected)
		}
	}
}

func TestPermute(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{[]int{}, [][]int{{}}},
		{[]int{1}, [][]int{{1}}},
		{[]int{1, 2, 3}, [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
	}

	for _, test := range tests {
		if got := permute(test.input); !reflect.DeepEqual(got, test.expected) {
			t.Errorf("permute(%v) = %v, want %v\n", test.input, got, test.expected)
		}
	}
}
