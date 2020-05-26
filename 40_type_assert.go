// 类型断言 - 不是java的那个断言，相当于强制类型转换
package main

import "fmt"

type Point struct {
	x int
	y int
}

func judgeType(items ...interface{}) {
	for idx, val := range items {
		switch val.(type) { // 固定写法, 用于判断类型
		case int:
			fmt.Printf("index-%v is int,val is %v \n", idx, val)
		case float32:
			fmt.Printf("index-%v is float32,val is %v \n", idx, val)
		case string:
			fmt.Printf("index-%v is string,val is %v \n", idx, val)
		}
	}
}

func main() {
	var a interface{} // 什么一个可以接受所有结构的接口
	var point Point = Point{1, 2}
	a = point
	fmt.Println(a)
	var b Point
	//b = a 无法转换，类型无法判定是否一致
	b, flag := a.(Point) // 这时候就要使用类型断言 - a.(是不是这个类型的) ,如果失败会抛出异常
	// 使用flag表示带检测的断言
	// if b, flag := a.(Point); flag {} 还可以这么写的
	if flag {
		fmt.Println(b)
	} else {
		fmt.Println("convert error")
	}

	// 类型判断
	var item1 = 10
	var item2 float32 = 11.2
	var item3 = "justin"
	judgeType(item1, item2, item3)
}
