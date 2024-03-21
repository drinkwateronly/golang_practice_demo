package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// 复制Mutex定义的常量
const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	// 如果能成功抢到锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	// 如果处于唤醒、加锁或饥饿状态，此次请求不参与竞争，返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	// 尝试在竞争状态请求锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}

func try() {
	var mu Mutex
	go func() {
		mu.Lock()
		time.Sleep(time.Second * 1)
		mu.Unlock()
	}()
	time.Sleep(time.Second * 2)
	ok := mu.TryLock()
	// 成功获取锁
	if ok {
		fmt.Println("get lock !")
		mu.Unlock()
		return
	}
	fmt.Println("can't get lock !")
}

func main() {
	try()
}
