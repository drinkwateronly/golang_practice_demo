package main

import (
	"fmt"
	"sync"
	"time"
)

// 封装计数器
type Counter struct {
	sync.Mutex
	// 或者
	// sync.RWMutex
	Count int
}

func (c *Counter) Incr() {
	c.Lock()
	c.Count++
	c.Unlock()
}

func (c *Counter) GetCount() int {
	c.Lock()
	defer c.Unlock() // 先return后释放锁
	return c.Count
}

func main() {
	counter := Counter{}
	count := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter.Incr() // 计数器自增
		}()
		go func() {
			count++ // 变量计数器自增
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(counter.GetCount(), count) // 5000 4972
}
