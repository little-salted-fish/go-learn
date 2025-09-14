package main

import (
	"fmt"
)

func main() {

	b := isPalindrome(0)
	fmt.Println(b)

	b = isPalindrome(121)
	fmt.Println(b)

	b = isPalindrome(1001)
	fmt.Println(b)

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	//x := int64(m)
	//格式化有问题
	//str := strconv.FormatInt(x, 10)

	//
	str := fmt.Sprintf("%d", x)
	//直接用数字转字符
	//str := strconv.Itoa(m)

	l := len(str)
	half := l / 2
	flag := true
	i := 0

	for {
		if i > half-1 {
			break
		}
		if str[i] != str[l-1-i] {
			flag = false
		}
		i++
	}

	return flag
}
