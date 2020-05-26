package main

import "fmt"

func main() {

	//算术运算符： + _ * / % ++ --

	var i = 10
	//var j = i++ go 不可以这样使用！++ 只能单独作为一个式子。并且没有++i这种东西
	i++
	var j = i
	fmt.Println(i, " ", j)
}
