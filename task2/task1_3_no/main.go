package main

import (
	"fmt"
)

func main() {

	fmt.Println(isValid("()}"))

}

func isValid(s string) bool {

	n := len(s)
	//奇数直接返回false
	if n%2 == 1 {
		return false
	}
	pairMap := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}

	for i := 0; i < n; i++ {

	}

	stack := []byte{}

	for _, b1 := range s {
		b := byte(b1)
		match := pairMap[b]

		if match > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != match {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, b)
		}
	}

	return len(stack) == 0
}
