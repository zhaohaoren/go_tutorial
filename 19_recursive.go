// 递归
package main

import (
	"fmt"
)

func main() {
	test(4)

	fmt.Println("feb num is: ", feb(6))
	fmt.Println("tao day1 is: ", tao(1))
}

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println(n)

}

func feb(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return feb(n-1) + feb(n-2)
}

// 猴子吃桃
// 每天吃桃子一半再多吃一个，第10天就只剩1个
// 转化为式子： f(n-1) = (f(n) + 1) * 2 -> f(n) = (f(n+1) +1 ) *2
func tao(n int) int {
	if n == 10 { // 第10天只有一个
		return 1
	}
	return (tao(n+1) + 1) * 2
}
