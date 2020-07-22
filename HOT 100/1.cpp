/*
 * Created by Reverie
 * 2020.07.22
 * 21:04:50
 */

#include <bits/stdc++.h>
using namespace std;

//暴力扫描,时间复杂度O(n^2),空间复杂度O(1)
class Solution1 {
  public:
    vector<int> twoSum(vector<int> &nums, int target) {
        vector<int> ans;
        for (int i = 0; i < nums.size(); i++) {
            for (int j = i + 1; j < nums.size(); j++) {
                if (nums[i] + nums[j] == target) {
                    ans.push_back(i);
                    ans.push_back(j);
                    return ans;
                }
            }
        }
        return ans;
    }
};

//排序+双指针,时间复杂度O(nlogn),空间复杂度O(n)
class Solution2 {
  public:
    vector<int> twoSum(vector<int> &nums, int target) {
        vector<int> ans;
        vector<int> temp = nums;
        sort(nums.begin(), nums.end());
        int i = 0, j = (int)nums.size() - 1;
        while (i < j) {
            if (nums[i] + nums[j] < target) {
                i++;
            } else if (nums[i] + nums[j] > target) {
                j--;
            } else {
                break;
            }
        }
        if (i < j) {
            for (int k = 0; k < nums.size(); k++) {
                if (temp[k] == nums[i]) {
                    ans.push_back(k);
                } else if (temp[k] == nums[j]) {
                    ans.push_back(k);
                }
            }
        }
        return ans;
    }
};

//哈希表,时间复杂度O(n),空间复杂度O(n)
class Solution3 {
  public:
    vector<int> twoSum(vector<int> &nums, int target) {
        vector<int> ans;
        unordered_map<int, int> appear;
        for (int i = 0; i < nums.size(); i++) {
            if (appear[target - nums[i]] == 0) {
                appear[nums[i]] = i + 1;
            } else {
                ans.push_back(i);
                ans.push_back(appear[target - nums[i]] - 1);
            }
        }
        return ans;
    }
};

// test
int main() {
    Solution3 s;
    vector<int> v = {2, 7, 11, 15};
    auto ans = s.twoSum(v, 9);
    if (ans.size() == 2) {
        cout << ans[0] << " " << ans[1] << endl;
    }
    return 0;
}