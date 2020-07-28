/*
 * Created by Reverie
 * 2020.07.28
 * 22:53:17
 */

#include <bits/stdc++.h>
using namespace std;

//滑动窗口,时间复杂度O(n)
class Solution {
  public:
    int maxSubArray(vector<int> &nums) {
        int ans = nums[0], sum = 0;
        for (int i = 0; i < nums.size(); i++) {
            sum += nums[i];
            if (sum > ans) {
                ans = sum;
            }
            if (sum < 0) {
                sum = 0;
            }
        }
        return ans;
    }
};

int main() {
    Solution s;
    vector<int> v = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
    cout << s.maxSubArray(v) << endl;
    return 0;
}