package main

import (
	"fmt"
	"strconv"
)

func main() {

	b := isPalindrome(124213)
	fmt.Println(b)

	b = isPalindrome(121)
	fmt.Println(b)

	b = isPalindrome(1001)
	fmt.Println(b)

}

func isPalindrome(m int) bool {
	x := int64(m)

	str := strconv.FormatInt(x, 10)

	l := len(str)
	half := l / 2
	flag := false
	i := 0

	for {
		if i > half-1 {
			break
		}
		if str[i] == str[l-1-i] {
			flag = true
		} else {
			flag = false
		}
		i++
	}

	return flag
}
