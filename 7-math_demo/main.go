package main

import (
	"fmt"
	"math"
)

func demo1() {
	// 绝对值与取整
	x := -1234.567
	// 返回x绝对值
	fmt.Println(math.Abs(x)) // 1234.567
	// 返回不小于x的最小正整数
	fmt.Println(math.Ceil(x)) // -1234
	// 返回不大于x的最大正整数
	fmt.Println(math.Floor(x)) // -1235
	// 返回四舍五入最接近的整数
	fmt.Println(math.Round(x)) // -1235，对于负值而言，五入是向下入，即入了后数值更小；正数则是向上入，即入了后数值更大。
}

func demo2() {
	// 次方
	x := 2.
	y := 4.
	// Pow(x,y) 返回x的y方
	fmt.Println(math.Pow(x, y)) // 16
	// Sqrt(x) 返回x的平方根（开方），等价于math.Pow(x, 0.5)
	fmt.Println(math.Sqrt(x), math.Pow(x, 0.5))
	// Exp(x) 返回e的x次幂，等价于math.Pow(math.E, x)与math.SqrtE
	fmt.Println(math.Exp(x), math.Pow(math.E, x))
}

func demo3() {
	// 对数
	// 以自然数为底的对数
	fmt.Println(math.Log(math.E)) // 1
	// 以10为底的对数
	fmt.Println(math.Log(100) / math.Log(10)) // 2
}

func demo4() {
	// 三角函数
	fmt.Println(math.Pi)
	fmt.Println(math.Cos(math.Pi / 2))
	fmt.Println(math.Sin(math.Pi / 2))
	fmt.Println(math.Tan(math.Pi / 2))
}

func demo() {
	var count int
	fmt.Scanf("%d", &count)
	result := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&result[i])
	}
	fmt.Println(result)
}

func main() {
	demo()
}
