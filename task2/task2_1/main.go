package main

import "fmt"

func main() {

	a := 10
	fmt.Println("外部a=", a)
	changeValue(&a)
	fmt.Println("指针参数函数增加后a=", a)
	changeValue1(a)
	fmt.Println("值参数函数增加后a=", a)

	nums := []int{1, 2, 3}
	fmt.Println("原始切片", nums)
	sliceMultiply(&nums)
	fmt.Println("调用函数后切片", nums)

	sliceMultiply1(nums)
	fmt.Println("调用普通函数后切片", nums)

}

func changeValue(a *int) {
	*a = *a + 10
}

func changeValue1(a int) {
	a = +10
}

// 指针传递
func sliceMultiply(nums *[]int) {
	for i, v := range *nums {
		(*nums)[i] = v * 2
	}

}

//切片在Go中是一个引用类型，它包含三个部分：

//指向底层数组的指针

//长度

// 容量
// 当你传递切片给函数时，虽然传递的是切片的副本（值传递），但这个副本仍然指向同一个底层数组。因此，在函数内部对切片元素的修改会反映到原始切片中
// 值传递，但底层的原始数据是引用
func sliceMultiply1(nums []int) {
	for i, v := range nums {
		(nums)[i] = v * 3
	}

}
