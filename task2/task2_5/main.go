package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	lockAdd()
	noLockAdd()

}

func lockAdd() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	num := 0
	mutex := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go add1000(&num, &mutex, &wg)
	}
	wg.Wait()
	fmt.Println("加锁结果", num)
}
func add1000(num *int, sync *sync.Mutex, wg *sync.WaitGroup) {
	sync.Lock()
	defer sync.Unlock()
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		*num = *num + 1
	}
}

func noLockAdd() {
	wg1 := sync.WaitGroup{}
	wg1.Add(10)

	autoInt := atomic.Int64{}

	for i := 0; i < 10; i++ {
		go add1000NoLock(&autoInt, &wg1)
	}
	wg1.Wait()
	fmt.Println("无锁结果", autoInt.Load())
}

func add1000NoLock(num *atomic.Int64, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		num.Add(1)
	}
}
