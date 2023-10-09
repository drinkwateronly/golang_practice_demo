package main

import (
	"fmt"
)

type Interfaces interface {
	getString() string
}

type structurer struct {
	Interfaces
	a int
}

func doPrint(i Interfaces) {
	fmt.Println(i.getString())
}

func main() {
	str := new(structurer)
	doPrint(str)
}
