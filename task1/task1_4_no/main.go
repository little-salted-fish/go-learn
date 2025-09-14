package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{0}, 6))
	fmt.Println(twoSum([]int{1, 1}, 2))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 13))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 17))
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 6))

}

func twoSum(nums []int, target int) []int {

	if len(nums) < 2 {
		return nil
	}
	pairMap := make(map[int]int)

	for _, v := range nums {
		if pairMap[target-v] > 0 {
			return []int{v, pairMap[target-v]}

		} else {
			pairMap[v] = v
		}

	}

	return nil

}
