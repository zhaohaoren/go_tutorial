// init 函数
package main

import "fmt"

var num = initVal()

func main() {
	fmt.Println("main function invoke")
}

// 每一个文件都可以有一个init函数，该函数会在main函数被调用前执行。
// 这个函数和main函数一样，在一个包中是可以重名的
func init() {
	fmt.Println("init function invoke")
}

// 如果有全局变量定义那么最先执行的是全部变量定义
func initVal() int {
	fmt.Println("init val invoke")
	return 90
}

// 执行顺序: 全局变量定义-> init方法 -> main方法，如果有引入的，则先执行引入包的方法
// init 一般用于一些初始化工作