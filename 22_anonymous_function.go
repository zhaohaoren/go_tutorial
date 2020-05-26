// 匿名函数
package main

import "fmt"

var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 + n2
	}
)

func main() {

	// 如果一个函数只会被使用一次可以使用匿名函数

	// 匿名函数1: 定义函数同时调用
	sum1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(1, 1)
	fmt.Println(sum1)

	// 匿名函数2: 定义函数给变量，通过变量多次调用
	sum2 := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Println(sum2(1, 1))

	// 匿名函数3: 全局匿名函数
	fmt.Println(Fun1(1, 2))
}
