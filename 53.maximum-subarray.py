#
# @lc app=leetcode id=53 lang=python3
#
# [53] Maximum Subarray
#
# https://leetcode.com/problems/maximum-subarray/description/
#
# algorithms
# Easy (45.22%)
# Likes:    5842
# Dislikes: 244
# Total Accepted:    722.5K
# Total Submissions: 1.6M
# Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
#
# Given an integer array nums, find the contiguous subarray (containing at
# least one number) which has the largest sum and return its sum.
#
# Example:
#
#
# Input: [-2,1,-3,4,-1,2,1,-5,4],
# Output: 6
# Explanation: [4,-1,2,1] has the largest sum = 6.
#
#
# Follow up:
#
# If you have figured out the O(n) solution, try coding another solution using
# the divide and conquer approach, which is more subtle.
#
#

# @lc code=start


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        numsLen = len(nums)
        if numsLen == 1:
            return nums[0]
        elif max(nums) <= 0:
            return max(nums)
        else:
            leftMax = self.maxSubArray(nums[:numsLen//2])
            rightMax = self.maxSubArray(nums[numsLen//2:])
            if leftMax < 0:
                return rightMax
            elif rightMax < 0:
                return leftMax
            else:
                return leftMax+rightMax
# @lc code=end
