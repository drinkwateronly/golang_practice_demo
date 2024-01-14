package main

import (
	"fmt"
	"strings"
)

func splitDemo() {
	a := "  hello  world !  "
	fmt.Printf("Fields: %q\n", strings.Fields(a))
	fmt.Printf("Split: %q\n", strings.Split(a, " "))
	fmt.Println()

	a = "__hello__world_!___"
	fmt.Printf("FieldsFunc: %q\n", strings.FieldsFunc(a, func(r rune) bool {
		return r == '_'
	}))
	fmt.Printf("Split: %q\n", strings.Split(a, "_"))
	fmt.Printf("SplitAfter: %q\n", strings.SplitAfter(a, "_"))
	fmt.Printf("SplitN: %q\n", strings.SplitN(a, "_", 3))
	fmt.Printf("SplitAfterN: %q\n", strings.SplitAfterN(a, "_", 3))
}

func demo() {
	a := "hElLo _WoRlD !"
	fmt.Println(strings.ToLower(a))
	fmt.Println(strings.ToUpper(a))
	fmt.Println(strings.ToTitle(a))
}

func trimDemo() {
	a := "__hello__world_!___"
	// 基础
	fmt.Printf("Trim: %q\n", strings.Trim(a, "_"))
	fmt.Printf("TrimLeft: %q\n", strings.TrimLeft(a, "_"))
	fmt.Printf("TrimRight: %q\n", strings.TrimRight(a, "_"))
	// Func
	delFunc := func(r rune) bool {
		return r == '_'
	}
	fmt.Printf("TrimFunc: %q\n", strings.TrimFunc(a, delFunc))
	fmt.Printf("TrimLeftFunc: %q\n", strings.TrimLeftFunc(a, delFunc))
	fmt.Printf("TrimRightFunc: %q\n", strings.TrimRightFunc(a, delFunc))
	// 前后缀
	fmt.Printf("TrimPrefix: %q\n", strings.TrimPrefix(a, "__h"))
	fmt.Printf("TrimSuffix: %q\n", strings.TrimSuffix(a, "_!___"))
	// 特殊
	a = "  hello  world !   "
	fmt.Printf("TrimSpace: %q\n", strings.TrimSpace(a))
}

func compareDemo() {
	// a < b < c
	a := "hello worlc"
	b := "hello world"
	c := "hello world !"
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(b, b))
	fmt.Println(strings.Compare(c, b))

	e := "Hello World !"
	fmt.Println(strings.EqualFold(c, e))
}

func replaceDemo() {
	a := "hello world ! hello world ! hello world !"
	fmt.Println(strings.Replace(a, "hello", "hi", 1))
	fmt.Println(strings.Replace(a, "hello", "hi", 2))
	fmt.Println(strings.Replace(a, "hello", "hi", -100))
	fmt.Println(strings.ReplaceAll(a, "hello", "hi"))
}

func joinAndRepeatDemo() {
	a := []string{" ", "hello", "world", "!"}
	fmt.Printf("Join: %q\n", strings.Join(a, "_"))
	fmt.Printf("Repeat: %q\n", strings.Repeat("#", 10))
}

func containDemo() {
	a := "hello world 你好! hello world 你好!"
	fmt.Println(strings.Count(a, "hello"))
	fmt.Println(strings.Contains(a, "hello"))
	fmt.Println(strings.ContainsAny(a, "abcd"))
	fmt.Println(strings.ContainsRune(a, '!'))
	fmt.Println(strings.HasPrefix(a, "hello"))
	fmt.Println(strings.HasSuffix(a, "你好!"))
}

func indexDemo() {
	a := "hello world 你好! hello world 你好!"
	fmt.Println(strings.Index(a, "hello"))
	fmt.Println(strings.IndexAny(a, "你好"))
	fmt.Println(strings.IndexRune(a, '你'))
	fmt.Println(strings.IndexByte(a, 'w'))
	fmt.Println(strings.IndexFunc(a, func(r rune) bool {
		return r == 'w'
	}))

	fmt.Println(strings.LastIndex(a, "hello"))
	fmt.Println(strings.LastIndexAny(a, "你好"))
	// 没有LastIndexRune
	fmt.Println(strings.LastIndexByte(a, 'w'))
	fmt.Println(strings.LastIndexFunc(a, func(r rune) bool {
		return r == 'w'
	}))
}
func main() {
	indexDemo()
}
