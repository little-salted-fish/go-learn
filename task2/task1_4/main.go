package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channelConnect()
	channelConnect1()
	channelConnectCache()
}

func channelConnect() {

	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}

	}()
	go func() {

		for j := range ch {
			fmt.Println(j)
		}
	}()

	time.Sleep(10 * time.Second)
}

func channelConnectCache() {

	ch := make(chan int, 10)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
			if i%10 == 0 {
				time.Sleep(1 * time.Second)
			}
		}

	}()
	go func() {

		for j := range ch {
			fmt.Println(j)
		}
	}()

	time.Sleep(10 * time.Second)
}

// 等待携程都完成
func channelConnect1() {

	var wg = sync.WaitGroup{}
	wg.Add(2)

	ch := make(chan int)

	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Printf("发送: %d\n", i)
			time.Sleep(1 * time.Second)
		}

	}()
	go func() {
		defer func() {
			wg.Done()
		}()
		for j := range ch {
			fmt.Println("接收", j)
		}
	}()
	wg.Wait() // 等待所有goroutine完成
	fmt.Println("所有操作完成!")
}
