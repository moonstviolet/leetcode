/*
 * Created by Reverie
 * 2020.07.28
 * 23:20:43
 */

package main

import "fmt"

//扫一遍,记录以已经出现的最小值,用当前值去减,维护答案
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n, ans, minVal := len(prices), 0, prices[0]
	for i := 1; i < n; i++ {
		if prices[i]-minVal > ans {
			ans = prices[i] - minVal
		}
		if prices[i] < minVal {
			minVal = prices[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
