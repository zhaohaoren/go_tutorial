// goto
package main

import "fmt"

func main() {

	// 可以使用goto但是不推荐使用
	var i = 10
	fmt.Println("你好1")
	if i == 10 {
		// 一般结合if条件进行goto跳转
		goto label
	}
	fmt.Println("你好2")
	fmt.Println("你好3")
	fmt.Println("你好4")
	fmt.Println("你好5")
	fmt.Println("你好6")
	fmt.Println("你好7")
label:
	fmt.Println("你好8")
}
