// 关系运算符

package main

import "fmt"

func main() {
	// 关系运算符
	var n1 = 9
	var n2 = 8
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)
	var flag = n1 == n2
	fmt.Println(flag)

	//逻辑运算符 -- 注意其短路运算
	var A = true
	var B = false
	fmt.Println(A && B)
	fmt.Println(A || B)
	fmt.Println(!A)
	var num = 10
	if num > 0 && num < 60 {
		fmt.Println(num)
	}
	if A || bool_func() {
		fmt.Println("it won't print false")
	}

	//  赋值运算符
	var f1 = 10
	var f2 = 20
	f1 += f2
	fmt.Println(f1)
	f1 -= f2
	fmt.Println(f1)
	f1 *= f2
	fmt.Println(f1)
	f1 /= f2
	fmt.Println(f1)
	// 交换变量
	fmt.Println(f1, f2)
	f1 += f2
	f2 = f1 - f2
	f1 = f1 - f2
	fmt.Println(f1, f2)

	//位运算符
	// * & << >> 等

}

func bool_func() bool {
	fmt.Println("return false")
	return false
}
