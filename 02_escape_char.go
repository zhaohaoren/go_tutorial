//转义

package main

import "fmt"

func main()  {
	fmt.Println("name\tjustin")
	fmt.Print("age\t24\n")
	fmt.Println("\\hi\\")
	fmt.Println("雪山飞狐\r找") // 只有这个不熟悉，回车符号会把前面的字符都给删除了（回车不换行从头再输出！）不同版本go表现好像不同
}