package leetcode

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, [][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}}},
	}

	for _, test := range tests {
		if rotate(test.input); !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("Result = %v, want %v\n", test.input, test.expected)
		}
	}
}

func TestGetRow(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{3, []int{1, 3, 3, 2}},
		// {2, "11"},
		// {3, "21"},
		// {4, "1211"},
		// {5, "111221"},
	}

	for _, test := range tests {
		if result := getRow(test.input); !reflect.DeepEqual(test.expected, result) {
			t.Errorf("getRow(%v) = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		num1     string
		num2     string
		expected string
	}{
		{"0", "3", "0"},
		{"2", "3", "6"},
		{"123", "456", "56088"},
	}

	for _, test := range tests {
		if result := multiply(test.num1, test.num2); test.expected != result {
			t.Errorf("multiply(%v, %v) = %v, expected %v\n", test.num1, test.num2, result, test.expected)
		}
	}
}

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{")", 0},
		{"()", 2},
		{"(()", 2},
		{")()())", 4},
		{"()(()", 2},
		// {"(()(((()", 2},
	}

	for _, test := range tests {
		if result := longestValidParentheses(test.input); result != test.expected {
			t.Errorf("longestValidParentheses %v = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestLongestPalindromeLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abccccdd", 7},
		// {"abbapa", "abba"},
		// {"cbbd", "bb"},
	}

	for _, test := range tests {
		if result := longestPalindromeLength(test.input); result != test.expected {
			t.Errorf("longestPalindromeLength(%v) = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestSubsets(t *testing.T) {
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
		if result := subsets(test.input); !reflect.DeepEqual(result, test.expected) {
			t.Errorf("\nsubsets(%v) = %v, \n\t\texpected %v\n", test.input, result, test.expected)
		}
	}
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		input    [][]byte
		expected int
	}{
		{[][]byte{}, 0},
		{[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}, 1},
		{[][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}, 3},
	}

	for _, test := range tests {
		if result := numIslands(test.input); result != test.expected {
			t.Errorf("numIslands(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestLexicalOrder(t *testing.T) {
	tests := []struct {
		input    int
		expected []int
	}{
		{1, []int{1}},
		{9, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{10, []int{1, 10, 2, 3, 4, 5, 6, 7, 8, 9}},
		{13, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		if result := lexicalOrder(test.input); !reflect.DeepEqual(test.expected, result) {
			t.Errorf("lexicalOrder(%v) = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestExist(t *testing.T) {
	tests := []struct {
		inputBoard [][]byte
		inputWord  string
		expected   bool
	}{
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCCED", true},
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "SEE", true},
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCB", false},
	}

	for _, test := range tests {
		if result := exist(test.inputBoard, test.inputWord); result != test.expected {
			t.Errorf("exist(%v,%v) = %v, expected %v", test.inputBoard, test.inputWord, result, test.expected)
		}
	}
}

func TestMinCut(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"aab", 1},
		{"abbab", 1},
		{"noonabbad", 2},
		{"adabdcaebdcebdcacaaaadbbcadabcbeabaadcbcaaddebdbddcbdacdbbaedbdaaecabdceddccbdeeddccdaabbabbdedaaabcdadbdabeacbeadbaddcbaacdbabcccbaceedbcccedbeecbccaecadccbdbdccbcbaacccbddcccbaedbacdbcaccdcaadcbaebebcceabbdcdeaabdbabadeaaaaedbdbcebcbddebccacacddebecabccbbdcbecbaeedcdacdcbdbebbacddddaabaedabbaaabaddcdaadcccdeebcabacdadbaacdccbeceddeebbbdbaaaaabaeecccaebdeabddacbedededebdebabdbcbdcbadbeeceecdcdbbdcbdbeeebcdcabdeeacabdeaedebbcaacdadaecbccbededceceabdcabdeabbcdecdedadcaebaababeedcaacdbdacbccdbcece", 273},
	}

	for _, test := range tests {
		if result := minCut(test.input); result != test.expected {
			t.Errorf("minCut(%v) = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
	}

	for _, test := range tests {
		if result := groupAnagrams(test.input); !reflect.DeepEqual(result, test.expected) {
			t.Errorf("groupAnagrams(%v) = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestReversePairs(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 3, 2, 3, 1}, 2},
		{[]int{2, 4, 3, 5, 1}, 3},
		{[]int{2147483647, 2147483647, -2147483647, -2147483647, -2147483647,
			2147483647}, 9},
	}

	for _, test := range tests {
		if result := reversePairs(test.input); test.expected != result {
			t.Errorf("reversePairs(%v) = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestIsHappy(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{1, true},
		{19, true},
		{20, false},
	}

	for _, test := range tests {
		if result := isHappy(test.input); test.expected != result {
			t.Errorf("isHappy(%v) = %v, expected %v\n", test.input, result, test.expected)
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte("hello"), []byte("olleh")},
	}

	for _, test := range tests {
		if reverseString(test.input); !reflect.DeepEqual(test.expected, test.input) {
			t.Errorf("reverseString(%v), expected %v\n", string(test.input), string(test.expected))
		}
	}
}

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		time     [][]int
		nodeSize int
		source   int
		expected int
	}{
		{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
	}

	for _, test := range tests {
		if result := networkDelayTime(test.time, test.nodeSize, test.source); result != test.expected {
			t.Errorf("networkDelayTime result %v, expected %v\n", result, test.expected)
		}
	}
}
