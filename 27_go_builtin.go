// golang的一些内置函数
package main

import "fmt"

func main() {
	// len
	// new: 分配内存 - 主要针对值类型（int float等，数组，结构体）
	var n1 = 100
	fmt.Printf("type:%T var:%v addr:%v \n", n1, n1, &n1)
	var n2 = new(int) // 没有值默认为0， 返回的是该分配的内存的地址。 n2是一个变量：存的是new的地址，&n2是n2这个变量的地址
	fmt.Printf("type:%T var:%v addr:%v", n2, n2, &n2)

	// make: 分配内存 - 主要针对引用类型（chan map slice）
}
