package main

//  go test -bench=Benchmark -benchmem .\string_concat_test.go
import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var s1 []string = []string{
	"Chen Jie",
	"Jie Chen",
	"Jhen Cie",
}

func concatByOperator(s []string) string {
	var result string
	for _, v := range s {
		result += v
	}
	return result
}

func concatByJoin(s []string) string {
	return strings.Join(s, "")
}

func concatByFmt(s []string) string {
	var result string
	for _, v := range s {
		result = fmt.Sprintf("%s%s", result, v)
	}
	return result
}

func concatByStringsBuilder(s []string) string {
	var result strings.Builder
	for _, v := range s {
		result.WriteString(v)
	}
	return result.String()
}

func concatByStringsBuilderWithInitSize(s []string) string {
	var result strings.Builder
	result.Grow(64) // 初始化
	for _, v := range s {
		result.WriteString(v)
	}
	return result.String()
}

func concatByBytesBuffer(s []string) string {
	var result bytes.Buffer
	for _, v := range s {
		result.WriteString(v)
	}
	return result.String()
}

func concatByBytesBufferWithInitSize(s []string) string {
	var result bytes.Buffer
	result.Grow(64) // 初始化
	for _, v := range s {
		result.WriteString(v)
	}
	return result.String()
}

func BenchmarkConcatStringByOperator(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatByOperator(s1)
	}
}
func BenchmarkConcatByJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatByJoin(s1)
	}
}
func BenchmarkConcatByFmt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatByFmt(s1)
	}
}
func BenchmarkConcatByStringsBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatByStringsBuilder(s1)
	}
}
func BenchmarkConcatByStringsBuilderWithInitSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatByStringsBuilderWithInitSize(s1)
	}
}
func BenchmarkConcatByBytesBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatByBytesBuffer(s1)
	}
}

func BenchmarkConcatByBytesBufferWithInitSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatByBytesBufferWithInitSize(s1)
	}
}
