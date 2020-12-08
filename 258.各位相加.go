/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	if num%9 == 0 && num != 0 {
		return 9
	}
	return num % 9
}

// @lc code=end

