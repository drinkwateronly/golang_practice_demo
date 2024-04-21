package main

import (
	"fmt"
	"sync"
)

func main() {
	var hash sync.Map
	hash.Store("A", 1)
	hash.Store(123, "b")
	hash.Store(1, "C")
	i := 0
	hash.Range(func(key, value any) bool {
		fmt.Println(key, value)
		i++
		if i == 2 {
			return false
		}
		return true
	})
}
