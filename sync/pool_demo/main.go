package main

import (
	"fmt"
	"sync"
)

type Student struct {
	id   int
	name string
}

func main() {
	studentPool := sync.Pool{
		//
		New: func() interface{} {
			return new(Student)
		},
	}
	for i := 0; i < 5; i++ {
		student := studentPool.Get().(*Student)
		fmt.Println("get from pool: ", student.id, student.name)
		student.id = i
		student.name = "a"
		fmt.Println("put to pool: ", student.id, student.name)
		studentPool.Put(student)
	}
}
