/*
 * Created by Reverie
 * 2020.07.28
 * 23:10:33
 */

#include <bits/stdc++.h>
using namespace std;

//扫一遍,记录以已经出现的最小值,用当前值去减,维护答案
class Solution {
  public:
    int maxProfit(vector<int> &prices) {
        if (prices.size() == 0) {
            return 0;
        }
        int ans = 0, minVal = prices[0];
        for (int i = 1; i < prices.size(); i++) {
            ans = max(ans, prices[i] - minVal);
            minVal = min(minVal, prices[i]);
        }
        return ans;
    }
};

int main() {
    Solution s;
    vector<int> v = {7, 1, 5, 3, 6, 4};
    cout << s.maxProfit(v) << endl;
    return 0;
}