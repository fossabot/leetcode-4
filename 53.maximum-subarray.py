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
        # numSum = [[val for val in range(i+1)] for i in range(len(nums))]
        # maxSum = nums[0]
        # for index in range(len(nums)):
        #     for smallIndex in range(index):
        #         localSum = numSum[index-1][smallIndex]+nums[index]
        #         if localSum > maxSum:
        #             maxSum = localSum
        #         numSum[index][smallIndex] = localSum

        #     if nums[index] > maxSum:
        #         maxSum = nums[index]
        #     numSum[index][index] = nums[index]
        if max(nums) <= 0:
            return max(nums)

        currentSum = [0]*len(nums)
        maxSum = 0
        for i, num in enumerate(nums):
            currentSum[i] = max(0, currentSum[i-1]+num)
        return max(currentSum)
# @lc code=end
