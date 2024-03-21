package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func main() {
	// 初始化没有人占用的锁
	l := RecursiveMutex{owner: -1}
	l.Lock()
	fmt.Println("第一次锁")
	l.Lock()
	fmt.Println("可重入，第二次锁")
	l.Unlock()
	l.Unlock()
}

// GoID 解析goroutine_id
func GoID() int64 {
	var buf [64]byte
	// 通过runtime.Stack获取栈帧信息，从该信息获取goroutine_id
	n := runtime.Stack(buf[:], false)
	//
	//fmt.Println(string(buf[:n]))
	idFiled := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine"))[0]
	id, err := strconv.ParseInt(idFiled, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

// 根据goroutine id可以获得一个可重入锁
type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 当前持有锁的goroutine id
	recursion int32 // 重入次数
}

func (m *RecursiveMutex) Lock() {
	gid := GoID()
	// 如果当前锁的持有者是本次调用Lock的goroutine
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	// 此时两种情况，未上锁时，上锁
	// 已上锁时，阻塞
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := GoID()
	// 当前锁的持有者不是调用Unlock的goroutine，错误
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong lock owner(%d): %d", m.owner, gid))
	}
	// 重入次数自减
	m.recursion--
	if m.recursion != 0 {
		return
	}
	// 重入次数为0了，释放该锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}
