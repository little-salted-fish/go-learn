package main

import (
	"fmt"
)

func main() {

	//str := []string{"hello", "hel world", "hel tom"}

	str := []string{"flower", "flow", "flight"}
	fmt.Println("返回", longestCommonPrefix(str))

}

// 列循环对比
func longestCommonPrefix(strs []string) string {

	var minLen = -1
	for _, v := range strs {
		if minLen == -1 {
			minLen = len(v)
		} else if minLen > len(v) {
			minLen = len(v)
		}
	}

	prefix := ""
out:
	for i := 0; i < minLen; i++ {

		a := ""
		for _, v := range strs {
			//fmt.Println(i, v[i:i+1])
			if a == "" {
				a = v[i : i+1]
			} else if a == v[i:i+1] {
				//prefix += a
			} else {
				break out
			}
		}
		prefix += a

	}
	return prefix
}
