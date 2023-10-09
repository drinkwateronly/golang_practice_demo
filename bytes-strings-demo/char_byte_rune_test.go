package main

import (
	"fmt"
	"strings"
	"testing"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func TestChar(t *testing.T) {
	fmt.Printf("%T \n", '1')
	fmt.Printf("%T \n", rune('1'))
	fmt.Printf("%T \n", byte('1'))
	fmt.Printf("%T \n", '陈')
	fmt.Printf("%T \n", rune('陈'))
	//fmt.Printf("%T", byte('陈'))  // 出错
	str := "123::456"
	fmt.Println(len(strings.Split(str, ":")))
	a := 1 << 32
	fmt.Println(-a)
	maxLeft := -(1 << 31)
	minRight := 1 << 31
	fmt.Println(max(maxLeft, minRight))
}
