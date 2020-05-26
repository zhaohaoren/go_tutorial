//进制

package main

import "fmt"

func main() {
	// 二进制：go不可以直接表示二进制，但可以输出二进制
	var i int = 5
	fmt.Printf("%b\n", i)

	// 八进制：go可以直接表示8进制，0开头
	var j int = 011
	fmt.Println(j)

	// 16进制：可以直接表示，0x或者0X开头
	var k int = 0x11
	fmt.Println(k)

	bit_cal()
}

// 位运算
func bit_cal() {

	fmt.Println("-------------")
	//3个位运算
	// & 1 1: 1
	fmt.Println(2 & 3) // 0010 0011 -> 0010
	// | 1 0: 1
	fmt.Println(2 | 3) // 0011
	// ^ 1 0: 1 不同为true 同都为false
	fmt.Println(2 ^ 3)  // 0001
	fmt.Println(-2 ^ 2) // 补码：1111_1110 0000_0010 -> 1111_1100 -> 反码：1111_1011 -> 原码：1000_0100(符号位不变其他取反)

	//2个移位运算
	fmt.Println(1 >> 2)
	fmt.Println(1 << 2)
}

/*
原码
	正数=本身
	负数=本身，头部为1的符号位 （反码符号位不变取反）
反码
	正数=本身
	负数=符号位不变，其他取反  （补码-1）
补码
	正数=本身
	负数=反码+1
 */
