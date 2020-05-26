// 反射
package main

import (
	"fmt"
	"reflect"
)

// 反射的价值在于自己写框架！ java亦然

// 反射常见应用场景
// 1:json格式化的时候做标签
// 2:适配器函数

// 基本介绍
// 1. 反射可以在运行时动态获取变量的各种信息，比如变量的类型，类别。如果是结构体变量还可以获取到结构体信息（结构体字段和方法）
// 2. 反射可以修改变量的值，可以调用关联的方法

// 反射里面有两个特别重要的类型：type 和 value
// 对于反射的开发中 变量，interface{}，reflect.Value类型是互相可以转换的
type RStudent struct {
	Name string `json:"name"`
	Age  int
}

func (s RStudent) Fun1() {
	fmt.Println("RStudent fun1")
}
func (s RStudent) Fun2(n int, m int) {
	fmt.Println("RStudent fun2", n, m)
}

func main() {

	var num int = 100
	reflectDemoBase(num)

	var stu = RStudent{
		Name: "justin",
		Age:  24,
	}
	reflectStruct(stu)

	changeVal(&num) // 想要reflect去修改值，那么必须传地址给reflect.XXOf()
	fmt.Println(num)
	changeStruct(&stu)
	fmt.Println(stu)
}

// 反射获取类和值信息
func reflectDemoBase(b interface{}) {

	// 所有的变量都可以由反射转为 Type和Value类型
	// 反射获取到 reflect.Type --- 是一个接口类型 （变量的类型自然是接口接受）
	rType := reflect.TypeOf(b)
	fmt.Println(rType)
	fmt.Println(rType.Kind()) // 获取类别

	// 反射获取到 reflect.Value --- 是一个结构体  （变量的值自然用结构体存储信息）
	rValue := reflect.ValueOf(b)
	fmt.Println(rValue)
	fmt.Println(rValue.Kind()) // 获取类别
	// Type是类型，Kind是类别，两个可能是相同的也可能是不同的
	// int 的type和kind都是 int； 但是 RStudent的type是 包名.RStudent，但是Kind是struct

	// 还可以将这些转回去 2步走： 反射 -提升为-> interface -强转为-> type
	iV := rValue.Interface() // 需要先转为interfaces
	num2 := iV.(int)         // 然后使用类型断言进行转回
	fmt.Println(num2)
}
func reflectStruct(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	rValue := reflect.ValueOf(b)
	fmt.Println(rValue)
	// 遍历结构体变量
	kd := rValue.Kind()
	if kd != reflect.Struct { //判断是不是结构体
		fmt.Println("no struct! exit...")
		return
	}
	num := rValue.NumField()
	fmt.Println("struct has field number:", num)
	for i := 0; i < num; i++ {
		fmt.Println("struct field:", i, rValue.Field(i)) // value.Field(index) 获取字段值
		// 获取字段的标签：就是json的别名,需要通过reflect.Type来获取
		tagVal := rType.Field(i).Tag.Get("json")
		if tagVal != "" { // 如果有tag就显示
			fmt.Printf("field: %d ,tag is: %v \n", i, tagVal)
		}
	}
	// 遍历结构体的方法
	numFunc := rValue.NumMethod()
	fmt.Println("func num is:", numFunc)
	rValue.Method(0).Call(nil) // 获取结构体第一个方法并调用， 方法的排序规则是按照方法名字进行排序的
	//带参数的调用
	var params []reflect.Value // 这里为什么不要make
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(11))
	rValue.Method(1).Call(params)

	// 还可以将这些转回去
	iV := rValue.Interface() // 需要先转为interfaces
	stu := iV.(RStudent)
	fmt.Println(stu)
}

// 反射修改值
func changeVal(b interface{}) {
	rVal := reflect.ValueOf(b)
	// 前面传入的是指针类型，所以需要通过Elem来获取指针类型指向的值来进行修改
	rVal.Elem().SetInt(20)
}

func changeStruct(b interface{}) {
	rValue := reflect.ValueOf(b)
	rValue.Elem().Field(0).SetString("jack")
}
