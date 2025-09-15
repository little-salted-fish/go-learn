package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//nums := getRandomIntSplice(5)
	//fmt.Println("随机int 切片", nums)
	nums := []int{1, 1, 2, 3, 3, 2, 5, 6, 5, 6, 7}
	result := singleNumber(nums)
	fmt.Println("结果是", result)

}

func singleNumber(nums []int) int {

	recordMap := make(map[int]int, len(nums))
	for _, v := range nums {
		if recordMap[v] == 0 {
			recordMap[v] = 1
		} else {
			recordMap[v] = recordMap[v] + 1
		}
	}

	for k, v := range recordMap {
		if v == 1 {
			return k
		}
	}
	panic("没有合适数据")
}

func getRandomIntSplice(length int) []int {

	if length <= 1 {
		panic("长度必须大于2")
	}
	arr := getRandomArr(length)
	nums := arr[:]
	randomInt := generateRandomInt(length-1, 0, len(arr)-1)
	nums = append(nums, arr[randomInt])
	return nums
}

// 获取随机int数组
func getRandomArr(length int) []int {
	// 创建新的随机数生成器（推荐方式）
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成随机数组
	randomArray := generateRandomIntArray(r, length-1, 1, 100)

	return randomArray
}
func generateRandomIntArray(r *rand.Rand, length, min, max int) []int {
	if length > (max - min + 1) {
		panic("范围太小，无法生成不重复的随机数")
	}

	uniqueMap := make(map[int]bool)
	result := make([]int, 0, length)

	for len(result) < length {
		num := r.Intn(max-min+1) + min
		if !uniqueMap[num] {
			uniqueMap[num] = true
			result = append(result, num)
		}
	}

	return result
}

func generateRandomInt(length, min, max int) int {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	a := r.Intn(max-min+1) + min
	return a
}
