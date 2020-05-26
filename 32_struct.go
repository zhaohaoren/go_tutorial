// 结构体
package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体，结构体是值类型的，引用直接指向的是内存空间，不像map和slice先指向一个地址，然后这个地址再指向内存空间
type Cat struct {
	Name  string //属性: 为值类型的默认值，如果是map slice 指针，则为nil即还未分配空间
	Age   int
	Color string
	pro   string // 不区分大小写-没有私有属性公有属性的概念
} // 命名规范：struct名字首字母大写

type Person struct {
	Scores [6]float64
	ptr    *int
	slice  []int
	map1   map[string]string // 如果需要使用这些引用类型属性，我们需要make才能使用
}

type A struct {
	// 结构体属性可以带上tag，在序列化和反序列化的时候将使用tag里面的名字进行序列化。
	Name string `json:"name"`
}
type B struct {
	Name string
}

func main() {
	// go只是通过结构体支持面向对象的特性，但不是面向对象的语言。面向对象是一种编程风格思想，不是语言的问题。
	// go是面向接口编程，通过接口进行关联

	var cat1 Cat
	fmt.Println(cat1)
	cat1.Age = 1
	cat1.Name = "Tom"
	cat1.Color = "grey"
	cat1.pro = "no"
	fmt.Println(cat1)

	var person1 Person
	fmt.Println(person1) //输出：{[0 0 0 0 0 0] <nil> [] map[]}，<nil> [] map[]本质上都是nil只是输出的形式不一样。
	person1.slice = make([]int, 10)
	person1.slice[0] = 1
	person1.slice[1] = 2
	person1.map1 = make(map[string]string)
	person1.map1["name"] = "justin"
	fmt.Println(person1)
	var p1 Person
	p1.ptr = new(int) // 指针用new的

	// 其他初始化方式
	// 1 -推荐
	var cat2 = Cat{"jack", 11, "red", "yes"}
	var catn = Cat{
		Name:  "nacy",
		Age:   20,
		Color: "pink",
		pro:   "no",
	}
	// 2
	var cat3 *Cat = new(Cat)
	(*cat3).Name = "jack2"
	cat3.Name = "jack3" // 同样起作用，底层go为对cat3简化加上了取地址
	// 3
	var cat4 *Cat = &Cat{"jack", 11, "red", "yes"}
	fmt.Println(cat2)
	fmt.Println(*cat3)
	fmt.Println(*cat4)
	fmt.Println(catn)

	// 内存：
	// 结构体所有字段在内存中是连续的，获取值可以通过地址的计算来进行偏移获取。

	// 注意事项：
	// 结构体和其他类型进行转换的时候需要有完全相同属性
	var a A
	var b B
	//a = b go中类型不同，不可以赋值
	a = A(b) //AB字段一样，才可以强制类型转换：名字类型个数都要完全一样
	fmt.Println(a, b)

	// 结构体tag -- 利用的是反射的原理-后续讲解
	var a1 = A{"a"}
	bytes, _ := json.Marshal(a1)
	fmt.Println(string(bytes)) // 将输出tag指定的字段名的序列化json串

}
