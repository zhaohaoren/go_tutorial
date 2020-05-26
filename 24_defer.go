// defer 关键字
package main

import "fmt"

func main() {
	// defer主要作用是在函数结束后去释放资源， 比如在打开文件的资源操作后，立刻调用defer 来释放资源，两者结对出现。

	deferDemo()
	deferDemo2()

	deferWithValue()

}

func deferDemo() {
	//defer 语句会不执行，而是放入到一个defer栈中，等函数执行完了会再去弹出defer栈中的内容
	defer fmt.Println("defer statement1")
	defer fmt.Println("defer statement2")
	fmt.Println("no defer statement")
}

// 也就是说defer 栈是针对每一个函数都是一个，每次函数执行完弹出defer栈。而不是整个程序在一个栈中，全结束才弹出。
func deferDemo2() {
	defer fmt.Println("defer2 statement1")
	defer fmt.Println("defer2 statement2")
	fmt.Println("no defer2 statement")
}

// defer栈会将变量值也带上 一并拷入栈中
func deferWithValue() {
	var n1 = 1
	var n2 = 2
	defer fmt.Println("defer n1=", n1)
	defer fmt.Println("defer n2=", n2)
	n1++
	n2++
	fmt.Println("n1++ =", n1)
	fmt.Println("n2++ =", n2)
}
