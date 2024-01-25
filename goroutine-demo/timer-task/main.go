package main

import (
	"fmt"
	"time"
)

func main() {
	// 方式1
	fmt.Println(time.Now().Format("15:04:05"))
	timer := time.NewTimer(time.Second * 2)
	select {
	case now := <-timer.C:
		fmt.Println(now.Format("15:04:05"))
	}
	// 方式2
	now := <-time.After(time.Second * 2)
	fmt.Println(now.Format("15:04:05"))
}
