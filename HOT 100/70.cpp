/*
 * Created by Reverie
 * 2020.07.28
 * 23:03:07
 */

#include <bits/stdc++.h>
using namespace std;

//简单dp,f(n)=f(n-1)+f(n-2),斐波那契数列
class Solution {
  public:
    int climbStairs(int n) {
        if (n == 1) {
            return 1;
        } else if (n == 2) {
            return 2;
        }
        int f1 = 1, f2 = 2;
        for (int i = 3; i <= n; i++) {
            int temp = f1 + f2;
            f1 = f2;
            f2 = temp;
        }
        return f2;
    }
};

int main() {
    Solution s;
    cout << s.climbStairs(3) << endl;
    return 0;
}