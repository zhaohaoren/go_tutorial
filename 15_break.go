// break语句
package main

import "fmt"

func main() {
	// break 语句

	fmt.Println("ssss")

	// break + label 可以表明需要终止的是哪一层循环，跳出对应的for循环
	// 不是指跳到对应位置代码块继续执行，而是指终止这层的循环，所以label必须加在循环上头不能乱加
label:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label
			}
		}
	}
}
