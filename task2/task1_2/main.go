package main

import (
	"fmt"
	"time"
)

func main() {
	task1()
	task2()

}

func task2() {

	for i := 0; i < 10; i++ {
		go job(i)
	}

	time.Sleep(2 * time.Second)
}

func job(index int) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("任务%d,执行耗时(纳秒): %d ns\n", index, elapsed.Nanoseconds())
	}()
	time.Sleep(time.Duration(time.Nanosecond * time.Duration(index)))
}

func task1() {
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				fmt.Println("奇数", i)
			}

		}
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("偶数数", i)
			}

		}
	}()

	time.Sleep(3 * time.Second)
}
