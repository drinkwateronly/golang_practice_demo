package main

func escape1() *int {
	var a int = 1
	return &a
}

func main() {
	escape1()
}
