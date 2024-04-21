package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 20; i++ {
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("goroutine: ", i)
			defer wg.Done()
		}(i)
	}
	// 等待任意10个协程运行完毕
	wg.Wait()
	fmt.Println("all goroutines done")
}
