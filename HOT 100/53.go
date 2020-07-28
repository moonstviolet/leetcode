/*
 * Created by Reverie
 * 2020.07.28
 * 22:59:18
 */

package main

import "fmt"

//滑动窗口,时间复杂度O(n)
func maxSubArray(nums []int) int {
	ans, sum := nums[0], 0
	for _, v := range nums {
		sum += v
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
