package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//https://www.liwenzhou.com/posts/Go/json-tricks/#c-0-0-0

func demo1() {
	type Partner struct {
		Name string `json:"firstName"`
		Age  uint   `json:"age"`
	}

	type Person struct {
		FirstName string  `json:"firstName"`       // 指定字段名
		LastName  string  `json:"-"`               // "-"忽略某个字段
		Email     string  `json:"email,omitempty"` // omitempty忽略空值字段，需要和指定字段名搭配使用
		Age       uint    `json:"age"`
		Height    float64 `json:"height"`
		Sex       bool    `json:"sex"`

		PartnerA Partner  `json:"partnerA"`           // 嵌套结构体
		PartnerB *Partner `json:"partnerB,omitempty"` // 只有嵌套结构体是指针时，才能忽略嵌套结构体的空值字段
	}

	p := Person{
		FirstName: "John",
		LastName:  "Smith",
		Email:     "john@example.com",
		Age:       11,
		Height:    180.01,
		Sex:       true,
	}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	s := `{"firstName":"John", "lastName":"Smith","email":"john@example.com","age":11,"height":180.01,"sex":True,"partnerA":{"firstName":"","age":0}}`
	var p2 Person
	err = json.Unmarshal([]byte(s), &p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", p2)
}

func Demo2() {
	type User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	type UserPublic struct {
		*User
		Password string `json:"-"`
	}
}

func demo3() {
	//处理数字和字符类型的数字
	// 在该json中，num是数字，numString是字符串数字
	s := `{"num": 1, "numString": "12.34"}`

	type Nums struct {
		Num       int     `json:"num"`
		NumString float64 `json:"numString,string"`
	}
	var nums Nums

	err := json.Unmarshal([]byte(s), &nums)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("%+v\n", nums)

}

func parseTimeInDefaultDemo() {
	type Times struct {
		CreatedAt time.Time `json:"createdAt"`
	}
	// 反序列化
	str := `{"createdAt": "2024-01-25 01:23:45"}`
	var times Times
	err := json.Unmarshal([]byte(str), &times)
	if err != nil {
		// parsing time "2024-01-25 01:23:45" as "2006-01-02T15:04:05Z07:00": cannot parse " 01:23:45" as "T"
		fmt.Println("unmarshal error:", err)
	}

	// 序列化
	b, err := json.Marshal(Times{CreatedAt: time.Now()})
	if err != nil {
		// 不会err
		return
	}
	fmt.Println(string(b))
}

type Order struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

const layout = "2006-01-02 15:04:05"

func (o *Order) MarshalJSON() ([]byte, error) {
	type TempOrder Order // 定义与Order字段一致的新类型，避免直接嵌套Order进入死循环
	return json.Marshal(struct {
		CreatedAt  string `json:"createdAt"`
		*TempOrder        // 避免直接嵌套Order进入死循环
	}{
		CreatedAt: o.CreatedAt.Format(layout),
		TempOrder: (*TempOrder)(o), // 类型转换，虽然TempOrder也有CreatedAt，但编译器会直接使用当前作用域已有的CreatedAt
	})
}

func (o *Order) UnmarshalJSON(data []byte) error {
	type TempOrder Order // 定义与Order字段一致的新类型，
	ot := struct {
		CreatedAt  string `json:"createdAt"`
		*TempOrder        // 避免直接嵌套Order进入死循环
	}{
		TempOrder: (*TempOrder)(o), // 类型转换，当前作用域没有的CreatedAt，编译器会直接使用TempOrder的CreatedAt
	}

	if err := json.Unmarshal(data, &ot); err != nil {
		return err
	}

	var err error
	o.CreatedAt, err = time.Parse(layout, ot.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func parseTimeInCustomDemo() {

}

func main() {
	//demo3()
	parseTimeInDefaultDemo()
}
