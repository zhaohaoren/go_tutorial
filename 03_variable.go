//go变量的使用

package main

import "fmt"

//全局变量
var (
	n1 = 100
	n2 = 200
	n3 = "hello"
)

var i = 10
//j := 11 // 这个是不可以的，这种赋值不属于 初始化，而是var j int; j = 11 的缩写。所以go中定义变量只存在 var 这种定义才是真的定义方式。

func main() {
	//申明方式1：定义变量，如果不赋值就使用默认值，int 默认值为0 string 为空串
	var i int
	//变量赋值
	i = 10
	var j int = 10
	//变量使用
	fmt.Println("i=", i)
	fmt.Println("j=", j)

	//申明方式2：类型推导，自动识别赋值的类型
	var num = 10.11
	fmt.Println("num=", num)

	//申明方式3：使用:= 来进行申明变量并赋值
	name := "justin"
	fmt.Println("name=", name)

	// 常量
	// const修饰，定义时候必须初始化，不可以修改，常量只能修饰bool，数值类型和string类型
	const constVal int = 10
	//constVal = 11
	fmt.Println(constVal)
	// 其他写法
	const (
		aC = 1
		bC = 2
	)
	const (
		zC       = iota       // 默认为0
		z2C                   //下面都自增1 = 1, 一种专业写法，其实省略了 = iota
		z3C                   //为2	省略了iota
		z4C, z5C = iota, iota // 如果两个在一行，那么两个值都一样
	)
	fmt.Println(aC, bC)
	fmt.Println(zC, z2C, z3C, z4C, z5C)

	// ----------------------------------

	// go一次申明定义多个变量
	var a, b, c int
	fmt.Println(a, b, c)

	var d, e, f = 1, 2, 3
	fmt.Println(d, e, f)

	j, h, k := 1, 2, 3
	fmt.Println(j, h, k)

	fmt.Println(n1, n2, n3)

}
