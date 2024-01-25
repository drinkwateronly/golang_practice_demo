package main

import (
	"fmt"
	"sync"
	"time"
)

var testMap = make(map[int]int, 10)
var lock sync.Mutex

func testNum(num int) {
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	lock.Lock()
	testMap[num] = res
	lock.Unlock()
}
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go testNum(i)
	}
	time.Sleep(10 * time.Second)
	lock.Lock()
	for key, val := range testMap {
		fmt.Println("key: ", key, "val: ", val)
	}
	lock.Unlock()
	end := time.Since(start)
	fmt.Println(end)
}
