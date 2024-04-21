package main

import (
	"fmt"
	"sync/atomic"
)

var x int64 = 256

func add(addr *int64, delta int64) {
	// 原子加
	atomic.AddInt64(addr, delta)
	fmt.Println("add: x =", *addr)
}

func load(addr *int64) {
	// 原子读
	fmt.Println("load: x =", atomic.LoadInt64(addr))
}

func store(addr *int64, newValue int64) {
	// 原子写
	atomic.StoreInt64(addr, newValue)
	fmt.Println("store: x =", *addr)
}

func CAS(addr *int64, oldValue int64, newValue int64) {
	// 原子对比交换
	if atomic.CompareAndSwapInt64(addr, oldValue, newValue) {
		fmt.Printf("CAS: oldx=%d to newx=%d \n", oldValue, *addr)
	} else {
		fmt.Printf("CAS: oldx %d != oldValue %d, not swap \n", *addr, oldValue)
	}
}

func swap(addr *int64, newValue int64) {
	// 原子交换
	oldValue := atomic.SwapInt64(addr, newValue)
	fmt.Printf("swap: oldx=%d to newx=%d \n", oldValue, *addr)
}

func main() {
	add(&x, 1)        // 256+1
	load(&x)          // 257
	store(&x, 256)    // 256
	CAS(&x, 256, 128) // 128
	CAS(&x, 256, 64)  // 127
	swap(&x, 64)      // 256
}
