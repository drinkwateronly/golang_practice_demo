package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseIntDemo() {
	// 9进制的11为10进制的10，解析成功
	nInt, err := strconv.ParseInt("11", 9, 8)
	fmt.Println(nInt, err)
	// 10 <nil>

	// base指定了字符串为16进制，此时字符串不应该带前缀，解析失败
	nInt, err = strconv.ParseInt("0x01", 16, 8)
	fmt.Println(nInt, err)
	// 0 strconv.ParseInt: parsing "0x01": invalid syntax

	// base未指定进制，字符串前缀符合16进制，解析成功
	nInt, err = strconv.ParseInt("0x0f", 0, 8)
	fmt.Println(nInt, err)
	// 15 <nil>

	// -1024超出了int8的范围，解析失败，此时nInt为int8类型最小值，即-128。
	nInt, err = strconv.ParseInt("-1024", 10, 8)
	fmt.Println(nInt, err)
	// 输出：-128 strconv.ParseInt: parsing "-1024": value out of range

	// -111未超出int8的范围，解析成功
	nInt, err = strconv.ParseInt("111", 10, 8)
	fmt.Println(nInt, err)
	// 输出：-111 <nil>
}

func parseFloatDemo() {
	// 32位浮点数最大值位3.4*10^38，40个0溢出，因此解析失败
	fmt.Println(strings.Count("30000000000000000000000000000000000000000", "0"))
	nFloat, err := strconv.ParseFloat("30000000000000000000000000000000000000000", 32)
	fmt.Println(nFloat, err)
	// +Inf strconv.ParseFloat: parsing "30000000000000000000000000000000000000000": value out of range

	// 超出32位浮点最小精度，解析成功，但不准确
	nFloat, err = strconv.ParseFloat("0.123456789", 32)
	fmt.Println(nFloat, err)
	// 0.12345679104328156 <nil>
}

func parseBoolDemo() {
	b, err := strconv.ParseBool("f")
	fmt.Println(b, err)

	b, err = strconv.ParseBool("ff")
	fmt.Println(b, err)
}

func formatIntDemo() {
	fmt.Println(strconv.FormatInt(1023, 2))  // 1111111111
	fmt.Println(strconv.FormatInt(1023, 10)) // 1023
	fmt.Println(strconv.FormatInt(1023, 16)) // 3ff
}

func formatFloatDemo() {
	fmt.Println(strconv.FormatFloat(0.123456789, 'b', 32, 32))
	fmt.Println(strconv.FormatFloat(0.123456789, 'e', 32, 32))
	fmt.Println(strconv.FormatFloat(0.123456789, 'E', 32, 32))
	fmt.Println(strconv.FormatFloat(0.123456789, 'f', 32, 32))
	fmt.Println(strconv.FormatFloat(0.123456789, 'g', 32, 32))
	fmt.Println(strconv.FormatFloat(0.123456789, 'G', 32, 32))
	//16570090p-27
	//1.23456791043281555175781250000000e-01
	//1.23456791043281555175781250000000E-01
	//0.12345679104328155517578125000000
	//0.12345679104328155517578125
	//0.12345679104328155517578125
}

func main() {
	formatFloatDemo()
}
