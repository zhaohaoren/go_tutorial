// 指针
package main

import "fmt"

func main() {

	// 获取变量的地址
	var i int = 10
	fmt.Println("i的地址为:", &i)

	// 指针类型：保存地址的变量，这个变量指向的地址空间存的值才是真实值
	// * 取地址对应的值 & 取变量的地址
	var ptr *int = &i // *int 带*号 代表这是一个指针变量
	fmt.Println(ptr)
	// 获取指针对应的值
	fmt.Println(*ptr)
	// 指针同样也有自己的地址
	fmt.Println(&ptr)
	// 通过指针修改指针指向空间的值从而修改变量
	*ptr = 11
	fmt.Println(i)
	// 下面反例：
	//var ptr2 *int = i  指针变量只能接受地址，不可以接受变量
	//var ptr2 *float32 = &i	指针类型和变量需要一致

}
