package main

import (
	"fmt"
)

func main() {

	digits := []int{1, 2, 9}
	fmt.Println(plusOne(digits))

}

// 比较9就行不用相加 优化
func plusOne(digits []int) []int {

	plus := 1

	for i := len(digits) - 1; i >= 0; i-- {

		v := digits[i]

		if v+plus == 10 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + plus
			return digits
		}
		if i == 0 && plus == 1 {
			res := make([]int, len(digits)+1)
			res[0] = 1
			return res
		}

	}

	return digits
}
