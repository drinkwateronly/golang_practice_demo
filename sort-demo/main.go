package main

import (
	"fmt"
	"sort"
)

type myStruct struct {
	a int
}

type myStructs []myStruct

func (s myStructs) Len() int { return len(s) }

func (s myStructs) Less(i, j int) bool { return s[i].a > s[j].a }

func (s myStructs) Swap(i, j int) { s[i].a, s[j].a = s[j].a, s[i].a }

type mySlice []int

func (s mySlice) Len() int { return len(s) }

func (s mySlice) Less(i, j int) bool { return s[i] > s[j] }

func (s mySlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

//func searchDemo() {
//	x := 23
//	i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
//	if i < len(data) && data[i] == x {
//		// x is present at data[i]
//	} else {
//		// x is not present in data,
//		// but i is the index where it would be inserted.
//	}
//}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s\n", &s)
		fmt.Println(s != "" && s[0] == 'y')
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

func typeConvertDemo() {
	s := mySlice{1, 9, 3, 7, 4, 5}
	//s.Sort()              // 不可以
	//search := s.Search(3) // 不可以

	// 强制类型转换
	s2 := sort.IntSlice(s)
	s2.Sort()              // 可以
	search := s2.Search(6) // 可以

	fmt.Println(s2, search)
	// [1 3 4 5 7 9] 4
}

func reverseDemo() {
	s := sort.IntSlice{1, 9, 3, 7, 4, 5}
	sort.Ints(s) // 递增排序
	fmt.Println(s)

	r := sort.Reverse(s) // 对r排序是递减
	sort.Sort(r)
	fmt.Println(r)

	s1 := mySlice{1, 9, 3, 7, 4, 5}
	sort.Sort(s1)
	fmt.Println(s1)
	r1 := sort.Reverse(s1) // 对r排序是递减
	sort.Sort(r1)
	fmt.Println(r1)

}

func main() {
	//s1 := myStructs{myStruct{3}, myStruct{1}, myStruct{2}, myStruct{1}, myStruct{-4}, myStruct{-8}}
	//sort.Sort(s1)
	//fmt.Println(s1)

	GuessingGame()
}
