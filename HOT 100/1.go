/*
 * Created by Reverie
 * 2020.07.22
 * 21:51:52
 */

package main

import (
	"fmt"
	"sort"
)

//暴力扫描,时间复杂度O(n^2),空间复杂度O(1)
func twoSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

type myInt struct {
	idx int
	val int
}
type intSlice []myInt

func (ints intSlice) Len() int {
	return len(ints)
}
func (ints intSlice) Less(i, j int) bool {
	return ints[i].val < ints[j].val
}
func (ints intSlice) Swap(i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

//排序+双指针,时间复杂度O(nlogn),空间复杂度O(n)
func twoSum2(nums []int, target int) []int {
	var ints intSlice
	for i, v := range nums {
		ints = append(ints, myInt{i, v})
	}
	sort.Sort(ints)
	i, j := 0, len(ints)-1
	for i < j {
		if ints[i].val+ints[j].val < target {
			i++
		} else if ints[i].val+ints[j].val > target {
			j--
		} else {
			return []int{ints[i].idx, ints[j].idx}
		}
	}
	return []int{}
}

//哈希表,时间复杂度O(n),空间复杂度O(n)
func twoSum3(nums []int, target int) []int {
	appear := make(map[int]int, len(nums))
	for i, v := range nums {
		if idx, ok := appear[target-v]; ok {
			return []int{idx, i}
		} else {
			appear[v] = i
		}
	}
	return []int{}
}

func main() {
	ans := twoSum3([]int{2, 7, 11, 15}, 9)
	fmt.Println(ans)
}
