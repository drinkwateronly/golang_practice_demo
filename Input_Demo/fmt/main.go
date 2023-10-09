package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//var a int
	//var b string
	//fmt.Println("请输入a和b：")
	////&a 获取a的内存地址
	//fmt.Scan(&a, &b)
	//fmt.Println("a is", a, "b is", b)

	var node TreeNode = TreeNode{
		Val: 1,
	}
	fmt.Println(node)
}
