// 闭包
package main

import (
	"fmt"
	"strings"
)

// 闭包指一个函数和其相关的引用环境组合成一个整体
// 闭包的本质是一个函数： 这个函数和自己外部引用的变量构成一个整体

func main() {
	fmt.Println(AddUpper()(1))

	// 返回的闭包，可以理解为一个对象，所以每次调用这个闭包函数的时候，里面的值都会在之前的值上进行操作。
	// 可以理解为一个对象只有一个方法，这个方法闭包的外部变量都是这个对象的属性，只要还使用这个方法，那么操作的就是该对象的属性
	var f = AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16

	// 实践：
	// 给一个后缀，返回一个闭包，该闭包可以给文件名添加后缀（如果有了该后缀就直接原返回）
	// 给我的感觉有点像工厂方法，依据传入的参数返回具有特定特性的东西
	var suffix = makeSuffix(".jpg")
	fmt.Println(suffix("a.jpg"))
	fmt.Println(suffix("b"))

}

// 闭包累加器
func AddUpper() func(int) int {
	var n = 10
	return func(i int) int {
		n = n + i
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(fileName string) string {
		if strings.HasSuffix(fileName, suffix) {
			return fileName
		} else {
			return fileName + suffix
		}
	}
}
