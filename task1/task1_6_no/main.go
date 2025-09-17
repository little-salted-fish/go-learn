package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{0, 1, 1, 3, 1, 32, 3}))

}

func removeDuplicates(nums []int) int {

	cur := nums[0]

	for i := 1; i < len(nums); i++ {

		fmt.Println("当前值", cur, "遍历值", nums[i])
		if cur != nums[i] {
			cur = nums[i]
			fmt.Println("不等", cur, nums[i])
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			fmt.Println("相等", nums)
		}

	}
	fmt.Println("剩余数组", nums)
	return len(nums)

}

// 参考官网
func removeDuplicates1(nums []int) int {

	res := 1

	for i := 1; i < len(nums); i++ {

		if nums[i] != nums[i-1] {
			nums[res] = nums[i]
			res++
		}

	}
	fmt.Println(nums)
	return res

}
