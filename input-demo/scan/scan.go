package main

import (
	"bufio"
	"fmt"
	"os"
)

func scan() {
	var a int
	var b string
	var c byte
	fmt.Scan(&a, &b, &c)
	fmt.Printf("%d %T \n", a, a)
	fmt.Printf("%s %T \n", b, b)
	fmt.Printf("%c %T \n", c, c)
}

func scanln() {
	var a int
	var b string
	var c byte
	fmt.Scanln(&a, &b, &c)
	fmt.Printf("%d %T \n", a, a)
	fmt.Printf("%s %T \n", b, b)
	fmt.Printf("%c %T \n", c, c)
}

func scanf() {
	var a int
	var b string
	var c byte
	fmt.Scanf("%d %s %d", &a, &b, &c)
	fmt.Printf("%d %T \n", a, a)
	fmt.Printf("%s %T \n", b, b)
	fmt.Printf("%c %T \n", c, c)
}

func buffio() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("请输入一行文本: ")
	scanner.Scan()
	input := scanner.Text()

	writer := bufio.NewWriter(os.Stdout)
	writer.Write([]byte(fmt.Sprintf("您输入的文本是：%s\n", input)))
	writer.Flush()
}

func main() {
	buffio()
}
