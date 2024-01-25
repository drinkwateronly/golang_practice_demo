package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	intChan <- -1
	fmt.Println(<-intChan)
}
