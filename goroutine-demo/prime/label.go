package main

import (
	"time"
)

// 不断产生数字
func producer(intChan chan<- int, maxNum int) {
	for i := 1; i <= maxNum; i++ {
		intChan <- i
	}
	close(intChan)
}

// 循环取数，直到无法
func isPrime(intChan <-chan int, exitChan chan<- bool) {
	for {
		num, ok := <-intChan
		if !ok { // 管道关闭
			break
		}
		isPrime := true
		for i := 2; i < num; i++ {
			if num%(i) == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			//fmt.Println(num)
		}
	}
	exitChan <- true
}

func concurrentPrime(maxNum int) {
	intChan := make(chan int)
	exitChan := make(chan bool)
	go producer(intChan, maxNum)
	routineNum := 10
	for i := 0; i < routineNum; i++ {
		go isPrime(intChan, exitChan)
	}
	for i := 0; i < routineNum; i++ {
		<-exitChan
	}
}

func singleCurPrime(maxNum int) {
	for num := 1; num <= maxNum; num++ {
		isPrime := true
		for i := 2; i < num; i++ {
			if num%(i) == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			//fmt.Println(i)
		}
	}
}

func main() {
	begin := time.Now()
	singleCurPrime(2000)
	println("花费时间", int(time.Since(begin)))
	begin = time.Now()
	concurrentPrime(2000)
	println("花费时间", int(time.Since(begin)))
}
