//第一个Go程序
//可以很清晰的感受到：Go语言对规范的唯一性，一种问题只有一种方案。
//go推荐使用行注释而不是块注释

package main //main函数的包名必须为main否则无法执行

import "fmt" // 引入包就必须要使用,不适用无法通过编译

func main() {
	var txt = "hello world" // 变量申明就必须要使用，如果不使用就报错
	fmt.Println(txt)        // 代码不需要加“;”,并且代码必须一行一行执行
}
