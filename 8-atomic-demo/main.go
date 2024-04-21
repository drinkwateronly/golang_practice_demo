package main

import "sync/atomic"

const x int64 = 1<<32 + 1

func main() {
	var i int64
	atomic.StoreInt64(&i, x)
	_ = i
}
