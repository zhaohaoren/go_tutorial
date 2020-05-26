// 异常
package main

import (
	"errors"
	"fmt"
)

func main() {

	// go为了简洁没有使用 try catch finally
	// go 使用 defer panic recover
	testException()
	fmt.Println("code after exception")

	err := customError("your.conf")
	if err != nil {
		//panic 将异常直接抛出 并 程序退出，后面程序不再执行（不是退出函数，是退出整个程序）
		panic(err) // panic 可以接受任何值做参数，如果接收到error变量，那么会直接输出错误信息并退出程序
	}
	fmt.Println(err)
	fmt.Println("code after custom error")
}

func testException() {
	// defer+recover 来处理错误
	defer func() {
		err := recover() // recover 内置函数可以捕获到异常
		if err != nil {
			fmt.Println(err)
		}
	}()
	// 使用defer 在异常抛出前定义好一个匿名操作函数并调用，在发生异常的时候函数就会结束，这时候调用defer函数里面recover会捕获到异常
	// 如果有结果就输出 -- 这也是比较好玩的一个思路

	var i = 10
	var j = 0
	var res = i / j
	fmt.Println(res)
}

// 自定义错误
func customError(fileName string) error {
	if fileName != "my.conf" {
		// 这里就是自定义错误，主要自定义错误的信息
		return errors.New("read config file fail")
	} else {
		return nil
	}
}
