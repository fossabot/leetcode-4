/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 *
 * https://leetcode.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (40.97%)
 * Likes:    754
 * Dislikes: 971
 * Total Accepted:    184.5K
 * Total Submissions: 432.1K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * Given an integer array nums, find the sum of the elements between indices i
 * and j (i ≤ j), inclusive.
 *
 * Example:
 *
 * Given nums = [-2, 0, 3, -5, 2, -1]
 *
 * sumRange(0, 2) -> 1
 * sumRange(2, 5) -> -1
 * sumRange(0, 5) -> -3
 *
 *
 *
 * Note:
 *
 * You may assume that the array does not change.
 * There are many calls to sumRange function.
 *
 *
 */

package leetcode

// @lc code=start
type NumArray struct {
	numSum []int
}

func NumArrayConstructor(nums []int) NumArray {
	numsLen := len(nums)
	numSum := make([]int, numsLen+1)
	for i := 0; i < numsLen; i++ {
		numSum[i+1] = (nums[i] + numSum[i])
	}
	return NumArray{numSum: numSum}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.numSum[j+1] - this.numSum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end
