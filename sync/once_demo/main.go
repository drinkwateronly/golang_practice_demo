package main

import (
	"fmt"
	"sync"
	"time"
)

var a int
var initOnce sync.Once

func main() {
	for i := 0; i < 5; i++ {
		go incr()
	}
	time.Sleep(time.Second)
}

func initialize() {
	a = 2
	fmt.Println("init a")
}

func incr() {
	initOnce.Do(initialize)
	a++
	fmt.Println(a)
}
