package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var status int64 = 0

func main() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c)
	}
	time.Sleep(time.Second * 2)
	broadcast(c)
	time.Sleep(time.Second * 2)
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	atomic.StoreInt64(&status, 1)
	c.Broadcast() // 唤醒所有goroutine
	c.Signal()    // 唤醒其中1个goroutine
	c.L.Unlock()
}

func listen(c *sync.Cond) {
	c.L.Lock()
	fmt.Println("listening")
	for atomic.LoadInt64(&status) != 1 {
		c.Wait()
	}
	fmt.Println("done")
	c.L.Unlock()
}
