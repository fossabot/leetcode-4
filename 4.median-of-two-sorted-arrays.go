/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (27.93%)
 * Likes:    5907
 * Dislikes: 897
 * Total Accepted:    592.4K
 * Total Submissions: 2.1M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 *
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 *
 * You may assume nums1 and nums2 cannot be both empty.
 *
 * Example 1:
 *
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * The median is 2.0
 *
 *
 * Example 2:
 *
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * The median is (2 + 3)/2 = 2.5
 *
 *
 */

package leetcode

import "math"

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)

	// If the first array is bigger than second array
	// then return the median of reversed sequence
	if l1 > l2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	for start, end := 0, l1; ; {
		nums1Med := (start + end) / 2
		nums2Med := (l2+l1+1)/2 - nums1Med
		// Values are initialized to terminal values. Changed later if they
		// aren't terminal
		nums1Left, nums1Right, nums2Left, nums2Right := math.MinInt64,
			math.MaxInt64, math.MinInt64, math.MaxInt64

		// The variables are set to normal values when the conditions aren't
		// terminal
		if nums1Med != 0 {
			nums1Left = nums1[nums1Med-1]
		}
		if nums1Med != l1 {
			nums1Right = nums1[nums1Med]
		}
		if nums2Med != 0 {
			nums2Left = nums2[nums2Med-1]
		}
		if nums2Med != l2 {
			nums2Right = nums2[nums2Med]
		}

		if nums1Left > nums2Right {
			end = nums1Med - 1
		} else if nums2Left > nums1Right {
			start = nums1Med + 1
		} else {
			if (l1+l2)%2 == 1 {
				return math.Max(float64(nums1Left), float64(nums2Left))
			}
			return (math.Max(float64(nums1Left), float64(nums2Left)) +
				math.Min(float64(nums1Right), float64(nums2Right))) / 2
		}
	}
}

// @lc code=end
