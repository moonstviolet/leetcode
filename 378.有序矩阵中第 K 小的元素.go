package main

import "fmt"

/*
 * 1 <= n <= 300
 * -10^9 <= matrix[i][j] <= 10^9
 * 1 <= k <= n^2
 */
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		m := l + (r-l)/2
		x, y, cnt := n-1, 0, 0
		for x >= 0 && x < n && y >= 0 && y < n {
			if matrix[x][y] > m {
				x--
			} else {
				cnt += x + 1
				y++
			}
		}
		// cnt表示小于等于m的数的数量
		if cnt == k {
			return m
		} else if cnt > k {
			r = m - 1
		} else {
			l = m
		}
	}

	return l
}

func main() {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	k := 8

	ans := kthSmallest(matrix, k)
	fmt.Println(ans)
	//13
}
