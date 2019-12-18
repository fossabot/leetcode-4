/*
 * @lc app=leetcode id=1 lang=cpp
 *
 * [1] Two Sum
 */

// @lc code=start
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
public:
  vector<int> twoSum(vector<int> &nums, int target) {
    unordered_map<int, int> numMap;

    for (int i = 0; i < nums.size(); i++) {
      int difference = target - nums[i];
      if (numMap.find(difference) != numMap.end()) {
        return vector<int>{numMap[difference], i};
      } else {
        numMap[nums[i]] = i;
      }
    }

    return {};
  }
};
// @lc code=end
