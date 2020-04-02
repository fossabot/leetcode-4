package leetcode

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	recursionReverse(&nums)
	fmt.Println(nums)
}

func recursionReverse(nums *[]int) {
	if len(*nums) <= 1 {
		return
	}

	tmp := (*nums)[0]
	*nums = (*nums)[1:]
	recursionReverse(nums)
	*nums = append(*nums, tmp)
}
