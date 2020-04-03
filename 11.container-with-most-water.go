/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (48.75%)
 * Likes:    4925
 * Dislikes: 519
 * Total Accepted:    555K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * Given n non-negative integers a1, a2, ..., an , where each represents a
 * point at coordinate (i, ai). n vertical lines are drawn such that the two
 * endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together
 * with x-axis forms a container, such that the container contains the most
 * water.
 *
 * Note: You may not slant the container and n is at least 2.
 *
 *
 *
 *
 *
 * The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In
 * this case, the max area of water (blue section) the container can contain is
 * 49.
 *
 *
 *
 * Example:
 *
 *
 * Input: [1,8,6,2,5,4,8,3,7]
 * Output: 49
 */

package leetcode

// @lc code=start

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0
	for left < right {
		minHeight := min(height[left], height[right])
		result = max(result, minHeight*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

// @lc code=end
