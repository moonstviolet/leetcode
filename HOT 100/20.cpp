/*
 * Created by Reverie
 * 2020.07.28
 * 21:55:13
 */

#include <bits/stdc++.h>
using namespace std;

//遇到左括号入栈,遇到右括号看栈顶是否匹配,最后看栈是否为空
class Solution {
  public:
    bool isValid(string s) {
        stack<char> st;
        for (int i = 0; i < s.size(); i++) {
            switch (s[i]) {
            case '(':
            case '[':
            case '{':
                st.push(s[i]);
                break;
            case ')':
                if (st.empty() || st.top() != '(') {
                    return false;
                }
                st.pop();
                break;
            case ']':
                if (st.empty() || st.top() != '[') {
                    return false;
                }
                st.pop();
                break;
            case '}':
                if (st.empty() || st.top() != '{') {
                    return false;
                }
                st.pop();
                break;
            default:
                return false;
            }
        }
        return st.empty();
    }
};

int main() {
    Solution s;
    cout << s.isValid("{[]}") << endl;
    return 0;
}