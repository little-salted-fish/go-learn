package main

import (
	"fmt"
)

func main() {

	//nums := [][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}, {0, 1}}
	nums := [][]int{{1, 2}, {3, 4}}
	fmt.Println(merge(nums))
	//nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	//fmt.Println("结果", merge(nums))

}

func merge(nums [][]int) [][]int {
	//排序
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j][0] < nums[i][0] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	merged := [][]int{}
	for i := 0; i < len(nums); i++ {
		l := nums[i][0]
		r := nums[i][1]

		if len(merged) == 0 || merged[len(merged)-1][1] < l {
			merged = append(merged, nums[i])

		} else {

			max := 0
			if merged[len(merged)-1][1] > r {
				max = merged[len(merged)-1][1]
			} else {
				max = r
			}
			merged[len(merged)-1][1] = max
		}

	}
	return merged
}

// 第一版有错误，leetcode没有通过
func merge1(nums [][]int) [][]int {

	//排序
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j][0] < nums[i][0] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp

			}
		}
	}
	res := [][]int{}
	res = append(res, nums[0])

	for i := 1; i < len(nums); i++ {
		resLen := len(res)

		a0 := res[resLen-1][1]
		a1 := res[resLen-1][0]
		b0 := nums[i][0]

		b1 := nums[i][1]

		if (a0 >= b0 && a0 <= b1 || (a1 < b1)) && b0-a0 <= 1 {
			res[resLen-1][1] = nums[i][1]

		} else if a1 > b1 && b0-a0 <= 1 {
			res[resLen-1][1] = nums[i][0]
		} else {
			res = append(res, nums[i])
		}
	}

	return res

}
