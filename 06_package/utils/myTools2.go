package utils

import "fmt"

func Main() {
	privateFunction()
}

// 同一个包下不可以重复定义变量和函数。
//func PublicFunction()  {
//
//}

func init() {
	fmt.Println("2 init")
}

func main() {
	fmt.Println("2 main")
}
