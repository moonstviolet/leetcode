/*
 * Created by Reverie
 * 2020.07.28
 * 23:07:11
 */

package main

import "fmt"

//简单dp,f(n)=f(n-1)+f(n-2),斐波那契数列
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	f1, f2 := 1, 2
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}

func main() {
	fmt.Println(climbStairs(3))
}
