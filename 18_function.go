// 函数
package main

import "fmt"

func main() {
	fmt.Println(cal(19, 12, '+'))

	_, _ = f1() // 使用_ 表示忽略该返回值：不用任何东西去接收

	fmt.Println(testVar(func(i int, i2 int) int {
		return i + i2
	}, 1, 2))

	testRet()

	testParams(1, 1, 1)
}

// 函数名(形参1 类型，....新参n 类型) 返回值类型
func cal(num1 float64, num2 float64, operator byte) float64 {
	if operator == '*' {
		return num2 * num1
	}
	return num1 + num2
}

// go 可以返回多个值
func f1() (int, int) {
	return 1, 1
}

//go 不支持重载
//func f1(n int){
//
//}

// 函数在go中本质也是变量，可以作为形参传递，（这样可以回调）
func testVar(myFunction func(int, int) int, num1 int, num2 int) int {
	return myFunction(num1, num2)
}

// 直接为返回值命名, return的时候直接取函数内部对应的变量名称的变量值返回
func testRet() (n1 int, n2 int) {
	n1 = 1 // 这里是直接复制的，不是申明变量并赋值，申明被返回值给先申明了
	n2 = 2
	return
}

// go中支持可变参数: 可变参数要放在参数的最后
func testParams(args ...int) int {
	// args 是一个slice切片
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
