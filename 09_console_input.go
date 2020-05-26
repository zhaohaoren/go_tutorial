// 获取用户控制台输入
package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("输入name:")
	fmt.Scanln(&name)
	fmt.Println("输入age:")
	fmt.Scanln(&age)

	fmt.Println(name, age)
}
