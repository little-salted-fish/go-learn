package main

import (
	"fmt"
)

func main() {

	//str := []string{"hello", "hel world", "hel tom"}

	str := []string{"flower", "flow", "flight"}
	fmt.Println("返回", longestCommonPrefix(str))

}

func longestCommonPrefix(strs []string) string {
	//官方写法，以下为自己写
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]

}

// 列循环对比（优化可以去除一个计算最小长度的for循环）
func longestCommonPrefix1(strArr []string) string {

	var minLen = -1
	for _, v := range strArr {
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
		for _, v := range strArr {
			//fmt.Println(i, v[i:i+1])
			if a == "" {
				a = v[i : i+1]
			} else if a == v[i:i+1] {
				//prefix += a
				//这步走没用可以跳过
			} else {
				break out
			}
		}
		prefix += a

	}
	return prefix
}
