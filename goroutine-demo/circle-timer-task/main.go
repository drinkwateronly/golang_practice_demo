package main

import (
	"fmt"
	"sync"
	"time"
)

func circleTaskWithTicker() {
	count := 0
	ticker := time.NewTicker(time.Second)
	exitChan := make(chan bool)
	go func() {
		for {
			t := <-ticker.C
			fmt.Println(t)
			count++
			if count > 2 {
				ticker.Stop()
				exitChan <- true
			}
		}
	}()
	select {
	case <-exitChan:
		fmt.Println("end")
	}
}

func circleTaskWithTimer() {
	count := 0
	timer := time.NewTimer(time.Second)
	exitChan := make(chan bool)
	go func() {
		for {
			t := <-timer.C
			timer.Reset(time.Second)
			fmt.Println(t)
			count++
			if count > 2 {
				timer.Stop()
				exitChan <- true
			}
		}
	}()
	select {
	case <-exitChan:
		fmt.Println("end")
	}
}

func circleTaskWithSync() {
	count := 0
	ticker := time.NewTicker(time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1) // 告知WaitGroup只需要等待一个任务完成
	go func() {
		defer wg.Done()
		defer ticker.Stop()
		for {
			t := <-ticker.C
			fmt.Println(t)
			count++
			if count > 2 {
				return // 时间到，任务完成
			}
		}
	}()
	wg.Wait() // 等待
	fmt.Println("end")
}

func main() {
	circleTaskWithTimer()
}
